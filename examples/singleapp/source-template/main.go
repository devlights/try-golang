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
	TIMEOUT = time.Second
)

var (
	ErrTooSlow = errors.New("too slow")
)

func init() {
	log.SetFlags(log.Lmicroseconds)
}

func main() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithTimeoutCause(rootCtx, TIMEOUT, ErrTooSlow)
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

	return ctx
}

func proc(_ context.Context) error {
	/*
		ここに実装入れる
	*/
	return nil
}
