package helloworld

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/devlights/gomy/chans"
	"github.com/devlights/gomy/ctxs"
)

// Async -- HelloWorld 非同期版
func Async() error {
	var (
		rootCtx             = context.Background()
		mainCtx, mainCancel = context.WithCancel(rootCtx)
		procCtx, procCancel = context.WithTimeout(mainCtx, 100*time.Millisecond)
	)

	defer mainCancel()
	defer procCancel()

	var tasks []context.Context
	for i := 0; i < 10; i++ {
		tasks = append(tasks, func(pCtx context.Context, no, delay int) context.Context {
			ctx, cancel := context.WithCancel(pCtx)

			go func() {
				defer cancel()

				select {
				case <-ctx.Done():
					fmt.Printf("[%d] Time out\t(%02d msec delay)\n", no, delay)
				case <-time.After(time.Duration(delay) * time.Millisecond):
					fmt.Printf("[%d] Hello World\t(%02d msec delay)\n", no, delay)
				}
			}()

			return ctx
		}(procCtx, i, rand.Intn(100)))
	}

	<-chans.WhenAll(ctxs.ToDoneCh(tasks...)...)

	return nil
}
