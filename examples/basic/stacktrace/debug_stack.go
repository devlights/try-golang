package stacktrace

import (
	"context"
	"runtime/debug"
	"time"

	"github.com/devlights/gomy/output"
)

// DebugStack -- debug.Stack() の サンプルです。
//
// REFERENCES
//   - https://pkg.go.dev/runtime/debug#Stack
//   - https://stackoverflow.com/questions/19094099/how-to-dump-goroutine-stacktraces
func DebugStack() error {
	// channels
	var (
		ch = make(chan []byte, 1)
	)
	defer close(ch)

	// funcs
	var (
		third = func(pCtx context.Context) context.Context {
			ctx, cxl := context.WithTimeout(pCtx, 100*time.Millisecond)
			defer cxl()

			// この処理が走っているgoroutineのスタックトレースを取得 (全部ではない)
			ch <- debug.Stack()

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

	// contexts
	var (
		ctx = first(context.Background())
	)

LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case buf := <-ch:
			output.Stdoutl("stack", string(buf))
		}
	}
	output.Stdoutl("done")

	return nil
}
