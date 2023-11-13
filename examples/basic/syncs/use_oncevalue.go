package syncs

import (
	"runtime"
	"sync"

	"github.com/devlights/gomy/output"
)

// UseOnceValue は、Go 1.21 で追加された sync.OnceValue() のサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/sync@go1.21.4#OnceValue
func UseOnceValue() error {
	//
	// Go 1.21 にて、sync.OnceValue() が追加された
	// 元々存在していた、sync.Once を使いやすくした感じ
	// 値を返却する関数を使える
	//

	var (
		v = 0
		f = sync.OnceValue(func() int {
			v++
			return v
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
			result := f()
			done <- true
			output.Stderrf("[done]", "result=%d\tgoroutine-%d\n", result, i)
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

		ENTER EXAMPLE NAME: syncs_use_oncevalue

		[Name] "syncs_use_oncevalue"
		[before]             0
		[done]               result=1   goroutine-5
		[done]               result=1   goroutine-11
		[done]               result=1   goroutine-4
		[done]               result=1   goroutine-3
		[done]               result=1   goroutine-2
		[done]               result=1   goroutine-12
		[done]               result=1   goroutine-14
		[done]               result=1   goroutine-7
		[done]               result=1   goroutine-6
		[done]               result=1   goroutine-8
		[done]               result=1   goroutine-0
		[done]               result=1   goroutine-1
		[done]               result=1   goroutine-9
		[done]               result=1   goroutine-13
		[done]               result=1   goroutine-15
		[done]               result=1   goroutine-10
		[after]              1


		[Elapsed] 640.98µs
	*/
}
