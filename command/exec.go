package command

import (
	"fmt"
	"time"

	"github.com/devlights/try-golang/mapping"
)

type (
	// ExecCommand -- 処理実行を行うコマンド
	ExecCommand struct {
		Args *ExecArgs // 引数
	}

	// ExecArgs -- ExecCommand の 引数データ を表します.
	ExecArgs struct {
		Target  string                 // 対象
		Mapping mapping.ExampleMapping // マッピング情報
	}
)

// NewExecArgs -- 新しい ExecArgs を生成して返します.
func NewExecArgs(target string, m mapping.ExampleMapping) *ExecArgs {
	a := new(ExecArgs)
	a.Target = target
	a.Mapping = m
	return a
}

// NewExecCommand -- 新しい ExecCommand を生成して返します.
func NewExecCommand(args *ExecArgs) *ExecCommand {
	c := new(ExecCommand)
	c.Args = args
	return c
}

// Run -- 実行します.
func (c *ExecCommand) Run() error {
	var (
		target  = mapping.ExampleKey(c.Args.Target)
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
