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

	/*
	    $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: goroutines_with_context_cancel

	   [Name] "goroutines_with_context_cancel"
	   [goroutine]          This line is printed


	   [Elapsed] 50.7µs
	*/

}
