package goroutines

import (
	"context"

	"github.com/devlights/gomy/output"
)

// WithContextCancel -- context.Contextを用いて待ち合わせを行うサンプルです.
func WithContextCancel() error {
	var (
		rootCtx             = context.Background()
		mainCtx, mainCancel = context.WithCancel(rootCtx)
	)

	defer mainCancel()

	ctx := func(pCtx context.Context) context.Context {
		ctx, cancel := context.WithCancel(pCtx)

		go func() {
			defer cancel()
			output.Stdoutl("[goroutine]", "This line is printed")
		}()

		return ctx
	}(mainCtx)

	<-ctx.Done()

	return nil
}
