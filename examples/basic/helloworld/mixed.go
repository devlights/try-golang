package helloworld

import (
	"context"
	"fmt"
	"time"

	"github.com/devlights/gomy/ctxs"
)

// Mixed -- 同期と非同期の両方で同じことをするサンプル
func Mixed() error {
	var (
		mainCtx          = context.Background()
		procCtx, procCxl = context.WithTimeout(mainCtx, 1*time.Second)
	)
	defer procCxl()

	<-sync(procCtx).Done()
	fmt.Println("--------------------------------")
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
			fmt.Printf("[%d] helloworld\n", v)
		}
	}()

	return ctx
}
