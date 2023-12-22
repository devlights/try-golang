package goroutines

import (
	"context"
	"fmt"
	"time"

	"github.com/devlights/gomy/output"
)

// WithContextCancelCause は、Go 1.20 で新規追加された context.WithCancelCause のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/context@go1.20.1#WithCancelCause
func WithContextCancelCause() error {
	mainCtx, mainCxl := context.WithTimeout(context.Background(), 1*time.Second)
	defer mainCxl()

	withCancelCauseCtx := func(ctx context.Context) context.Context {
		//
		// Go 1.20 から WithCancelCause() が追加された
		// これにより、cancel 関数に対して任意のエラーを指定出来るようになっている
		//
		// cancel 時に指定したエラーは context.Cause(context.Context) で取得することに注意
		// 今までの context.Err() からは取得できない。
		//
		ctx, cxl := context.WithCancelCause(ctx)

		go func() {
			select {
			case <-time.After(500 * time.Millisecond):
				cxl(fmt.Errorf("my error"))
			case <-ctx.Done():
			}
		}()

		return ctx
	}(mainCtx)

	withCancelCtx := func(ctx context.Context) context.Context {
		ctx, cxl := context.WithCancel(ctx)

		go func() {
			select {
			case <-time.After(2 * time.Second):
				cxl()
			case <-ctx.Done():
			}
		}()

		return ctx
	}(withCancelCauseCtx)

	select {
	case <-withCancelCtx.Done():
	case <-withCancelCauseCtx.Done():
	case <-mainCtx.Done():
	}

	output.Stdoutl("[withCancelCtx.Err       ]", withCancelCtx.Err())
	output.Stdoutl("[withCancelCauseCtx.Err  ]", withCancelCauseCtx.Err())
	output.Stdoutl("[mainCtx.Err             ]", mainCtx.Err())

	// 当然であるが、下位のコンテキストで設定した任意のエラーは
	// 上位のコンテキストを指定しても取得できない。
	// 下位のコンテキストには伝播する。
	output.StdoutHr()
	output.Stdoutl("[withCancelCtx.Cause     ]", context.Cause(withCancelCtx))
	output.Stdoutl("[withCancelCauseCtx.Cause]", context.Cause(withCancelCauseCtx))
	output.Stdoutl("[mainCtx.Cause           ]", context.Cause(mainCtx))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: goroutines_with_context_cancelcause

	   [Name] "goroutines_with_context_cancelcause"
	   [withCancelCtx.Err       ] context canceled
	   [withCancelCauseCtx.Err  ] context canceled
	   [mainCtx.Err             ] <nil>
	   --------------------------------------------------
	   [withCancelCtx.Cause     ] my error
	   [withCancelCauseCtx.Cause] my error
	   [mainCtx.Cause           ] <nil>


	   [Elapsed] 500.744322ms
	*/

}
