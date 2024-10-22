package main

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type (
	// counterHandler は、内部でログ出力回数をカウントし、その内容を出力するカスタムハンドラです。
	// ログ出力自体は、作成時に指定した元ハンドラに任せます。
	//
	// slogでカスタムハンドラを作成する場合, `slog.Handler` インターフェースを満たす必要があります。
	// `slog.Handler` インターフェースには以下のメソッドが定義されています。(https://pkg.go.dev/log/slog@go1.23.2#Handler)
	//
	// 	- Enabled(context.Context, Level) bool
	// 	- Handle(context.Context, Record) error
	// 	- WithAttrs(attrs []Attr) Handler
	// 	- WithGroup(name string) Handler
	//
	// コアとなる処理は `slog.Handler.Handle` メソッドです。
	// `go doc` にある通り、独自実装する場合は以下のルールに従うべきであると記載されています。
	//
	// 	- If r.Time is the zero time, ignore the time.(r.Timeがゼロ時刻の場合は、時刻を無視する)
	// 	- If r.PC is zero, ignore it.(r.PCがゼロの場合、無視する)
	// 	- Attr's values should be resolved.(Attrの値は解決されるべき)
	// 	- If an Attr's key and value are both the zero value, ignore the Attr. This can be tested with attr.Equal(Attr{}).(Attrのkeyとvalueの両方がゼロの場合、そのAttrは無視される。これは attr.Equal(Attr{}) でテスト可能)
	// 	- If a group's key is empty, inline the group's Attrs.(グループのキーが空の場合、グループのAttrsをインライン化する)
	// 	- If a group has no Attrs (even if it has a non-empty key), ignore it.(グループにAttrがない場合は (空でないキーがある場合でも)、無視する)
	//
	// 詳細な解説については [A Guide to Writing slog Handlers](https://github.com/golang/example/blob/master/slog-handler-guide/README.md) を参照。
	// (日本語訳を添えてあるものを [slog-handler-guide-ja.md](./slog-handler-guide-ja.md) として置いています。)
	//
	// 現実的に、フルカスタムのハンドラをゼロから自前で作成することはほぼ無いと思います。
	// 基本は、元から存在している何かのハンドラ(slog.JSONHandlerなど)を元にして、その上に独自の処理を付け加えることになります。
	//
	// その場合、ログ出力の部分は元ハンドラに移譲する形で完了するが、実装時に以下の点には注意する必要がある。
	//
	// 	1. `WithAttrs`と`WithGroup`メソッドを上書きでメソッド定義しておかないと、カスタムハンドラをハンドラとして生成したロガーにて `Logger.With()` または `Logger.WithGroup()` を使ってロガーを複製した場合に属性が引き継がれない状態が発生する。
	// 	2. 内部で何らかの状態を保持する必要がある場合は、複数のゴルーチンにて同時に呼び出されることを考慮してロックする。
	// 	3. 2の件に加えて、`slog.Logger.With()` または `slog.Logger.WithGroup()` が利用されることが想定される場合は値やミューテックスをポインタで持っておく。（これらのメソッドではハンドラをコピーするため）
	//
	counterHandler struct {
		slog.Handler             // 内部で利用する元ハンドラ
		count        *int        // カウンタ
		mu           *sync.Mutex // ロック用
	}
)

func newHandler(handler slog.Handler) *counterHandler {
	var (
		count = 0
		mu    sync.Mutex
	)
	return &counterHandler{Handler: handler, count: &count, mu: &mu}
}

func (me *counterHandler) Handle(ctx context.Context, r slog.Record) error {
	me.mu.Lock()
	defer me.mu.Unlock()

	//
	// ミリ秒と呼び出し回数を出力
	//
	r.AddAttrs(slog.String("millis", time.Now().Format(".000")))
	*me.count++
	r.AddAttrs(slog.Int("count", *me.count))

	return me.Handler.Handle(ctx, r)
}

func (me *counterHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &counterHandler{Handler: me.Handler.WithAttrs(attrs), count: me.count, mu: me.mu}
}

func (me *counterHandler) WithGroup(name string) slog.Handler {
	return &counterHandler{Handler: me.Handler.WithGroup(name), count: me.count, mu: me.mu}
}

func main() {
	var (
		level = &slog.LevelVar{}
		opt   = &slog.HandlerOptions{
			Level:       level,
			ReplaceAttr: replaceAttr,
		}
		writer  = os.Stderr
		handler = newHandler(slog.NewTextHandler(writer, opt))
		logger  = slog.New(handler)

		wg             sync.WaitGroup
		goroutineCount = runtime.NumCPU() * 2
		loopCount      = 2
	)
	logger.Info("Start", "NumCPU", runtime.NumCPU(), "goroutineCount", goroutineCount)

	//
	// 複数のゴルーチンから並行してログ出力
	//
	wg.Add(goroutineCount)
	for i := range goroutineCount {
		go func(logger *slog.Logger) {
			defer wg.Done()
			for i := range loopCount {
				logger.Info(strconv.Itoa(i))
			}
		}(logger.With("goroutine", i))
	}

	wg.Wait()
}

func replaceAttr(g []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey {
		return slog.Attr{}
	}

	return a
}
