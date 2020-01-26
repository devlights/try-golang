package main

import "fmt"

type (
	ExecError struct {
		Name string
		Err  error
	}
)

func (e *ExecError) Error() string {
	return fmt.Sprintf("[Error] %s (%s)", e.Err.Error(), e.Name)
}

func (e *ExecError) Unwrap() error {
	return e.Err
}
