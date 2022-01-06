package helloworld

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/devlights/gomy/ctxs"
)

// Async -- HelloWorld 非同期版
func Async() error {
	// main contexts
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
	)
	defer mainCxl()

	// proc context
	var (
		timeLimit        = 100 * time.Millisecond
		procCtx, procCxl = context.WithTimeout(mainCtx, timeLimit)
	)
	defer procCxl()

	// start tasks
	var tasks []context.Context
	for i := 0; i < 10; i++ {
		tasks = append(tasks, func(pCtx context.Context, no, delay int) context.Context {
			var (
				ctx, cxl = context.WithCancel(pCtx)
			)

			go func() {
				defer cxl()

				select {
				case <-ctx.Done():
					fmt.Printf("[%02d]\tTime out\t(%02d msec delay)\n", no, delay)
				case <-time.After(time.Duration(delay) * time.Millisecond):
					fmt.Printf("[%02d]\tHello World\t(%02d msec delay)\n", no, delay)
				}
			}()

			return ctx
		}(procCtx, i+1, rand.Intn(100)))
	}

	// wait until all tasks are completed
	<-ctxs.WhenAll(procCtx, tasks...).Done()

	return nil
}
