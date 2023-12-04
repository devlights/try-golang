package errs

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/devlights/gomy/ctxs"
	"github.com/devlights/gomy/output"
)

// ErrorJoin は、Go 1.20 で追加された errors.Join のサンプルです.
//
// Go 1.20 で、errors.Join が追加され、複数のエラーを纏めて一つに
// することが標準ライブラリで出来るようになった.
//
// # REFERENCES
//   - https://pkg.go.dev/errors@go1.20.1#Join
func ErrorJoin() error {
	var (
		rootCtx            = context.Background()
		mainCtx, mainCxl   = context.WithCancel(rootCtx)
		proc1Ctx, proc1Cxl = context.WithCancel(mainCtx)
		proc2Ctx, proc2Cxl = context.WithCancel(mainCtx)
		errCh              = make(chan error)
		rnd                = rand.New(rand.NewSource(time.Now().UnixNano()))
		sleep              = func() {
			duration, _ := time.ParseDuration(fmt.Sprintf("%dms", rnd.Intn(500)))
			time.Sleep(duration)
		}
	)
	defer mainCxl()

	// 単一のerror
	go func() {
		defer proc1Cxl()

		output.Stderrl("[proc1]", "start")
		sleep()

		errCh <- errors.New("proc1 error")
	}()

	// 複数のerror
	go func() {
		defer proc2Cxl()

		output.Stderrl("[proc2]", "start")
		sleep()

		err1 := errors.New("proc2 error1")
		errCh <- err1
		sleep()

		err2 := errors.New("proc2 error2")
		errCh <- err2
		sleep()

		err3 := errors.New("proc2 error3")
		errCh <- err3
		sleep()

		errCh <- errors.Join(errors.New("proc2 all errors"), err1, err2, err3)
	}()

	go func(ctx context.Context) {
		defer close(errCh)
		<-ctx.Done()
	}(ctxs.WhenAll(mainCtx, proc1Ctx, proc2Ctx))

	for e := range errCh {
		output.Stdoutl("[Error]", e)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: error_join

	   [Name] "error_join"
	   [proc1]              start
	   [proc2]              start
	   [Error]              proc2 error1
	   [Error]              proc1 error
	   [Error]              proc2 error2
	   [Error]              proc2 error3
	   [Error]              proc2 all errors
	   proc2 error1
	   proc2 error2
	   proc2 error3


	   [Elapsed] 951.91721ms
	*/

}
