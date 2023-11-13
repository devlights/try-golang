package syncs

import (
	"runtime"
	"sync"

	"github.com/devlights/gomy/output"
)

// UseOnceValues は、Go 1.21 で追加された sync.OnceValues() のサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/sync@go1.21.4#OnceValues
func UseOnceValues() error {
	var (
		v = 0
		f = sync.OnceValues(func() (int, error) {
			v++
			return v, nil
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
			result, err := f()
			output.Stderrf("[done]", "result=%d\terr=%v\tgoroutine-%d\n", result, err, i)
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
		task: Task "build" is up to date
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: syncs_use_oncevalues

		[Name] "syncs_use_oncevalues"
		[before]             0
		[done]               result=1   err=<nil>       goroutine-15
		[done]               result=1   err=<nil>       goroutine-13
		[done]               result=1   err=<nil>       goroutine-14
		[done]               result=1   err=<nil>       goroutine-10
		[done]               result=1   err=<nil>       goroutine-1
		[done]               result=1   err=<nil>       goroutine-0
		[done]               result=1   err=<nil>       goroutine-2
		[done]               result=1   err=<nil>       goroutine-6
		[done]               result=1   err=<nil>       goroutine-5
		[done]               result=1   err=<nil>       goroutine-11
		[done]               result=1   err=<nil>       goroutine-3
		[done]               result=1   err=<nil>       goroutine-8
		[done]               result=1   err=<nil>       goroutine-9
		[done]               result=1   err=<nil>       goroutine-7
		[done]               result=1   err=<nil>       goroutine-4
		[done]               result=1   err=<nil>       goroutine-12
		[after]              1


		[Elapsed] 460.969µs
	*/
}
