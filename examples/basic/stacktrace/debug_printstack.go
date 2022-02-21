package stacktrace

import (
	"context"
	"runtime/debug"
	"time"

	"github.com/devlights/gomy/output"
)

// DebugPrintStack -- runtime/debug.PrintStack のサンプルです.
func DebugPrintStack() error {
	<-first(context.Background()).Done()
	return nil
}

func first(pCtx context.Context) context.Context {
	ctx, cxl := context.WithCancel(pCtx)
	go func() {
		defer cxl()
		<-second(ctx).Done()
	}()

	return ctx
}

func second(pCtx context.Context) context.Context {
	ctx, cxl := context.WithCancel(pCtx)
	go func() {
		defer cxl()
		<-third(ctx).Done()
	}()

	return ctx
}

func third(pCtx context.Context) context.Context {
	ctx, cxl := context.WithTimeout(pCtx, 100*time.Millisecond)
	defer cxl()

	// この処理が走っているgoroutineのスタックトレースを出力 (全部ではない)
	output.StderrHr()
	debug.PrintStack()

	<-ctx.Done()

	return ctx
}
