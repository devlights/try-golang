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
	log.SetFlags(log.Ltime)
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
		WORKER_COUNT = 3
	)
	var (
		barrier = NewCyclicBarrier(WORKER_COUNT)
		wg      sync.WaitGroup
	)

	// 3つのワーカーを起動し、全員揃ったら先に進むを繰り返す
	for i := 0; i < WORKER_COUNT; i++ {
		wg.Add(1)
		go worker(i+1, &wg, barrier)
	}

	wg.Wait()

	return nil
}

func worker(id int, wg *sync.WaitGroup, barrier *CyclicBarrier) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		log.Printf("Worker-[%2d] 準備作業 %2d週目", id, i+1)
		time.Sleep(time.Duration(id) * time.Second)

		log.Printf("Worker-[%2d] 待機開始", id)
		{
			if err := barrier.Await(); err != nil {
				log.Printf("Worker-[%2d] エラー: %v", id, err)
				return
			}
		}
		log.Printf("Worker-[%2d] 待機解除", id)
	}
}
