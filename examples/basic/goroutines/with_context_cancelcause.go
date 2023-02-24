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
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithTimeout(rootCtx, 2*time.Second)
		procCtx          context.Context
	)
	defer mainCxl()

	procCtx = func(ctx context.Context) context.Context {
		//
		// Go 1.20 から WithCancelCause() が追加された
		// これにより、cancel 関数に対して任意のエラーを指定出来るようになっている
		//
		// cancel 時に指定したエラーは context.Cause(context.Context) で取得することに注意
		// 今までの context.Err() からは取得できない。
		//
		ctx, cxl := context.WithCancelCause(ctx)

		go func() {
			<-time.After(1 * time.Second)
			cxl(fmt.Errorf("my error"))
		}()

		return ctx
	}(mainCtx)

	select {
	case <-procCtx.Done():
	case <-mainCtx.Done():
	}

	output.Stdoutl("[procctx.Err]", procCtx.Err())
	output.Stdoutl("[mainctx.Err]", mainCtx.Err())

	// 当然であるが、下位のコンテキストで設定した任意のエラーは
	// 上位のコンテキストを指定しても取得できない。
	output.Stdoutl("[procctx.Cause]", context.Cause(procCtx))
	output.Stdoutl("[mainctx.Cause]", context.Cause(mainCtx))
	output.Stdoutl("[rootctx.Cause]", context.Cause(rootCtx))

	return nil
}
