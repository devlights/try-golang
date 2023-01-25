package goroutines

import (
	"context"
	"time"

	"github.com/devlights/gomy/output"
)

// ContextAndTimeAfterFunc は、Context と time.AfterFunc でキャンセルするサンプルです.
func ContextAndTimeAfterFunc() error {
	//
	// 下の２つは同じ結果になる
	//
	useTimeAfterFunc()
	useWithTimeout()

	return nil
}

// useTimeAfterFunc は、context.WithCancel + time.AfterFunc でタイムアウトさせます
func useTimeAfterFunc() {
	var (
		ctx, cxl = context.WithCancel(context.Background())
	)
	time.AfterFunc(2*time.Second, func() { cxl() })

	output.Stdoutl("[useTimeAfterFunc]", time.Now().Format("15:04:05"))
	{
		<-ctx.Done()
	}
	output.Stdoutl("[useTimeAfterFunc]", time.Now().Format("15:04:05"))
}

// useWithTimeout は、context.WithTimeout でタイムアウトさせます
func useWithTimeout() {
	var (
		ctx, cxl = context.WithTimeout(context.Background(), 2*time.Second)
	)
	defer cxl()

	output.Stdoutl("[useWithTimeout]", time.Now().Format("15:04:05"))
	{
		<-ctx.Done()
	}
	output.Stdoutl("[useWithTimeout]", time.Now().Format("15:04:05"))
}
