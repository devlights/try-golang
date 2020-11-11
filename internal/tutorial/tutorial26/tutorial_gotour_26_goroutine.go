package tutorial26

import (
	"fmt"
	"sync"
	"time"
)

// Goroutine は、 Tour of Go - Goroutines (https://tour.golang.org/concurrency/1) の サンプルです。
func Goroutine() error {
	// ------------------------------------------------------------
	// goroutine (ゴルーチン) は、Goのランタイムで管理される軽量スレッドのこと。
	// 軽量スレッドと書いているが、実際はグレーンスレッドではなく
	// pythonのコルーチンと同じイメージを持ったほうが分かりやすい。
	// コルーチンなので、割り込みされることがない。（ノンプリエンプティブ）
	// コルーチンなので、一時停止及び再開可能。
	// 停止と再開はGoのランタイム側で管理される。
	// goroutineは、Goのプログラムでの最も基本的な構成単位である。
	//
	// Go言語は、言語のコア機能の一部として並行処理機能を提供している。
	// goroutine は、OSスレッドではないので、とても軽量。goroutineを一つ起動するのに
	// かかるメモリフットプリントは 数KB 程度とされている。
	// (runtime/stack.go の _StackMin の値は Go1.13 で 2048 と設定されている)
	//
	// 書籍「Go言語による並行処理（オライリー・ジャパン）」でも
	//   "Goの並行処理における哲学は以下のようにまとめられます。
	//    簡潔さを求め、チャネルをできる限り使い、ゴルーチンを湯水のように使いましょう。" (P.35)
	// と記載されており、基本的に起動する非同期処理の数は気にしなくても良いレベル。
	//
	// 通常の関数は、そのままだともちろん「同期」処理となる。
	// goroutine にすることで、「非同期」に出来る。
	//
	// 特定の関数を「非同期」にすることは、Go言語ではとても簡単で
	//   go func1(x, y)
	// と、関数の前に 「go」 とつけるだけである。これで func1 の実行が非同期になる.
	// (つまり、goroutine が生成される)
	//
	// 大事な点として、以下がある。
	//   - func1, x, および y は、実行元（current)の goroutine で評価される
	//   - func1の実行は新しい goroutine で実行される
	//
	// 他の言語と同じ考え方となるが、Goでも実行時には最低でも一つの goroutine が
	// 動作している。それが メインゴルーチン となる。
	//
	// goroutine は、スレッドと同じで同じアドレス空間で実行されるため
	// 共有メモリを利用してアクセスする場合は、必ず同期する必要がある。
	//
	// Go言語では、この点をチャネル (channel) という概念を用いて
	// 扱いやすいようになるようにしてくれている。共有メモリを扱う場合は
	// syncパッケージに関連する型や関数があるので、それらを利用する。
	//
	// 非同期処理を書いていると、複数のチャネルを扱う必要性がどうしても
	// 出てきてしまうが、Goでは、この点を select ステートメントを利用することで
	// 扱いやすいようにできている。
	//
	// - チャネルは goroutine を束ねる糊のようなもの
	// - select は チャネル   を束ねる糊のようなもの
	//
	// と考えるとわかりやすい。
	//
	// なので、Goの非同期処理は
	//   - goroutine
	//   - channel
	//   - select
	// を使って処理を行うのが基本パターン。
	// 必要な場合に、syncパッケージにあるライブラリも利用して処理する。
	//
	// Go言語でも共有メモリを利用することは勿論できるが、推奨される方法は
	// 共有メモリを使わずに、チャネルを利用してデータを通信して行く方法。
	//
	// Do not communicate by sharing memory;
	// instead, share memory by communicating.
	//    -- Effective Go (http://bit.ly/2NUdLRA))
	//
	// Go言語の並行処理の考え方の元となった
	// のは、CSP (http://bit.ly/2NOw8Y5) という理論。
	//
	// REFERENCES::
	//   - https://golang.org/doc/effective_go.html#concurrency
	//   - https://golang.org/doc/faq#Concurrency
	//   - https://golang.org/ref/spec#Go_statements
	//   - https://qiita.com/niconegoto/items/3952d3c53d00fccc363b
	// ------------------------------------------------------------
	// 普通の関数を定義
	f := func(prefix string) {
		fmt.Printf("[%-5s] こんにちわ世界\n", prefix)
	}

	// これをそのまま呼ぶと当然「同期」呼び出し
	f("sync")

	// Goでは、非同期呼び出しするには goroutine を使う
	// 関数を goroutine にするには、go 関数名 とする
	// これだけで、元々普通の関数だったものが非同期呼び出しに変わる
	// (つまり、Goランタイムによって非同期処理されるようにスケジューリングされる）
	go f("async")

	// ------------------------------------------------------------
	// sync.WaitGroup をつかって終了待機
	//   補足：チャネルを使っても同じような事は可能だが複数の非同期処理を
	//        待機する場合は、sync.WaitGroupの方が楽
	// ------------------------------------------------------------
	var (
		wg sync.WaitGroup
	)

	run := func(wait time.Duration, prefix string, wg *sync.WaitGroup) {
		if wg != nil {
			defer wg.Done()
		}

		fmt.Printf("[%-5s] func begin\n", prefix)
		time.Sleep(wait)
		fmt.Printf("[%-5s] func end\n", prefix)
	}

	// 非同期呼び出し
	wg.Add(1)
	go run(2*time.Second, "async", &wg)

	// 同期で呼び出し
	run(1*time.Second, "sync", nil)

	// 非同期処理の終了待機
	wg.Wait()

	return nil
}
