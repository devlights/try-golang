package syncs

import (
	"runtime"
	"sync"

	"github.com/devlights/gomy/output"
)

// UseOnce は、sync.Onceのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/sync@go1.21.3#Once
func UseOnce() error {
	//
	// sync.Once は、一度だけ実行されることが保証される型
	//
	// 標準ライブラリの説明には以下の記載がある。
	//
	// > In the terminology of the Go memory model, the return from f “synchronizes before” the return from any call of once.Do(f).
	//
	// これは具体的には、fのリターンはonce.Do(f)のリターンよりも先に「同期」されるということを意味する。
	// これにより、f内での全てのメモリアクセスが、once.Do(f)呼び出し後のメモリアクセスよりも先に発生し、
	// それによってf内で行われた全ての操作がonce.Do(f)呼び出し後に見えることを保証する。
	//
	// 例えば、f内で変数を設定し、別のゴルーチンがonce.Do(f)を呼び出した後にその変数を読む場合、f内で設定された値が見えることが保証される。
	//

	type (
		_result struct {
			value string
		}
	)

	var (
		wg     sync.WaitGroup
		once   sync.Once
		result _result
		fn     = func() {
			output.Stderrl("[fn]", "called")
			result.value += "helloworld"
		}
	)

	wg.Add(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		i := i
		go func() {
			defer wg.Done()
			output.Stderrf("[goroutine]", ">>> %02d: call once.Do\n", i)
			once.Do(fn)
		}()
	}

	wg.Wait()
	output.Stdoutl("[result]", result.value)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: syncs_use_once

	   [Name] "syncs_use_once"
	   [goroutine]          >>> 15: call once.Do
	   [fn]                 called
	   [goroutine]          >>> 06: call once.Do
	   [goroutine]          >>> 07: call once.Do
	   [goroutine]          >>> 08: call once.Do
	   [goroutine]          >>> 05: call once.Do
	   [goroutine]          >>> 03: call once.Do
	   [goroutine]          >>> 09: call once.Do
	   [goroutine]          >>> 10: call once.Do
	   [goroutine]          >>> 11: call once.Do
	   [goroutine]          >>> 04: call once.Do
	   [goroutine]          >>> 12: call once.Do
	   [goroutine]          >>> 13: call once.Do
	   [goroutine]          >>> 14: call once.Do
	   [goroutine]          >>> 01: call once.Do
	   [goroutine]          >>> 02: call once.Do
	   [goroutine]          >>> 00: call once.Do
	   [v]                  helloworld


	   [Elapsed] 492.64µs
	*/

}
