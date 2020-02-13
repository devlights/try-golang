package main

import (
	"fmt"
	"github.com/devlights/try-golang/lib"
)

type (
	ExecError struct {
		Name lib.SampleKey
		Err  error
	}
)

func (e *ExecError) Error() string {
	return fmt.Sprintf("[Error] %s (%s)", e.Err.Error(), e.Name)
}

func (e *ExecError) Unwrap() error {
	return e.Err
}