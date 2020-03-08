package main

import (
	"fmt"

	"github.com/devlights/try-golang/interfaces"
)

type (
	ExecCommand struct {
		Args *ExecArgs
	}

	ExecArgs struct {
		Target  string
		Mapping interfaces.ExampleMapping
	}
)

func NewExecArgs(target string, mapping interfaces.ExampleMapping) *ExecArgs {
	a := new(ExecArgs)
	a.Target = target
	a.Mapping = mapping
	return a
}

func NewExecCommand(args *ExecArgs) *ExecCommand {
	c := new(ExecCommand)
	c.Args = args
	return c
}

func (c *ExecCommand) Run() error {
	var (
		target  = interfaces.ExampleKey(c.Args.Target)
		mapping = c.Args.Mapping
	)

	if v, ok := mapping[target]; ok {
		fmt.Printf("[Name] %q\n", target)
		if err := v(); err != nil {
			return &ExecError{
				Name: target,
				Err:  err,
			}
		}
	}

	return nil
}
