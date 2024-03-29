package helloworld

import (
	"context"
	"time"

	"github.com/devlights/gomy/ctxs"
	"github.com/devlights/gomy/output"
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

	// start tasks
	var (
		syncCtx  = syncOp(procCtx)
		asyncCtx = asyncOp(procCtx)
	)

	// wait until all tasks are completed
	<-ctxs.WhenAll(procCtx, syncCtx, asyncCtx).Done()

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: helloworld_mixed

	   [Name] "helloworld_mixed"
	   async                [02] helloworld
	   sync                 [01] helloworld
	   sync                 [02] helloworld
	   sync                 [03] helloworld
	   sync                 [04] helloworld
	   sync                 [05] helloworld
	   sync                 [06] helloworld
	   async                [01] helloworld
	   sync                 [07] helloworld
	   sync                 [08] helloworld
	   sync                 [09] helloworld
	   sync                 [10] helloworld
	   async                [04] helloworld
	   async                [06] helloworld
	   async                [08] helloworld
	   async                [03] helloworld
	   async                [07] helloworld
	   async                [05] helloworld
	   async                [09] helloworld
	   async                [10] helloworld


	   [Elapsed] 988.12µs
	*/

}

func syncOp(pCtx context.Context) context.Context {
	var (
		ctx, cxl = context.WithCancel(pCtx)
	)

	go func() {
		defer cxl()

		for v := range items() {
			v := v
			<-exec(ctx, v+1, "sync ").Done()
		}
	}()

	return ctx
}

func asyncOp(pCtx context.Context) context.Context {
	var (
		ctx, cxl = context.WithCancel(pCtx)
		tasks    = make([]context.Context, 0)
	)

	for v := range items() {
		v := v
		tasks = append(tasks, exec(ctx, v+1, "async"))
	}

	go func() {
		defer cxl()
		<-ctxs.WhenAll(ctx, tasks...).Done()
	}()

	return ctx
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

func exec(pCtx context.Context, v int, prefix string) context.Context {
	var (
		ctx, cxl = context.WithCancel(pCtx)
	)

	go func() {
		defer cxl()

		select {
		case <-ctx.Done():
			return
		default:
			output.Stderrf(prefix, "[%02d] helloworld\n", v)
		}
	}()

	return ctx
}
