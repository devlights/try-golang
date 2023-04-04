package errs

import (
	"errors"
	"fmt"
	"time"

	"github.com/devlights/gomy/output"
)

type MyError1 struct {
	Message string
	Code    int
}

type MyError2 struct {
	Message string
	Code    int
}

type MyError3 struct {
	Message string
	Code    int
}

func (me MyError1) Error() string {
	return fmt.Sprintf("%d:%s", me.Code, me.Message)
}

func (me MyError2) Error() string {
	return fmt.Sprintf("%d:%s", me.Code, me.Message)
}

func (me MyError3) Error() string {
	return fmt.Sprintf("%d:%s", me.Code, me.Message)
}

// ErrorJoinErrorsAs は、Go1.20から追加された errors.Join() と errors.As() の関係性についてのサンプルです.
func ErrorJoinErrorsAs() error {
	var (
		errCh = make(chan error, 1)
	)
	defer close(errCh)

	go func(errCh chan<- error) {
		time.Sleep(100 * time.Millisecond)

		e1 := &MyError1{
			Message: "MyError1",
			Code:    1,
		}
		e2 := &MyError2{
			Message: "MyError2",
			Code:    2,
		}

		errCh <- errors.Join(e1, e2)
	}(errCh)

	select {
	case <-time.After(1 * time.Second):
	case err, ok := <-errCh:
		if ok {
			var e1 *MyError1
			var e2 *MyError2
			var e3 *MyError3

			output.Stdoutl("[err  ]", err)
			output.Stdoutl("[MyError1?]", errors.As(err, &e1))
			output.Stdoutl("[MyError2?]", errors.As(err, &e2))
			output.Stdoutl("[MyError3?]", errors.As(err, &e3))

			output.Stdoutl("[e1]", *e1)
			output.Stdoutl("[e2]", *e2)
			output.Stdoutl("[e3]", e3) // e3はnilポインタとなっているのでデリファレンスするとpanicする
		}
	}

	return nil
}
