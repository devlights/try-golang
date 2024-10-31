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
	var (
		gate = NewGate()
		wg   sync.WaitGroup
	)

	// 10個のゴルーチンがゲート前に待機する
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			log.Printf("[%2d] 待機開始", i)
			gate.Await()
			log.Printf("[%2d] 待機解除", i)
		}(i)
	}

	// 何か準備処理などを行っているとする
	<-time.After(time.Second)
	log.Println("-------------------------------------")

	// ゲートを開き、待機解除したゴルーチン達が全完了するのを待つ
	gate.Open()
	wg.Wait()

	// 一度開いたゲートは開きっぱなしになるため、開いた後のAwait呼び出しは即座に返る.
	gate.Await()
	gate.Await()

	return nil
}
