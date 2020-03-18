package effectivego25

import (
	"strings"
	"time"

	"github.com/devlights/gomy/output"
)

// ShareByCommunicating -- Effective Go - Concurrency の 内容についてのサンプルです。
func ShareByCommunicating() error {
	/*
		https://golang.org/doc/effective_go.html#sharing

		非同期処理のやり方には大きく以下の２つの方法がある。

		(1) 共有メモリを置いて各非同期処理がそれを処理する方法
			- Shared-memory communication
			- スレッド間でメモリを共有するスタイル
			- 発想としてはシンプルなので、このやり方で処理を作る人も多い
				- 仕組みを作りやすい反面、データをきっちりと守るのは難しく、ちょっとしたことでデータ競合などが起こせてしまう
		(2) 共有メモリを置かずにスレッド間でメッセージをパッシングする方法
			- Message-passing communication
			- 要は、共有メモリを使わずにスレッド間でデータをやり取りして処理する方法
			- Goが推奨しているのはこちらの方式
				- チャネルはこのために存在している
			- データを共有していないので、ある瞬間で一つのデータを操作するのはただ一つの処理のみとなる

		Goでは非同期処理を行う上で以下の有名なスローガンを掲げている。

		> Do not communicate by sharing memory; instead, share memory by communicating.

		>> メモリを共有して通信しないでください。代わりに、通信してメモリを共有します。

		データは一度にたった一つのゴルーチンのみがアクセスできるべきであり、この設計上、データの競合は発生しない。
		概念的に、Unixのパイプラインに近いイメージで、それをタイプセーフにした感じ。以下のように記載されている。

		> One way to think about this model is to consider a typical single-threaded program running on one CPU.
		> It has no need for synchronization primitives.
		> Now run another such instance; it too needs no synchronization.
		> Now let those two communicate; if the communication is the synchronizer, there's still no need for other synchronization.
		> Unix pipelines, for example, fit this model perfectly.
		> Although Go's approach to concurrency originates in Hoare's Communicating Sequential Processes (CSP),
		> it can also be seen as a type-safe generalization of Unix pipes.

		>> このモデルについて考える1つの方法は、1つのCPUで実行される典型的なシングルスレッドプログラムを検討することです。
		>> 同期プリミティブは必要ありません。
		>> 次に、別のそのようなインスタンスを実行します。 同期も不要です。
		>> 次に、これら2つの通信を行います。 通信がシンクロナイザーである場合、他の同期の必要はありません。
		>> たとえば、Unixパイプラインはこのモデルに完全に適合します。
		>> Goの並行性へのアプローチはHoareのCommunicating Sequential Processes（CSP）に由来していますが、
		>> Unixパイプのタイプセーフな一般化として見ることもできます。

		- ゴルーチンは「同じアドレス空間」で他のゴルーチンと平行して実行することが出来る機能
		- ゴルーチンを実行するには、関数の呼び出しに `go` というキーワードを付与するだけ
		- ゴルーチンで呼び出された処理が完了するとゴルーチンは静かに終了する
			- シェルの `&` 実行みたいなイメージ
			- 完了待機などはユーザ側で組み込まないといけない
	*/

	// 以下は同期の呼び出し
	output.Stdoutl("sync call", "helloworld")
	// 同じ呼び出しをゴルーチン化するには `go` キーワードを付与する
	go output.Stdoutl("async call", "helloworld goroutine")
	// ゴルーチンを利用する際は、クロージャを多用することが多い
	message := "helloworld gotourine message"
	go func() {
		m := strings.ToUpper(message)
		output.Stdoutl("closure", m)
	}()

	// そのままだと、平行して実行される前にメインゴルーチンが終了してしまうので
	// 意図的に待つ。本来、このような待機でスリープを利用するのはいいやり方ではなく
	// チャネルや他の待機するための構造を利用するのがいいが、それは後のサンプルで記述する。
	time.Sleep(100 * time.Millisecond)

	/*
		上記のサンプルを最後の time.Sleep(100 * time.Millisecond) がない状態で実行すると
		大抵以下のようになる。（待機するための仕組みを実装していないので、このように表示されないときもある。)

			ENTER EXAMPLE NAME: effectivego_25_concurrency_share_by_communicating
			[Name] "effectivego_25_concurrency_share_by_communicating"
			sync call            helloworld


			[Elapsed] 0s
			async call           helloworld goroutine
			closure              HELLOWORLD GOTOURINE MESSAGE

		サンプル実行が完了した後で、ゴルーチンの出力が出ている。
		time.Sleep() を入れた場合は、大抵以下のようになる。
		（待機するための仕組みを実装していないので、このように表示されないときもある。)

			ENTER EXAMPLE NAME: effectivego_25_concurrency_share_by_communicating
			[Name] "effectivego_25_concurrency_share_by_communicating"
			sync call            helloworld
			closure              HELLOWORLD GOTOURINE MESSAGE
			async call           helloworld goroutine


			[Elapsed] 100.4088ms
	*/

	return nil
}
