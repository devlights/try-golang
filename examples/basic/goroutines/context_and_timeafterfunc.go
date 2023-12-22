package goroutines

import (
	"context"
	"time"

	"github.com/devlights/gomy/output"
)

// ContextAndTimeAfterFunc は、Context と time.AfterFunc でキャンセルするサンプルです.
func ContextAndTimeAfterFunc() error {
	//
	// 下の２つは同じ結果になるが、context.Err() の内容は異なる
	//
	useTimeAfterFunc()
	useWithTimeout()

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: goroutines_context_and_timeafterfunc

	   [Name] "goroutines_context_and_timeafterfunc"
	   [useTimeAfterFunc]   04:36:13
	   [useTimeAfterFunc]   context canceled
	   [useTimeAfterFunc]   04:36:15
	   [useWithTimeout]     04:36:15
	   [useWithTimeout]     context deadline exceeded
	   [useWithTimeout]     04:36:17


	   [Elapsed] 4.000603645s
	*/

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
		output.Stdoutl("[useTimeAfterFunc]", ctx.Err())
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
		output.Stdoutl("[useWithTimeout]", ctx.Err())
	}
	output.Stdoutl("[useWithTimeout]", time.Now().Format("15:04:05"))
}
