package tutorial

import (
	"fmt"
	"sync"
	"time"
)

func GoTourGoroutine() error {
	// ------------------------------------------------------------
	// goroutine (ゴルーチン) は、Goのランタイムで管理される軽量スレッドのこと。
	// pythonのコルーチンと同じイメージを持ったほうが分かりやすい。
	// なので一時停止及び再開可能。停止と再開はGoのランタイム側で管理される。
	//
	// Go言語は、言語のコア機能の一部として並行処理機能を提供している。
	// goroutine は、OSスレッドではないので、とても軽量。goroutineを一つ起動するのに
	// かかるメモリフットプリントは 数KB 程度の模様。
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
	// 非同期処理は、現実問題投げっぱなしというシチュエーションはあまりなく
	// なんらかの形で終了を待機する必要がある。goroutine を使うだけでは
	// 非同期にした関数を投げっぱなしにしてしまうので、sync.WaitGroupを
	// 利用するのが多い。
	// ------------------------------------------------------------
	var (
		wg sync.WaitGroup
	)

	run := func(wait time.Duration, prefix string, wg *sync.WaitGroup) {
		if wg != nil {
			defer wg.Done()
		}

		fmt.Printf("[%s] func begin\n", prefix)
		time.Sleep(wait)
		fmt.Printf("[%s] func end\n", prefix)
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
