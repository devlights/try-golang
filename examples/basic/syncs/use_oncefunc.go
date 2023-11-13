package syncs

import (
	"runtime"
	"sync"

	"github.com/devlights/gomy/output"
)

// UseOnceFunc は、Go 1.21 で追加された sync.OnceFunc() のサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/sync@go1.21.4#OnceFunc
func UseOnceFunc() error {
	//
	// Go 1.21 にて、sync.OnceFunc() が追加された
	// 元々存在していた、sync.Once を使いやすくした感じ
	//

	var (
		v = 0
		f = sync.OnceFunc(func() {
			v++
		})
	)
	output.Stdoutl("[before]", v)

	var (
		numCpu = runtime.NumCPU()
		done   = make(chan bool, numCpu)
	)
	for i := 0; i < numCpu; i++ {
		i := i
		go func() {
			f()
			output.Stderrf("[done]", "goroutine-%d\n", i)
			done <- true
		}()
	}

	for i := 0; i < numCpu; i++ {
		<-done
	}

	output.Stdoutl("[after]", v)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: syncs_use_oncefunc

	   [Name] "syncs_use_oncefunc"
	   [before]             0
	   [done]               goroutine-15
	   [done]               goroutine-10
	   [done]               goroutine-3
	   [done]               goroutine-0
	   [done]               goroutine-11
	   [done]               goroutine-14
	   [done]               goroutine-6
	   [done]               goroutine-1
	   [done]               goroutine-2
	   [done]               goroutine-5
	   [done]               goroutine-12
	   [done]               goroutine-4
	   [done]               goroutine-7
	   [done]               goroutine-8
	   [done]               goroutine-9
	   [done]               goroutine-13
	   [after]              1


	   [Elapsed] 563.92µs
	*/
}
