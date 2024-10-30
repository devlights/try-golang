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
		latchCount = 3
	)
	var (
		ce = NewCountdownLatch(latchCount)
		wg sync.WaitGroup
	)

	for range 2 {
		ce.Reset(latchCount)

		for i := range 5 {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				log.Printf("[%2d] 待機開始", i)
				ce.Wait()
				log.Printf("[%2d] 待機解除", i)
			}(i)
		}

		for range 3 {
			<-time.After(time.Second)

			log.Printf("現在のカウント: %d\n", ce.CurrentCount())
			ce.Signal()
		}

		wg.Wait()
		log.Println("-------------------------------------")
	}

	return nil
}
