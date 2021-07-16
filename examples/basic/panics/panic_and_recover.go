package panics

import (
	"errors"

	"github.com/devlights/gomy/output"
)

// PanicAndRecover -- panicとrecoverのサンプルです.
func PanicAndRecover() error {
	if err := catch(); err != nil {
		output.Stdoutl("[catch]", err)
	}

	return nil
}

func catch() (err error) {
	defer func() {
		if obj := recover(); obj != nil {
			if _, ok := obj.(error); ok {
				err = obj.(error)
			}
		} else {
			err = errors.New("何か発生した")
		}
	}()

	raise()
	return nil
}

func raise() {
	panic(errors.New("error occurred in raise()"))
}
