package runtimes

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/devlights/gomy/output"
)

// Gosched -- runtime.Gosched() のサンプルです。
//
// # REFERENCES
//   - https://dev.to/abh1navv/12-common-uses-of-java-streams-1pgk
//   - https://journal.lampetty.net/entry/concurrency-in-go-goroutines
func Gosched() error {
	const (
		numGoroutines = 5
	)

	var (
		wg sync.WaitGroup
		fn = func(wg *sync.WaitGroup, i int, prefix string, enableYield bool) {
			defer wg.Done()

			output.Stderrf(prefix, "hello [%d]\n", i)

			if enableYield {
				runtime.Gosched()
			}

			output.Stderrf(prefix, "world [%d]\n", i)
		}
	)

	//
	// runtime.Gosched() は、他の言語での yield に相当する (C#とかJavaとか)
	// 呼び出すと、他のゴルーチンに処理コンテキストを譲ってもいいよという通知をGoランタイムに教えるイメージ。
	// (かならず、コンテキストがスイッチするわけではない)
	// 「実行権の放棄」とも言ったりする。
	//

	var (
		enableYield = false
	)

	// runtime.Gosched() を呼び出さない版
	//   yield しないので、基本的に処理コンテキストを持っているゴルーチンは
	//   Goランタイムから一時停止などをされない限り、自身の処理を完結させることが出来る。
	//   今回のサンプルで実行している関数は、I/O待機などがないので
	//   出力は各々のゴルーチン毎に出力されていく。
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go fn(&wg, i, fmt.Sprintf("[goroutine-%02d]", i), enableYield)
	}
	wg.Wait()

	output.StdoutHr()

	// runtime.Gosched() を呼び出す版
	//   yield するように要請するので、それぞれのゴルーチンは
	//   可能であれば、自身の処理コンテキストを他のゴルーチンに
	//   譲る挙動をする。なので、ゴルーチンによっては自身の処理を
	//   １回で完結させることが出来ない場合がある。
	//   (runtime.Gosched()の呼び出しで他のゴルーチンにコンテキストが移ることがあるため)
	//   なので、出力は 前半部と後半部に別々の時間軸で出力される場合がある。
	enableYield = true
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go fn(&wg, i, fmt.Sprintf("[goroutine-%02d]", i), enableYield)
	}
	wg.Wait()

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: runtime_gosched

	   [Name] "runtime_gosched"
	   [goroutine-04]       hello [4]
	   [goroutine-04]       world [4]
	   [goroutine-02]       hello [2]
	   [goroutine-02]       world [2]
	   [goroutine-03]       hello [3]
	   [goroutine-03]       world [3]
	   [goroutine-01]       hello [1]
	   [goroutine-01]       world [1]
	   [goroutine-00]       hello [0]
	   [goroutine-00]       world [0]
	   --------------------------------------------------
	   [goroutine-04]       hello [4]
	   [goroutine-00]       hello [0]
	   [goroutine-01]       hello [1]
	   [goroutine-02]       hello [2]
	   [goroutine-03]       hello [3]
	   [goroutine-04]       world [4]
	   [goroutine-00]       world [0]
	   [goroutine-02]       world [2]
	   [goroutine-03]       world [3]
	   [goroutine-01]       world [1]

	   [Elapsed] 486.98µs
	*/
}
