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
	)

	defer mainCancel()

	var tasks []context.Context
	for i := 0; i < 10; i++ {
		tasks = append(tasks, func(pCtx context.Context, no, delay int) context.Context {
			ctx, cancel := context.WithTimeout(pCtx, time.Duration(delay)*time.Millisecond)

			go func() {
				defer cancel()

				select {
				case <-ctx.Done():
				}

				fmt.Printf("[%d] Hello World\t(%02d msec delay)\n", no, delay)
			}()

			return ctx
		}(mainCtx, i, rand.Intn(100)))
	}

	<-chans.WhenAll(ctxs.ToDoneCh(tasks...)...)

	return nil
}
