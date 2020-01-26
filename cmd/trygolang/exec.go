package main

import (
	"fmt"
	"github.com/devlights/try-golang/lib"
)

type (
	ExecCommand struct {
		Args *ExecArgs
	}

	ExecArgs struct {
		Target  string
		Mapping lib.SampleMapping
	}
)

func NewExecArgs(target string, mapping lib.SampleMapping) *ExecArgs {
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
		target  = c.Args.Target
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
