package stacktrace

import (
	"context"
	"runtime/debug"
	"time"

	"github.com/devlights/gomy/output"
)

// DebugPrintStack -- runtime/debug.PrintStack のサンプルです.
//
// REFERENCES
//   - https://pkg.go.dev/runtime/debug#PrintStack
//   - https://stackoverflow.com/questions/19094099/how-to-dump-goroutine-stacktraces
func DebugPrintStack() error {
	var (
		third = func(pCtx context.Context) context.Context {
			ctx, cxl := context.WithTimeout(pCtx, 100*time.Millisecond)
			defer cxl()

			// この処理が走っているgoroutineのスタックトレースを出力 (全部ではない)
			output.StderrHr()
			debug.PrintStack()

			<-ctx.Done()

			return ctx
		}
		second = func(pCtx context.Context) context.Context {
			ctx, cxl := context.WithCancel(pCtx)
			go func() {
				defer cxl()
				<-third(ctx).Done()
			}()

			return ctx
		}
		first = func(pCtx context.Context) context.Context {
			ctx, cxl := context.WithCancel(pCtx)
			go func() {
				defer cxl()
				<-second(ctx).Done()
			}()

			return ctx
		}
	)

	<-first(context.Background()).Done()

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: stacktrace_debug_printstack

	   [Name] "stacktrace_debug_printstack"
	   --------------------------------------------------
	   goroutine 19 [running]:
	   runtime/debug.Stack()
	           /home/gitpod/go/src/runtime/debug/stack.go:24 +0x5e
	   runtime/debug.PrintStack()
	           /home/gitpod/go/src/runtime/debug/stack.go:16 +0x13
	   github.com/devlights/try-golang/examples/basic/stacktrace.DebugPrintStack.func1({0xa984f0?, 0xc000118050?})
	           /workspace/try-golang/examples/basic/stacktrace/debug_printstack.go:24 +0x85
	   github.com/devlights/try-golang/examples/basic/stacktrace.DebugPrintStack.func2.1()
	           /workspace/try-golang/examples/basic/stacktrace/debug_printstack.go:34 +0x3c
	   created by github.com/devlights/try-golang/examples/basic/stacktrace.DebugPrintStack.func2 in goroutine 18
	           /workspace/try-golang/examples/basic/stacktrace/debug_printstack.go:32 +0xa5


	   [Elapsed] 100.413649ms
	*/

}
