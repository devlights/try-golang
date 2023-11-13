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
			done <- true
			output.Stderrf("[done]", "goroutine-%d\n", i)
		}()
	}

	for i := 0; i < numCpu; i++ {
		<-done
	}

	output.Stdoutl("[after]", v)

	return nil
}
