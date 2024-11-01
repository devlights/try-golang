package main

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

const (
	MainTimeout = 20 * time.Second
	ProcTimeout = 10 * time.Second
)

var (
	ErrMainTooSlow = errors.New("(MAIN) TOO SLOW")
	ErrProcTooSlow = errors.New("(PROC) TOO SLOW")
)

func init() {
	log.SetFlags(0)
}

func main() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithTimeoutCause(rootCtx, MainTimeout, ErrMainTooSlow)
		procCtx          = run(mainCtx)
		err              error
	)
	defer mainCxl()

	select {
	case <-mainCtx.Done():
		err = context.Cause(mainCtx)
	case <-procCtx.Done():
		if err = context.Cause(procCtx); errors.Is(err, context.Canceled) {
			err = nil
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}

func run(pCtx context.Context) context.Context {
	var (
		ctx, cxl = context.WithCancelCause(pCtx)
	)

	go func() {
		cxl(proc(ctx))
	}()
	go func() {
		<-time.After(ProcTimeout)
		cxl(ErrProcTooSlow)
	}()

	return ctx
}

func proc(_ context.Context) error {
	const (
		numLatchs     = 3
		numGoroutines = 5
	)
	var (
		latch = NewCountdownLatch(numLatchs)
	)
	for range 2 {
		var (
			wg sync.WaitGroup
		)

		latch.Reset(numLatchs)

		for i := range numGoroutines {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				log.Printf("[%2d] 待機開始", i)
				latch.Wait()
				log.Printf("[%2d] 待機解除", i)
			}(i)
		}

		for range numLatchs {
			<-time.After(time.Second)

			log.Printf("現在のカウント: %d\n", latch.CurrentCount())
			latch.Signal()
		}

		wg.Wait()
		log.Println("----------------")
	}

	return nil
}
