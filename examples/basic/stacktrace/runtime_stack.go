package stacktrace

import (
	"runtime"

	"github.com/devlights/gomy/output"
)

// RuntimeStack -- runtime.Stack() についてのサンプルです。
//
// REFERENCES
//   - https://pkg.go.dev/runtime#Stack
//   - https://stackoverflow.com/questions/19094099/how-to-dump-goroutine-stacktraces
func RuntimeStack() error {
	// channels
	var (
		chSingle = make(chan []byte, 1)
		chAll    = make(chan []byte, 1)
	)

	// funcs
	var (
		getStack = func(all bool) []byte {
			// From src/runtime/debug/stack.go
			var (
				buf = make([]byte, 1024)
			)

			for {
				n := runtime.Stack(buf, all)
				if n < len(buf) {
					return buf[:n]
				}
				buf = make([]byte, 2*len(buf))
			}
		}
	)

	// runtime.Stack() の 第２引数 に false を渡すと、現在のgoroutineのみが対象
	go func(ch chan<- []byte) {
		defer close(ch)
		ch <- getStack(false)
	}(chSingle)

	// runtime.Stack() の 第２引数 に true  を渡すと、全てのgoroutineが対象
	go func(ch chan<- []byte) {
		defer close(ch)
		ch <- getStack(true)
	}(chAll)

	// runtime.Stack(false) の 結果
	for v := range chSingle {
		output.Stdoutl("=== stack-single", string(v))
	}

	// runtime.Stack(true) の 結果
	for v := range chAll {
		output.Stdoutl("=== stack-all   ", string(v))
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: stacktrace_runtime_stack

	   [Name] "stacktrace_runtime_stack"
	   === stack-single     goroutine 6 [running]:
	   github.com/devlights/try-golang/examples/basic/stacktrace.RuntimeStack.func1(0x0)
	           /workspace/try-golang/examples/basic/stacktrace/runtime_stack.go:30 +0x65
	   github.com/devlights/try-golang/examples/basic/stacktrace.RuntimeStack.func2(0xc0000a4300)
	           /workspace/try-golang/examples/basic/stacktrace/runtime_stack.go:42 +0x54
	   created by github.com/devlights/try-golang/examples/basic/stacktrace.RuntimeStack in goroutine 1
	           /workspace/try-golang/examples/basic/stacktrace/runtime_stack.go:40 +0xc5

	   === stack-all        goroutine 7 [running]:
	   github.com/devlights/try-golang/examples/basic/stacktrace.RuntimeStack.func1(0x1)
	           /workspace/try-golang/examples/basic/stacktrace/runtime_stack.go:30 +0x65
	   github.com/devlights/try-golang/examples/basic/stacktrace.RuntimeStack.func3(0xc0000a4360)
	           /workspace/try-golang/examples/basic/stacktrace/runtime_stack.go:48 +0x57
	   created by github.com/devlights/try-golang/examples/basic/stacktrace.RuntimeStack in goroutine 1
	           /workspace/try-golang/examples/basic/stacktrace/runtime_stack.go:46 +0x139

	   goroutine 1 [chan receive]:
	   github.com/devlights/try-golang/examples/basic/stacktrace.RuntimeStack()
	           /workspace/try-golang/examples/basic/stacktrace/runtime_stack.go:52 +0x1e6
	   github.com/devlights/try-golang/runner.(*Exec).Run(0x8070102?)
	           /workspace/try-golang/runner/exec.go:52 +0x131
	   github.com/devlights/try-golang/runner.(*Loop).exec(0xc00002e2d0?, {0x9c2a31, 0x18}, 0xc0000bda40)
	           /workspace/try-golang/runner/loop.go:126 +0x85
	   github.com/devlights/try-golang/runner.(*Loop).Run(0xc00007c410)
	           /workspace/try-golang/runner/loop.go:79 +0x23e
	   github.com/devlights/try-golang/cmd.Execute()
	           /workspace/try-golang/cmd/root.go:66 +0x612
	   main.main()
	           /workspace/try-golang/main.go:6 +0xf

	   goroutine 6 [runnable]:
	   github.com/devlights/try-golang/examples/basic/stacktrace.RuntimeStack.gowrap1()
	           /workspace/try-golang/examples/basic/stacktrace/runtime_stack.go:40
	   runtime.goexit({})
	           /home/gitpod/go/src/runtime/asm_amd64.s:1695 +0x1
	   created by github.com/devlights/try-golang/examples/basic/stacktrace.RuntimeStack in goroutine 1
	           /workspace/try-golang/examples/basic/stacktrace/runtime_stack.go:40 +0xc5



	   [Elapsed] 334.57µs
	*/

}
