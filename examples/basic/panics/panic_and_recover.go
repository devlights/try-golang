package panics

import (
	"errors"
	"fmt"

	"github.com/devlights/gomy/output"
)

// PanicAndRecover -- panicとrecoverのサンプルです.
func PanicAndRecover() error {
	if err := catch(); err != nil {
		output.Stdoutl("[catch]", err)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: panics_panic_and_recover

	   [Name] "panics_panic_and_recover"
	   [panic]              panic時deferは呼ばれる
	   [catch]              error occurred in raise()


	   [Elapsed] 51.68µs
	*/

}

func catch() (err error) {
	defer output.Stdoutl("[panic]", "panic時deferは呼ばれる")
	defer func() {
		// panicはrecoverで補足できるが、取得できる値は interface{} となる
		if obj := recover(); obj != nil {
			// error かどうか判別
			if e, ok := obj.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", obj)
			}
		} else {
			// panic(nil) としていると、ここに入る
			err = errors.New("何か発生した")
		}
	}()

	raise()
	return nil
}

func raise() {
	panic(errors.New("error occurred in raise()"))
	//panic("error じゃなく 文字列 を指定")
	//panic(nil)
}
