/*
	ちょっとしたサンプルを作成する際などに利用するソースのテンプレート（自分用）
*/

package main

import (
	"context"
	"errors"
	"log"
	"time"
)

const (
	MainTimeout = time.Second
	ProcTimeout = 200*time.Millisecond
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
	/*
		ここに実装入れる
	*/
	return nil
}
