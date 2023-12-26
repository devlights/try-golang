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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: helloworld_async

	   [Name] "helloworld_async"
	   [06]    Hello World     (01 msec delay)
	   [10]    Hello World     (03 msec delay)
	   [01]    Hello World     (28 msec delay)
	   [08]    Hello World     (49 msec delay)
	   [09]    Hello World     (64 msec delay)
	   [07]    Hello World     (65 msec delay)
	   [04]    Hello World     (66 msec delay)
	   [05]    Hello World     (69 msec delay)
	   [02]    Hello World     (75 msec delay)
	   [03]    Hello World     (91 msec delay)


	   [Elapsed] 91.850415ms
	*/

}
