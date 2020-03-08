package main

import (
	"fmt"

	"github.com/devlights/try-golang/interfaces"
)

type (
	ExecError struct {
		Name interfaces.ExampleKey
		Err  error
	}
)

func (e *ExecError) Error() string {
	return fmt.Sprintf("[Error] %s (%s)", e.Err.Error(), e.Name)
}

func (e *ExecError) Unwrap() error {
	return e.Err
}
