package main

import (
	"fmt"
	"time"

	"github.com/devlights/try-golang/mappings"
)

type (
	ExecCommand struct {
		Args *ExecArgs
	}

	ExecArgs struct {
		Target  string
		Mapping mappings.ExampleMapping
	}
)

func NewExecArgs(target string, mapping mappings.ExampleMapping) *ExecArgs {
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
		target  = mappings.ExampleKey(c.Args.Target)
		mapping = c.Args.Mapping
	)

	if v, ok := mapping[target]; ok {
		fmt.Printf("[Name] %q\n", target)

		defer func(start time.Time) {
			fmt.Printf("\n\n[Elapsed] %s\n", time.Since(start))
		}(time.Now())

		if err := v(); err != nil {
			return &ExecError{
				Name: target,
				Err:  err,
			}
		}
	}

	return nil
}
