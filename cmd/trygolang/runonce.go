package main

import "github.com/devlights/try-golang/lib"

type (
	RunOnceCommand struct {
		Args *RunOnceArgs
	}

	RunOnceArgs struct {
		ExecArgs
	}
)

func NewRunOnceArgs(target string, mapping lib.SampleMapping) *RunOnceArgs {
	a := new(RunOnceArgs)
	a.Target = target
	a.Mapping = mapping
	return a
}

func NewRunOnceCommand(args *RunOnceArgs) *RunOnceCommand {
	c := new(RunOnceCommand)
	c.Args = args
	return c
}

func (c *RunOnceCommand) Run() error {
	var (
		target  = c.Args.Target
		mapping = c.Args.Mapping
	)

	execArgs := NewExecArgs(target, mapping)
	execCmd := NewExecCommand(execArgs)

	if err := execCmd.Run(); err != nil {
		return err
	}

	return nil
}
