package helloworld

import (
	"context"
	"fmt"
	"time"

	"github.com/devlights/gomy/ctxs"
)

// Mixed -- 同期と非同期の両方で同じことをするサンプル
func Mixed() error {
	// main contexts
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
	)
	defer mainCxl()

	// proc context
	var (
		procCtx, procCxl = context.WithTimeout(mainCtx, 1*time.Second)
	)
	defer procCxl()

	// synchronous
	<-sync(procCtx).Done()

	fmt.Println("--------------------------------")

	// asynchronous
	<-async(procCtx).Done()

	return nil
}

func sync(pCtx context.Context) context.Context {
	var (
		ctx, cxl = context.WithCancel(pCtx)
	)
	defer cxl()

	for v := range items() {
		v := v
		<-exec(ctx, v+1).Done()
	}

	return ctx
}

func async(pCtx context.Context) context.Context {
	var (
		ctx, cxl = context.WithCancel(pCtx)
		tasks    = make([]context.Context, 0)
	)
	defer cxl()

	for v := range items() {
		v := v
		tasks = append(tasks, exec(ctx, v+1))
	}

	return ctxs.WhenAll(ctx, tasks...)
}

func items() <-chan int {
	var (
		ch = make(chan int)
	)

	go func() {
		defer close(ch)

		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	return ch
}

func exec(pCtx context.Context, v int) context.Context {
	var (
		ctx, cxl = context.WithCancel(pCtx)
	)

	go func() {
		defer cxl()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Printf("[%02d] helloworld\n", v)
		}
	}()

	return ctx
}
