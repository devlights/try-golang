<!-- https://github.com/golang/example/blob/master/slog-handler-guide/README.md を元にしたメモ -->
<!-- 2024-10-21 時点の内容です -->

# `slog` ハンドラの書き方

This document is maintained by Jonathan Amsterdam `jba@google.com`.

> この文書はジョナサン・アムステルダム jba@google.com によって管理されています。

# Contents

1. [Introduction](#introduction)
1. [Loggers and their handlers](#loggers-and-their-handlers)
1. [Implementing `Handler` methods](#implementing-`handler`-methods)
	1. [The `Enabled` method](#the-`enabled`-method)
	1. [The `Handle` method](#the-`handle`-method)
	1. [The `WithAttrs` method](#the-`withattrs`-method)
	1. [The `WithGroup` method](#the-`withgroup`-method)
	1. [Testing](#testing)
1. [General considerations](#general-considerations)
	1. [Copying records](#copying-records)
	1. [Concurrency safety](#concurrency-safety)
	1. [Robustness](#robustness)
	1. [Speed](#speed)


# Introduction

The standard library’s `log/slog` package has a two-part design.
A "frontend," implemented by the `Logger` type,
gathers structured log information like a message, level, and attributes,
and passes them to a "backend," an implementation of the `Handler` interface.
The package comes with two built-in handlers that usually should be adequate.
But you may need to write your own handler, and that is not always straightforward.
This guide is here to help.

> 標準ライブラリのlog/slogパッケージは、2つの部分で構成されている。 Logger型によって実装された "フロントエンド "は、メッセージ、レベル、属性などの構造化されたログ情報を収集し、ハンドラーインターフェースの実装である "バックエンド "に渡します。 パッケージには2つの組み込みハンドラが付属しており、通常はこれで十分です。 しかし、独自のハンドラを書く必要があるかもしれないし、それは必ずしも簡単ではない。 このガイドはその手助けをする。

# Loggers and their handlers

Writing a handler requires an understanding of how the `Logger` and `Handler`
types work together.

> ハンドラーを書くには、ロガーとハンドラータイプがどのように連動するかを理解する必要があります。

Each logger contains a handler. Certain `Logger` methods do some preliminary work,
such as gathering key-value pairs into `Attr`s, and then call one or more
`Handler` methods. These `Logger` methods are `With`, `WithGroup`,
and the output methods.

> 各ロガーはハンドラを含む。 特定のロガーメソッドは、キーと値のペアをAttrsに集めるなどの予備作業を行い、1つ以上のハンドラーメソッドを呼び出します。 これらのロガーメソッドはWith、WithGroup、およびoutputメソッドです。

An output method fulfills the main role of a logger: producing log output.
Here is a call to an output method:

> 出力メソッドは、ロガーの主な役割であるログ出力を行う。 以下は出力メソッドの呼び出しです：


    logger.Info("hello", "key", value)

There are two general output methods, `Log`, and `LogAttrs`. For convenience,
there is an output method for each of four common levels (`Debug`, `Info`,
`Warn` and `Error`), and corresponding methods that take a context (`DebugContext`,
`InfoContext`, `WarnContext` and `ErrorContext`).

> 一般的な出力メソッドには、LogとLogAttrsの2つがある。 便宜上、4つの一般的なレベル（Debug、Info、Warn、Error）それぞれに対応する出力メソッドがあり、対応するメソッドはコンテキスト（DebugContext、InfoContext、WarnContext、ErrorContext）を受け取ります。

Each `Logger` output method first calls its handler's `Enabled` method. If that call
returns true, the method constructs a `Record` from its arguments and calls
the handler's `Handle` method.

> 各ロガー出力メソッドは、最初にハンドラーの Enabled メソッドを呼び出します。 この呼び出しが真を返す場合、メソッドは引数からレコードを構築し、ハンドラの Handle メソッドを呼び出します。

As a convenience and an optimization, attributes can be added to
`Logger` by calling the `With` method:

> 利便性と最適化として、属性は With メソッドを呼び出すことでロガーに追加できます：


    logger = logger.With("k", v)

This call creates a new `Logger` value with the argument attributes; the
original remains unchanged.
All subsequent output from `logger` will include those attributes.
A logger's `With` method calls its handler's `WithAttrs` method.

> この呼び出しは、引数属性で新しいロガー値を作成します。 元の値は変更されません。ロガーからの後続の出力はすべて、それらの属性を含みます。 ロガーの With メソッドは、ハンドラーの WithAttrs メソッドを呼び出します。

The `WithGroup` method is used to avoid key collisions in large programs
by establishing separate namespaces.
This call creates a new `Logger` value with a group named "g":

> WithGroupメソッドは、別々のネームスペースを確立することで、大きなプログラムでのキーの衝突を避けるために使用されます。 この呼び出しは、"g" という名前のグループを持つ新しいロガー値を作成します：


    logger = logger.WithGroup("g")

All subsequent keys for `logger` will be qualified by the group name "g".
Exactly what "qualified" means depends on how the logger's handler formats the
output.
The built-in `TextHandler` treats the group as a prefix to the key, separated by
a dot: `g.k` for a key `k`, for example.
The built-in `JSONHandler` uses the group as a key for a nested JSON object:

> それ以降のロガーのキーはすべて、グループ名 "g "で修飾される。 修飾された "が何を意味するかは、ロガーのハンドラーがどのように出力を フォーマットするかによって異なります。 組み込みのTextHandlerは、グループをキーのプレフィックスとして扱い、ドットで区切ります。 組み込みのJSONHandlerは、グループをネストされたJSONオブジェクトのキーとして使用します：


    {"g": {"k": v}}

A logger's `WithGroup` method calls its handler's `WithGroup` method.

> ロガーの WithGroup メソッドは、ハンドラーの WithGroup メソッドを呼び出します。

# Implementing `Handler` methods

We can now talk about the four `Handler` methods in detail.
Along the way, we will write a handler that formats logs using a format
reminiscent of YAML. It will display this log output call:

> これで4つのハンドラーメソッドについて詳しく説明できます。 途中で、YAMLを連想させるフォーマットを使ってログをフォーマットするハンドラーを書きます。 このログ出力は以下のように呼び出します：


    logger.Info("hello", "key", 23)

something like this:

> こんな感じです


    time: 2023-05-15T16:29:00
    level: INFO
    message: "hello"
    key: 23
    ---

Although this particular output is valid YAML,
our implementation doesn't consider the subtleties of YAML syntax,
so it will sometimes produce invalid YAML.
For example, it doesn't quote keys that have colons in them.
We'll call it `IndentHandler` to forestall disappointment.

> この特定の出力は有効なYAMLですが、私たちの実装はYAMLの構文の繊細さを考慮していないので、ときどき無効なYAMLを生成します。 たとえば、コロンを含むキーを引用しません。 がっかりしないように、IndentHandlerと呼ぶことにしよう。

We begin with the `IndentHandler` type
and the `New` function that constructs it from an `io.Writer` and options:

> IndentHandler型と、それをio.Writerとオプションから構築するNew関数から始めます:

```
type IndentHandler struct {
	opts Options
	// TODO: state for WithGroup and WithAttrs
	mu  *sync.Mutex
	out io.Writer
}

type Options struct {
	// Level reports the minimum level to log.
	// Levels with lower levels are discarded.
	// If nil, the Handler uses [slog.LevelInfo].
	Level slog.Leveler
}

func New(out io.Writer, opts *Options) *IndentHandler {
	h := &IndentHandler{out: out, mu: &sync.Mutex{}}
	if opts != nil {
		h.opts = *opts
	}
	if h.opts.Level == nil {
		h.opts.Level = slog.LevelInfo
	}
	return h
}
```

We'll support only one option, the ability to set a minimum level in order to
suppress detailed log output.
Handlers should always declare this option to be a `slog.Leveler`.
The `slog.Leveler` interface is implemented by both `Level` and `LevelVar`.
A `Level` value is easy for the user to provide,
but changing the level of multiple handlers requires tracking them all.
If the user instead passes a `LevelVar`, then a single change to that `LevelVar`
will change the behavior of all handlers that contain it.
Changes to `LevelVar`s are goroutine-safe.

> 詳細なログ出力を抑制するための最小レベルを設定する機能です。 ハンドラは常にこのオプションをslog.Levelerであると宣言する必要があります。 slog.LevelerインターフェイスはLevelとLevelVarの両方によって実装されます。 Level 値をユーザが提供するのは簡単ですが、複数のハンドラのレベルを変更するには、それらすべてを追跡する必要があります。 代わりにユーザが LevelVar を渡す場合、その LevelVar への単一の変更は、それを含むすべてのハンドラの動作を変更します。 LevelVar への変更はゴルーチンセーフです。

You might also consider adding a `ReplaceAttr` option to your handler,
like the [one for the built-in
handlers](https://pkg.go.dev/log/slog#HandlerOptions.ReplaceAttr).
Although `ReplaceAttr` will complicate your implementation, it will also
make your handler more generally useful.

> また、ビルトインハンドラのように、ReplaceAttrオプションをハンドラに追加することも考えてみてください。 ReplaceAttrは実装を複雑にしますが、ハンドラをより一般的に使えるようにします。

The mutex will be used to ensure that writes to the `io.Writer` happen atomically.
Unusually, `IndentHandler` holds a pointer to a `sync.Mutex` rather than holding a
`sync.Mutex` directly.
There is a good reason for that, which we'll explain [later](#getting-the-mutex-right).

> このミューテックスは、io.Writerへの書き込みがアトミックに行われるようにするために使用される。 珍しいことに、IndentHandlerはsync.Mutexを直接保持するのではなく、sync.Mutexへのポインタを保持する。 これには理由があり、後で説明する。

Our handler will need additional state to track calls to `WithGroup` and `WithAttrs`.
We will describe that state when we get to those methods.

> 私たちのハンドラーは、WithGroupとWithAttrsの呼び出しを追跡するために追加のステートを必要とする。 このステートについては、これらのメソッドのところで説明する。

## The `Enabled` method

The `Enabled` method is an optimization that can avoid unnecessary work.
A `Logger` output method will call `Enabled` before it processes any of its arguments,
to see if it should proceed.

> Enabledメソッドは、不要な作業を回避する最適化です。 ロガーの出力メソッドは、引数を処理する前にEnabledを呼び出し、処理を続行するかどうかを確認します。

The signature is

> シグネチャは


    Enabled(context.Context, Level) bool

The context is available to allow decisions based on contextual information.
For example, a custom HTTP request header could specify a minimum level,
which the server adds to the context used for processing that request.
A handler's `Enabled` method could report whether the argument level
is greater than or equal to the context value, allowing the verbosity
of the work done by each request to be controlled independently.

> コンテキストは、コンテキスト情報に基づいた決定を可能にするために 利用できる。 例えば、カスタムHTTPリクエストヘッダは最小レベルを指定することができ、 サーバはそのリクエストの処理に使われるコンテキストにそれを追加する。 ハンドラの Enabled メソッドは、引数レベルがコンテキスト値以上であるかどうかを 報告することができ、各リクエストで行われる処理の冗長性を独立に制御することが できます。

Our `IndentHandler` doesn't use the context. It just compares the argument level
with its configured minimum level:

> 私たちのIndentHandlerはコンテキストを使用しません。 引数のレベルを設定された最小レベルと比較するだけです：

```
func (h *IndentHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}
```

## The `Handle` method

The `Handle` method is passed a `Record` containing all the information to be
logged for a single call to a `Logger` output method.
The `Handle` method should deal with it in some way.
One way is to output the `Record` in some format, as `TextHandler` and `JSONHandler` do.
But other options are to modify the `Record` and pass it on to another handler,
enqueue the `Record` for later processing, or ignore it.

> Handleメソッドは、Logger出力メソッドへの1回の呼び出しでログに記録されるすべての情報を含むRecordを渡されます。 Handleメソッドは、何らかの方法でそれを処理する必要があります。 一つの方法は、TextHandlerやJSONHandlerが行うように、何らかのフォーマットでRecordを出力することです。 しかし、他のオプションとしては、Recordを修正して別のハンドラに渡す、後で処理するためにRecordをエンキューする、または無視する、などがあります。

The signature of `Handle` is

> `Handle` メソッドのシグネチャは


    Handle(context.Context, Record) error

The context is provided to support applications that provide logging information
along the call chain. In a break with usual Go practice, the `Handle` method
should not treat a canceled context as a signal to stop work.

> コンテキストは、コールチェーンに沿ってロギング情報を提供するアプリケーションをサポートするために提供されます。 通常のGoの慣習に反して、Handleメソッドはキャンセルされたコンテキストを作業を停止するシグナルとして扱うべきではありません。

If `Handle` processes its `Record`, it should follow the rules in the
[documentation](https://pkg.go.dev/log/slog#Handler.Handle).
For example, a zero `Time` field should be ignored, as should zero `Attr`s.

> HandleがそのRecordを処理する場合、ドキュメントのルールに従うべきである。 例えば、ゼロのTimeフィールドは、ゼロのAttrsと同様に無視されるべきです。

A `Handle` method that is going to produce output should carry out the following steps:

> 出力を行う `Handle` メソッドは、以下のステップを実行しなければなりません：

1. Allocate a buffer, typically a `[]byte`, to hold the output.
It's best to construct the output in memory first,
then write it with a single call to `io.Writer.Write`,
to minimize interleaving with other goroutines using the same writer.

> バッファ（通常は[]byte）を確保し、出力を保持する。 同じライターを使用する他のゴルーチンとのインターリーブを最小限にするために、最初にメモリ上に出力を構築し、io.Writer.Writeを1回呼び出すだけで書き込むのが最善です。

2. Format the special fields: time, level, message, and source location (PC).
As a general rule, these fields should appear first and are not nested in
groups established by `WithGroup`.

> 時間、レベル、メッセージ、およびソースの場所 (PC) という特別なフィールドをフォーマットします。 一般的なルールとして、これらのフィールドは最初に表示され、WithGroup によって確立されたグループにネストされません。

3. Format the result of `WithGroup` and `WithAttrs` calls.

> WithGroupおよびWithAttrs呼び出しの結果をフォーマットする。

4. Format the attributes in the `Record`.

> レコードの属性をフォーマットする。

5. Output the buffer.

> バッファを出力する。

That is how `IndentHandler.Handle` is structured:

> `IndentHandler.Handle`の構造は以下のようになります:

```
func (h *IndentHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := make([]byte, 0, 1024)
	if !r.Time.IsZero() {
		buf = h.appendAttr(buf, slog.Time(slog.TimeKey, r.Time), 0)
	}
	buf = h.appendAttr(buf, slog.Any(slog.LevelKey, r.Level), 0)
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		buf = h.appendAttr(buf, slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", f.File, f.Line)), 0)
	}
	buf = h.appendAttr(buf, slog.String(slog.MessageKey, r.Message), 0)
	indentLevel := 0
	// TODO: output the Attrs and groups from WithAttrs and WithGroup.
	r.Attrs(func(a slog.Attr) bool {
		buf = h.appendAttr(buf, a, indentLevel)
		return true
	})
	buf = append(buf, "---\n"...)
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}
```

The first line allocates a `[]byte` that should be large enough for most log
output.
Allocating a buffer with some initial, fairly large capacity is a simple but
significant optimization: it avoids the repeated copying and allocation that
happen when the initial slice is empty or small.
We'll return to this line in the section on [speed](#speed)
and show how we can do even better.

> 最初の行は、ほとんどのログ出力に十分な大きさの `[]byte` を割り当てています。 初期のかなり大きな容量のバッファを割り当てることは、単純だが重要な最適化です。 速度に関するセクションでこの行に戻り、さらに良い方法を示します。

The next part of our `Handle` method formats the special attributes,
observing the rules to ignore a zero time and a zero PC.

> Handleメソッドの次の部分は、時間ゼロとPCゼロを無視するルールを守って、特別な属性をフォーマットします。

Next, the method processes the result of `WithAttrs` and `WithGroup` calls.
We'll skip that for now.

> 次に、このメソッドはWithAttrsとWithGroupの呼び出しの結果を処理します。今は省略します。

Then it's time to process the attributes in the argument record.
We use the `Record.Attrs` method to iterate over the attributes
in the order the user passed them to the `Logger` output method.
Handlers are free to reorder or de-duplicate the attributes,
but ours does not.

> 次に、引数レコードの属性を処理します。 Record.Attrsメソッドを使用して、ユーザーがLogger出力メソッドに渡した順序で属性を繰り返し処理します。 ハンドラは自由に属性を並べ替えたり、重複を取り除いたりできますが、私たちのハンドラはそうしません。

Lastly, after adding the line "---" to the output to separate log records,
our handler makes a single call to `h.out.Write` with the buffer we've accumulated.
We hold the lock for this write to make it atomic with respect to other
goroutines that may be calling `Handle` at the same time.

> 最後に、"---"という行を出力に追加してログレコードを区切った後、ハンドラは、蓄積したバッファを使ってh.out.Writeを1回呼び出します。 同時にHandleを呼び出すかもしれない他のゴルーチンに対してアトミックにするために、この書き込みのロックを保持します。

At the heart of the handler is the `appendAttr` method, responsible for
formatting a single attribute:

> ハンドラの中心にあるのはappendAttrメソッドで、一つの属性をフォーマットする役割を担っています：

```
func (h *IndentHandler) appendAttr(buf []byte, a slog.Attr, indentLevel int) []byte {
	// Resolve the Attr's value before doing anything else.
	a.Value = a.Value.Resolve()
	// Ignore empty Attrs.
	if a.Equal(slog.Attr{}) {
		return buf
	}
	// Indent 4 spaces per level.
	buf = fmt.Appendf(buf, "%*s", indentLevel*4, "")
	switch a.Value.Kind() {
	case slog.KindString:
		// Quote string values, to make them easy to parse.
		buf = fmt.Appendf(buf, "%s: %q\n", a.Key, a.Value.String())
	case slog.KindTime:
		// Write times in a standard way, without the monotonic time.
		buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value.Time().Format(time.RFC3339Nano))
	case slog.KindGroup:
		attrs := a.Value.Group()
		// Ignore empty groups.
		if len(attrs) == 0 {
			return buf
		}
		// If the key is non-empty, write it out and indent the rest of the attrs.
		// Otherwise, inline the attrs.
		if a.Key != "" {
			buf = fmt.Appendf(buf, "%s:\n", a.Key)
			indentLevel++
		}
		for _, ga := range attrs {
			buf = h.appendAttr(buf, ga, indentLevel)
		}
	default:
		buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value)
	}
	return buf
}
```

It begins by resolving the attribute, to run the `LogValuer.LogValue` method of
the value if it has one. All handlers should resolve every attribute they
process.

> まず属性を解決し、その値にLogValuer.LogValueメソッドがあればそれを実行します。 すべてのハンドラは、処理するすべての属性を解決する必要があります。

Next, it follows the handler rule that says that empty attributes should be
ignored.

> 次に、空の属性は無視されるべきだというハンドラ・ルールに従います。

Then it switches on the attribute kind to determine what format to use. For most
kinds (the default case of the switch), it relies on `slog.Value`'s `String` method to
produce something reasonable. It handles strings and times specially:
strings by quoting them, and times by formatting them in a standard way.

> 次に、属性の種類を切り替えて、使用する書式を決定します。 ほとんどの種類（スイッチのデフォルトケース）では、slog.ValueのStringメソッドに依存して、妥当なものを生成します。 文字列は引用符で囲み、時刻は標準的な方法でフォーマットします。

When `appendAttr` sees a `Group`, it calls itself recursively on the group's
attributes, after applying two more handler rules.
First, a group with no attributes is ignored&mdash;not even its key is displayed.
Second, a group with an empty key is inlined: the group boundary isn't marked in
any way. In our case, that means the group's attributes aren't indented.

> appendAttrがグループを見つけると、さらに2つのハンドラ・ルールを適用した後、グループの属性に対して再帰的に自分自身を呼び出します。 まず、属性のないグループは無視され、キーも表示されません。 第二に、キーが空のグループはインライン化され、グループの境界は何もマークされません。 この場合、グループの属性はインデントされません。

## The `WithAttrs` method

One of `slog`'s performance optimizations is support for pre-formatting
attributes. The `Logger.With` method converts key-value pairs into `Attr`s and
then calls `Handler.WithAttrs`.
The handler may store the attributes for later consumption by the `Handle` method,
or it may take the opportunity to format the attributes now, once,
rather than doing so repeatedly on each call to `Handle`.

> slogのパフォーマンス最適化の1つは、事前フォーマット属性のサポートです。 Logger.Withメソッドはキーと値のペアをAttrsに変換し、Handler.WithAttrsを呼び出します。 ハンドラーは、後でHandleメソッドで使用するために属性を保存することもできますし、Handleへの呼び出しごとに繰り返し行うのではなく、一度だけ属性をフォーマットする機会を取ることもできます。

The signature of the `WithAttrs` method is

> `WithAttrs` メソッドのシグネチャは以下です。


    WithAttrs(attrs []Attr) Handler

The attributes are the processed key-value pairs passed to `Logger.With`.
The return value should be a new instance of your handler that contains
the attributes, possibly pre-formatted.

> 属性は、Logger.With に渡される、処理されたキーと値のペアです。戻り値は、属性を含むハンドラーの新しいインスタンス (事前にフォーマットされている可能性があります) である必要があります。

`WithAttrs` must return a new handler with the additional attributes, leaving
the original handler (its receiver) unchanged. For example, this call:

> WithAttrsは、元のハンドラー（そのレシーバー）は変更せずに、追加属性を持つ新しいハンドラーを返さなければならない。 例えば、この呼び出しは


    logger2 := logger1.With("k", v)

creates a new logger, `logger2`, with an additional attribute, but has no
effect on `logger1`.

> 追加の属性を持つ新しいロガー、logger2 を作成しますが、logger1 には影響しません。

We will show example implementations of `WithAttrs` below, when we discuss `WithGroup`.

> 後でWithAttrsの実装例について示しますが、まずWithGroupについて考えます。

## The `WithGroup` method

`Logger.WithGroup` calls `Handler.WithGroup` directly, with the same
argument, the group name.
A handler should remember the name so it can use it to qualify all subsequent
attributes.

> Logger.WithGroupは、同じ引数、グループ名で、Handler.WithGroupを直接呼び出します。 ハンドラーはこの名前を覚えておく必要があり、後続のすべての属性を修飾するために使用できます。

The signature of `WithGroup` is:

> `WithGroup` のシグネチャは


    WithGroup(name string) Handler

Like `WithAttrs`, the `WithGroup` method should return a new handler, not modify
the receiver.

> `WithAttrs` と同様に、`WithGroup` メソッドは新たなハンドラを作成して返す必要があり、レシーバーは変更してはいけません。

The implementations of `WithGroup` and `WithAttrs` are intertwined.
Consider this statement:

> WithGroupとWithAttrsの実装は絡み合っています。 次の文を考えてみましょう。


    logger = logger.WithGroup("g1").With("k1", 1).WithGroup("g2").With("k2", 2)

Subsequent `logger` output should qualify key "k1" with group "g1",
and key "k2" with groups "g1" and "g2".
The order of the `Logger.WithGroup` and `Logger.With` calls must be respected by
the implementations of `Handler.WithGroup` and `Handler.WithAttrs`.

> 後続のロガー出力は、グループ「g1」でキー「k1」を、グループ「g1」と「g2」でキー「k2」を修飾する必要があります。 Logger.WithGroupとLogger.With呼び出しの順序は、Handler.WithGroupとHandler.WithAttrsの実装によって尊重されなければなりません。

We will look at two implementations of `WithGroup` and `WithAttrs`, one that pre-formats and
one that doesn't.

> WithGroup と WithAttrs の 2 つの実装を見ていきます。1 つは事前フォーマットするもの、もう 1 つは事前フォーマットしないものです。

### Without pre-formatting

Our first implementation will collect the information from `WithGroup` and
`WithAttrs` calls to build up a slice of group names and attribute lists,
and loop over that slice in `Handle`. We start with a struct that can hold
either a group name or some attributes:

> 最初の実装では、WithGroupとWithAttrsの呼び出しから情報を収集し、グループ名と属性リストのスライスを構築し、Handleでそのスライスをループします。 まず、グループ名か属性を保持できる構造体から始めます：

```
// groupOrAttrs holds either a group name or a list of slog.Attrs.
type groupOrAttrs struct {
	group string      // group name if non-empty
	attrs []slog.Attr // attrs if non-empty
}
```

Then we add a slice of `groupOrAttrs` to our handler:

> 私たちのハンドラに `groupOrAttrs` のスライスを追加します。

```
type IndentHandler struct {
	opts Options
	goas []groupOrAttrs
	mu   *sync.Mutex
	out  io.Writer
}
```

As stated above, The `WithGroup` and `WithAttrs` methods should not modify their
receiver.
To that end, we define a method that will copy our handler struct
and append one `groupOrAttrs` to the copy:

> 前述したように、WithGroupメソッドとWithAttrsメソッドはレシーバーを変更すべきではありません。 そのため、ハンドラ構造体をコピーし、そのコピーにgroupOrAttrsを1つ追加するメソッドを定義します：

```
func (h *IndentHandler) withGroupOrAttrs(goa groupOrAttrs) *IndentHandler {
	h2 := *h
	h2.goas = make([]groupOrAttrs, len(h.goas)+1)
	copy(h2.goas, h.goas)
	h2.goas[len(h2.goas)-1] = goa
	return &h2
}
```

Most of the fields of `IndentHandler` can be copied shallowly, but the slice of
`groupOrAttrs` requires a deep copy, or the clone and the original will point to
the same underlying array. If we used `append` instead of making an explicit
copy, we would introduce that subtle aliasing bug.

> `IndentHandler`のほとんどのフィールドは浅くコピーできますが、`groupOrAttrs`のスライスは深くコピーする必要があります。 明示的なコピーの代わりにappendを使うと、微妙なエイリアシングのバグが導入されてしまいます。

Using `withGroupOrAttrs`, the `With` methods are easy:

> withGroupOrAttrsを使えば、Withメソッドは簡単です:

```
func (h *IndentHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	return h.withGroupOrAttrs(groupOrAttrs{group: name})
}

func (h *IndentHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	return h.withGroupOrAttrs(groupOrAttrs{attrs: attrs})
}
```

The `Handle` method can now process the groupOrAttrs slice after
the built-in attributes and before the ones in the record:

> `Handle`メソッドは、`groupOrAttrs`スライスを、組み込み属性の後、レコード内の属性の前に処理できるようになりました:

```
func (h *IndentHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := make([]byte, 0, 1024)
	if !r.Time.IsZero() {
		buf = h.appendAttr(buf, slog.Time(slog.TimeKey, r.Time), 0)
	}
	buf = h.appendAttr(buf, slog.Any(slog.LevelKey, r.Level), 0)
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		buf = h.appendAttr(buf, slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", f.File, f.Line)), 0)
	}
	buf = h.appendAttr(buf, slog.String(slog.MessageKey, r.Message), 0)
	indentLevel := 0
	// Handle state from WithGroup and WithAttrs.
	goas := h.goas
	if r.NumAttrs() == 0 {
		// If the record has no Attrs, remove groups at the end of the list; they are empty.
		for len(goas) > 0 && goas[len(goas)-1].group != "" {
			goas = goas[:len(goas)-1]
		}
	}
	for _, goa := range goas {
		if goa.group != "" {
			buf = fmt.Appendf(buf, "%*s%s:\n", indentLevel*4, "", goa.group)
			indentLevel++
		} else {
			for _, a := range goa.attrs {
				buf = h.appendAttr(buf, a, indentLevel)
			}
		}
	}
	r.Attrs(func(a slog.Attr) bool {
		buf = h.appendAttr(buf, a, indentLevel)
		return true
	})
	buf = append(buf, "---\n"...)
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}
```

You may have noticed that our algorithm for
recording `WithGroup` and `WithAttrs` information is quadratic in the
number of calls to those methods, because of the repeated copying.
That is unlikely to matter in practice, but if it bothers you,
you can use a linked list instead,
which `Handle` will have to reverse or visit recursively.
See the
[github.com/jba/slog/withsupport](https://github.com/jba/slog/tree/main/withsupport)
package for an implementation.

> WithGroupとWithAttrsの情報を記録するアルゴリズムは、コピーを繰り返すため、これらのメソッドの呼び出し回数の2次関数になることにお気づきでしょうか。 しかし、もしそれが気になるのであれば、リンクリストを代わりに使うことができます。 実装は[github.com/jba/slog/withsupport](https://github.com/jba/slog/tree/main/withsupport)パッケージを参照してください。

#### Getting the mutex right

Let us revisit the last few lines of `Handle`:

> ハンドルの最後の数行をもう一度見てみましょう：

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
    return err

This code hasn't changed, but we can now appreciate why `h.mu` is a
pointer to a `sync.Mutex`. Both `WithGroup` and `WithAttrs` copy the handler.
Both copies point to the same mutex.
If the copy and the original used different mutexes and were used concurrently,
then their output could be interleaved, or some output could be lost.
Code like this:

> このコードに変更はないが、なぜh.muがsync.Mutexへのポインタなのかが理解できるようになった。 WithGroupもWithAttrsもハンドラをコピーする。 どちらのコピーも同じミューテックスを指す。 コピーとオリジナルが異なるミューテックスを使用し、同時に使用された場合、それらの出力がインターリーブされたり、一部の出力が失われたりする可能性がある。 こんなコードだ：


    l2 := l1.With("a", 1)
    go l1.Info("hello")
    l2.Info("goodbye")

could produce output like this:

> このような出力が得られる：

    hegoollo a=dbye1

See [this bug report](https://go.dev/issue/61321) for more detail.

> 詳しくは [this bug report](https://go.dev/issue/61321) を参照してください。

### With pre-formatting

Our second implementation implements pre-formatting.
This implementation is more complicated than the previous one.
Is the extra complexity worth it?
That depends on your circumstances, but here is one circumstance where
it might be.
Say that you wanted your server to log a lot of information about an incoming
request with every log message that happens during that request. A typical
handler might look something like this:

> 私たちの2番目の実装は、プリフォーマッティングを実装している。 この実装は前のものよりも複雑だ。 その複雑さに見合うだけの価値があるのだろうか？ それはあなたの状況次第ですが、ここに一つの状況があります。 例えば、入ってきたリクエストに関する多くの情報を、そのリクエスト中に発生する すべてのログメッセージに記録させたいとします。 典型的なハンドラは次のようなものです：

    func (s *Server) handleWidgets(w http.ResponseWriter, r *http.Request) {
        logger := s.logger.With(
            "url", r.URL,
            "traceID": r.Header.Get("X-Cloud-Trace-Context"),
            // many other attributes
            )
        // ...
    }

A single `handleWidgets` request might generate hundreds of log lines.
For instance, it might contain code like this:

> 一つのhandleWidgetsリクエストが何百行ものログを生成するかもしれません。 例えば、次のようなコードが含まれるかもしれません：

    for _, w := range widgets {
        logger.Info("processing widget", "name", w.Name)
        // ...
    }

For every such line, the `Handle` method we wrote above will format all
the attributes that were added using `With` above, in addition to the
ones on the log line itself.

> そのような行ごとに、上で書いたHandleメソッドは、ログ行自体の属性に加えて、上でWithを使って追加されたすべての属性をフォーマットする。

Maybe all that extra work doesn't slow down your server significantly, because
it does so much other work that time spent logging is just noise.
But perhaps your server is fast enough that all that extra formatting appears high up
in your CPU profiles. That is when pre-formatting can make a big difference,
by formatting the attributes in a call to `With` just once.

> ロギングに費やす時間はノイズに過ぎません。しかし、恐らくあなたのサーバーは十分に速く、余分なフォーマットがCPUプロファイルの上位に表示されるでしょう。 そのような場合、`With`への呼び出しで属性を一度だけフォーマットすることで、事前フォーマットが大きな違いを生むことがあります。

To pre-format the arguments to `WithAttrs`, we need to keep track of some
additional state in the `IndentHandler` struct.

> `WithAttrs`の引数を事前にフォーマットするために、`IndentHandler`構造体で追加の状態を追跡する必要がある。

```
type IndentHandler struct {
	opts           Options
	preformatted   []byte   // data from WithGroup and WithAttrs
	unopenedGroups []string // groups from WithGroup that haven't been opened
	indentLevel    int      // same as number of opened groups so far
	mu             *sync.Mutex
	out            io.Writer
}
```

Mainly, we need a buffer to hold the pre-formatted data.
But we also need to keep track of which groups
we've seen but haven't output yet. We'll call those groups "unopened."
We also need to track how many groups we've opened, which we can do
with a simple counter, since an opened group's only effect is to change the
indentation level.

> 主に、フォーマット済みのデータを保持するためのバッファが必要です。 しかし、まだ出力していないグループを追跡する必要もあります。 これらのグループを "未開封 "と呼ぶことにします。 また、いくつのグループを開いたかを追跡する必要があるが、開いたグループの効果はインデントレベルを変更することだけなので、単純なカウンターで出来ます。

The `WithGroup` implementation is a lot like the previous one: just remember the
new group, which is unopened initially.

> WithGroupの実装は、前のものとよく似ています。最初は未開封の新しいグループを覚えておくだけです。

```
func (h *IndentHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	h2 := *h
	// Add an unopened group to h2 without modifying h.
	h2.unopenedGroups = make([]string, len(h.unopenedGroups)+1)
	copy(h2.unopenedGroups, h.unopenedGroups)
	h2.unopenedGroups[len(h2.unopenedGroups)-1] = name
	return &h2
}
```

`WithAttrs` does all the pre-formatting:

> WithAttrsは事前フォーマットの処理をすべて行う：

```
func (h *IndentHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	h2 := *h
	// Force an append to copy the underlying array.
	pre := slices.Clip(h.preformatted)
	// Add all groups from WithGroup that haven't already been added.
	h2.preformatted = h2.appendUnopenedGroups(pre, h2.indentLevel)
	// Each of those groups increased the indent level by 1.
	h2.indentLevel += len(h2.unopenedGroups)
	// Now all groups have been opened.
	h2.unopenedGroups = nil
	// Pre-format the attributes.
	for _, a := range attrs {
		h2.preformatted = h2.appendAttr(h2.preformatted, a, h2.indentLevel)
	}
	return &h2
}

func (h *IndentHandler) appendUnopenedGroups(buf []byte, indentLevel int) []byte {
	for _, g := range h.unopenedGroups {
		buf = fmt.Appendf(buf, "%*s%s:\n", indentLevel*4, "", g)
		indentLevel++
	}
	return buf
}
```

It first opens any unopened groups. This handles calls like:

> まず未オープンのグループを開きます。 これは次のような呼び出しに対応します：

    logger.WithGroup("g").WithGroup("h").With("a", 1)

Here, `WithAttrs` must output "g" and "h" before "a". Since a group established
by `WithGroup` is in effect for the rest of the log line, `WithAttrs` increments
the indentation level for each group it opens.

> ここで、`WithAttrs`は "a "の前に "g "と "h "を出力しなければならない。 `WithGroup`によって確立されたグループは、ログ行の残りの部分に対して有効であるため、`WithAttrs`は、開くグループごとにインデントレベルを増加させる。

Lastly, `WithAttrs` formats its argument attributes, using the same `appendAttr`
method we saw above.

> 最後に、`WithAttrs`は、上で見たのと同じ`appendAttr`メソッドを使って、引数の属性をフォーマットする。

It's the `Handle` method's job to insert the pre-formatted material in the right
place, which is after the built-in attributes and before the ones in the record:

> あらかじめフォーマットされた素材を適切な場所に挿入するのは`Handle`メソッドの仕事であり、それは組み込み属性の後でレコード内の属性の前である：

```
func (h *IndentHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := make([]byte, 0, 1024)
	if !r.Time.IsZero() {
		buf = h.appendAttr(buf, slog.Time(slog.TimeKey, r.Time), 0)
	}
	buf = h.appendAttr(buf, slog.Any(slog.LevelKey, r.Level), 0)
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		buf = h.appendAttr(buf, slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", f.File, f.Line)), 0)
	}
	buf = h.appendAttr(buf, slog.String(slog.MessageKey, r.Message), 0)
	// Insert preformatted attributes just after built-in ones.
	buf = append(buf, h.preformatted...)
	if r.NumAttrs() > 0 {
		buf = h.appendUnopenedGroups(buf, h.indentLevel)
		r.Attrs(func(a slog.Attr) bool {
			buf = h.appendAttr(buf, a, h.indentLevel+len(h.unopenedGroups))
			return true
		})
	}
	buf = append(buf, "---\n"...)
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}
```

It must also open any groups that haven't yet been opened. The logic covers
log lines like this one:

> また、まだ開いていないグループも開かなければならない。 ロジックは以下のようなログ呼び出しをカバーする：


    logger.WithGroup("g").Info("msg", "a", 1)

where "g" is unopened before `Handle` is called and must be written to produce
the correct output:

> ここで "g "は`Handle`が呼び出される前は未開封であり、正しい出力を得るためには書き込まれなければならない：

    level: INFO
    msg: "msg"
    g:
        a: 1

The check for `r.NumAttrs() > 0` handles this case:

> この場合、`r.NumAttrs() > 0` のチェックで対応する：

    logger.WithGroup("g").Info("msg")

Here there are no record attributes, so no group to open.

> ここではレコード属性がないので、開くべきグループはない。

## Testing

The [`Handler` contract](https://pkg.go.dev/log/slog#Handler) specifies several
constraints on handlers.
To verify that your handler follows these rules and generally produces proper
output, use the [testing/slogtest package](https://pkg.go.dev/testing/slogtest).

That package's `TestHandler` function takes an instance of your handler and
a function that returns its output formatted as a slice of maps. Here is the test function
for our example handler:

```
func TestSlogtest(t *testing.T) {
	var buf bytes.Buffer
	err := slogtest.TestHandler(New(&buf, nil), func() []map[string]any {
		return parseLogEntries(t, buf.Bytes())
	})
	if err != nil {
		t.Error(err)
	}
}
```

Calling `TestHandler` is easy. The hard part is parsing your handler's output.
`TestHandler` calls your handler multiple times, resulting in a sequence of log
entries.
It is your job to parse each entry into a `map[string]any`.
A group in an entry should appear as a nested map.

If your handler outputs a standard format, you can use an existing parser.
For example, if your handler outputs one JSON object per line, then you
can split the output into lines and call `encoding/json.Unmarshal` on each.
Parsers for other formats that can unmarshal into a map can be used out
of the box.
Our example output is enough like YAML so that we can use the `gopkg.in/yaml.v3`
package to parse it:

```
func parseLogEntries(t *testing.T, data []byte) []map[string]any {
	entries := bytes.Split(data, []byte("---\n"))
	entries = entries[:len(entries)-1] // last one is empty
	var ms []map[string]any
	for _, e := range entries {
		var m map[string]any
		if err := yaml.Unmarshal([]byte(e), &m); err != nil {
			t.Fatal(err)
		}
		ms = append(ms, m)
	}
	return ms
}
```

If you have to write your own parser, it can be far from perfect.
The `slogtest` package uses only a handful of simple attributes.
(It is testing handler conformance, not parsing.)
Your parser can ignore edge cases like whitespace and newlines in keys and
values. Before switching to a YAML parser, we wrote an adequate custom parser
in 65 lines.

# General considerations

## Copying records

Most handlers won't need to copy the `slog.Record` that is passed
to the `Handle` method.
Those that do must take special care in some cases.

A handler can make a single copy of a `Record` with an ordinary Go
assignment, channel send or function call if it doesn't retain the
original.
But if its actions result in more than one copy, it should call `Record.Clone`
to make the copies so that they don't share state.
This `Handle` method passes the record to a single handler, so it doesn't require `Clone`:

    type Handler1 struct {
        h slog.Handler
        // ...
    }

    func (h *Handler1) Handle(ctx context.Context, r slog.Record) error {
        return h.h.Handle(ctx, r)
    }

This `Handle` method might pass the record to more than one handler, so it
should use `Clone`:

    type Handler2 struct {
        hs []slog.Handler
        // ...
    }

    func (h *Handler2) Handle(ctx context.Context, r slog.Record) error {
        for _, hh := range h.hs {
            if err := hh.Handle(ctx, r.Clone()); err != nil {
                return err
            }
        }
        return nil
    }

## Concurrency safety

A handler must work properly when a single `Logger` is shared among several
goroutines.
That means that mutable state must be protected with a lock or some other mechanism.
In practice, this is not hard to achieve, because many handlers won't have any
mutable state.

- The `Enabled` method typically consults only its arguments and a configured
  level. The level is often either set once initially, or is held in a
  `LevelVar`, which is already concurrency-safe.

- The `WithAttrs` and `WithGroup` methods should not modify the receiver,
  for reasons discussed above.

- The `Handle` method typically works only with its arguments and stored fields.

Calls to output methods like `io.Writer.Write` should be synchronized unless
it can be verified that no locking is needed.
As we saw in our example, storing a pointer to a mutex enables a logger and
all of its clones to synchronize with each other.
Beware of facile claims like "Unix writes are atomic"; the situation is a lot more nuanced than that.

Some handlers have legitimate reasons for keeping state.
For example, a handler might support a `SetLevel` method to change its configured level
dynamically.
Or it might output the time between successive calls to `Handle`,
which requires a mutable field holding the last output time.
Synchronize all accesses to such fields, both reads and writes.

The built-in handlers have no directly mutable state.
They use a mutex only to sequence calls to their contained `io.Writer`.

## Robustness

Logging is often the debugging technique of last resort. When it is difficult or
impossible to inspect a system, as is typically the case with a production
server, logs provide the most detailed way to understand its behavior.
Therefore, your handler should be robust to bad input.

For example, the usual advice when when a function discovers a problem,
like an invalid argument, is to panic or return an error.
The built-in handlers do not follow that advice.
Few things are more frustrating than being unable to debug a problem that
causes logging to fail;
it is better to produce some output, however imperfect, than to produce none at all.
That is why methods like `Logger.Info` convert programming bugs in their list of
key-value pairs, like missing values or malformed keys,
into `Attr`s that contain as much information as possible.

One place to avoid panics is in processing attribute values. A handler that wants
to format a value will probably switch on the value's kind:

    switch attr.Value.Kind() {
    case KindString: ...
    case KindTime: ...
    // all other Kinds
    default: ...
    }

What should happen in the default case, when the handler encounters a `Kind`
that it doesn't know about?
The built-in handlers try to muddle through by using the result of the value's
`String` method, as our example handler does.
They do not panic or return an error.
Your own handlers might in addition want to report the problem through your production monitoring
or error-tracking telemetry system.
The most likely explanation for the issue is that a newer version of the `slog` package added
a new `Kind`&mdash;a backwards-compatible change under the Go 1 Compatibility
Promise&mdash;and the handler wasn't updated.
That is certainly a problem, but it shouldn't deprive
readers from seeing the rest of the log output.

There is one circumstance where returning an error from `Handler.Handle` is appropriate.
If the output operation itself fails, the best course of action is to report
this failure by returning the error. For instance, the last two lines of the
built-in `Handle` methods are

    _, err := h.w.Write(*state.buf)
    return err

Although the output methods of `Logger` ignore the error, one could write a
handler that does something with it, perhaps falling back to writing to standard
error.

## Speed

Most programs don't need fast logging.
Before making your handler fast, gather data&mdash;preferably production data,
not benchmark comparisons&mdash;that demonstrates that it needs to be fast.
Avoid premature optimization.

If you need a fast handler, start with pre-formatting. It may provide dramatic
speed-ups in cases where a single call to `Logger.With` is followed by many
calls to the resulting logger.

If log output is the bottleneck, consider making your handler asynchronous.
Do the minimal amount of processing in the handler, then send the record and
other information over a channel. Another goroutine can collect the incoming log
entries and write them in bulk and in the background.
You might want to preserve the option to log synchronously
so you can see all the log output to debug a crash.

Allocation is often a major cause of a slow system.
The `slog` package already works hard at minimizing allocations.
If your handler does its own allocation, and profiling shows it to be
a problem, then see if you can minimize it.

One simple change you can make is to replace calls to `fmt.Sprintf` or `fmt.Appendf`
with direct appends to the buffer. For example, our IndentHandler appends string
attributes to the buffer like so:

	buf = fmt.Appendf(buf, "%s: %q\n", a.Key, a.Value.String())

As of Go 1.21, that results in two allocations, one for each argument passed to
an `any` parameter. We can get that down to zero by using `append` directly:

	buf = append(buf, a.Key...)
	buf = append(buf, ": "...)
	buf = strconv.AppendQuote(buf, a.Value.String())
	buf = append(buf, '\n')

Another worthwhile change is to use a `sync.Pool` to manage the one chunk of
memory that most handlers need:
the `[]byte` buffer holding the formatted output.

Our example `Handle` method began with this line:

	buf := make([]byte, 0, 1024)

As we said above, providing a large initial capacity avoids repeated copying and
re-allocation of the slice as it grows, reducing the number of allocations to
one.
But we can get it down to zero in the steady state by keeping a global pool of buffers.
Initially, the pool will be empty and new buffers will be allocated.
But eventually, assuming the number of concurrent log calls reaches a steady
maximum, there will be enough buffers in the pool to share among all the
ongoing `Handler` calls. As long as no log entry grows past a buffer's capacity,
there will be no allocations from the garbage collector's point of view.

We will hide our pool behind a pair of functions, `allocBuf` and `freeBuf`.
The single line to get a buffer at the top of `Handle` becomes two lines:

	bufp := allocBuf()
	defer freeBuf(bufp)

One of the subtleties involved in making a `sync.Pool` of slices
is suggested by the variable name `bufp`: your pool must deal in
_pointers_ to slices, not the slices themselves.
Pooled values must always be pointers. If they aren't, then the `any` arguments
and return values of the `sync.Pool` methods will themselves cause allocations,
defeating the purpose of pooling.

There are two ways to proceed with our slice pointer: we can replace `buf`
with `*bufp` throughout our function, or we can dereference it and remember to
re-assign it before freeing:

	bufp := allocBuf()
	buf := *bufp
	defer func() {
		*bufp = buf
		freeBuf(bufp)
	}()


Here is our pool and its associated functions:

```
var bufPool = sync.Pool{
	New: func() any {
		b := make([]byte, 0, 1024)
		return &b
	},
}

func allocBuf() *[]byte {
	return bufPool.Get().(*[]byte)
}

func freeBuf(b *[]byte) {
	// To reduce peak allocation, return only smaller buffers to the pool.
	const maxBufferSize = 16 << 10
	if cap(*b) <= maxBufferSize {
		*b = (*b)[:0]
		bufPool.Put(b)
	}
}
```

The pool's `New` function does the same thing as the original code:
create a byte slice with 0 length and plenty of capacity.
The `allocBuf` function just type-asserts the result of the pool's
`Get` method.

The `freeBuf` method truncates the buffer before putting it back
in the pool, so that `allocBuf` always returns zero-length slices.
It also implements an important optimization: it doesn't return
large buffers to the pool.
To see why this important, consider what would happen if there were a single,
unusually large log entry&mdash;say one that was a megabyte when formatted.
If that megabyte-sized buffer were put in the pool, it could remain
there indefinitely, constantly being reused, but with most of its capacity
wasted.
The extra memory might never be used again by the handler, and since it was in
the handler's pool, it might never be garbage-collected for reuse elsewhere.
We can avoid that situation by excluding large buffers from the pool.