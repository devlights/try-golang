package errs

import (
	"errors"
	"fmt"
	"time"

	"github.com/devlights/gomy/output"
)

// ErrorJoinFmtErrorf は、Go1.20から追加された複数エラーを纏める機能が追加されたfmt.Errorfのサンプルです.
func ErrorJoinFmtErrorf() error {
	//
	// Go1.20 より、fmt.Errorf でも複数のエラーを纏めることが出来るようになった
	//
	var (
		err1  = errors.New("this is err1 message")
		err2  = errors.New("this is err2 message")
		err3  = errors.New("this is err3 message")
		errCh = make(chan error, 1)
	)

	var (
		fn1 = func() error {
			time.Sleep(100 * time.Millisecond)
			return err1
		}
		fn2 = func() error {
			time.Sleep(150 * time.Millisecond)
			return err2
		}
		fn3 = func(errCh chan<- error) {
			errors := make([]any, 0)

			if err := fn1(); err != nil {
				errors = append(errors, err)
			}

			if err := fn2(); err != nil {
				errors = append(errors, err)
			}

			// エラーを纏める
			errCh <- fmt.Errorf("errors: (%w) (%w)", errors...)
		}
	)
	defer close(errCh)

	go fn3(errCh)

	select {
	case err := <-errCh:
		output.Stdoutl("[err  ]", err)
		output.Stdoutl("[err1?]", errors.Is(err, err1))
		output.Stdoutl("[err2?]", errors.Is(err, err2))
		output.Stdoutl("[err3?]", errors.Is(err, err3))
	case <-time.After(1 * time.Second):
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: error_join_fmt_errorf

	   [Name] "error_join_fmt_errorf"
	   [err  ]              errors: (this is err1 message) (this is err2 message)
	   [err1?]              true
	   [err2?]              true
	   [err3?]              false


	   [Elapsed] 250.685168ms
	*/

}
