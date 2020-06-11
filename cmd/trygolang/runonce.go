package main

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	// RunOnceCommand -- 一度だけ実行するコマンド
	RunOnceCommand struct {
		Args *RunOnceArgs // 引数
	}

	// RunOnceArgs -- RunOnceCommand の引数データを表します.
	RunOnceArgs struct {
		ExecArgs // 引数
	}
)

// NewRunOnceArgs -- 新しい RunOnceArgs を生成して返します.
func NewRunOnceArgs(target string, mapping mappings.ExampleMapping) *RunOnceArgs {
	a := new(RunOnceArgs)
	a.Target = target
	a.Mapping = mapping
	return a
}

// NewRunOnceCommand -- 新しい RunOnceCommand を生成して返します.
func NewRunOnceCommand(args *RunOnceArgs) *RunOnceCommand {
	c := new(RunOnceCommand)
	c.Args = args
	return c
}

// Run -- 実行します.
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
