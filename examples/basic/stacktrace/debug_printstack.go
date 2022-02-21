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
}
