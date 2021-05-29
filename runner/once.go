package runner

import "github.com/devlights/try-golang/mapping"

type (
	// Once -- 一度だけ実行するコマンド
	Once struct {
		Args *OnceArgs // 引数
	}

	// OnceArgs -- Once の引数データを表します.
	OnceArgs struct {
		ExecArgs // 引数
	}
)

// NewOnceArgs -- 新しい OnceArgs を生成して返します.
func NewOnceArgs(target string, m mapping.ExampleMapping) *OnceArgs {
	a := new(OnceArgs)
	a.Target = target
	a.Mapping = m
	return a
}

// NewOnce -- 新しい RunOnceCommand を生成して返します.
func NewOnce(args *OnceArgs) *Once {
	c := new(Once)
	c.Args = args
	return c
}

// Run -- 実行します.
func (c *Once) Run() error {
	var (
		target  = c.Args.Target
		mapping = c.Args.Mapping
	)

	execArgs := NewExecArgs(target, mapping)
	execCmd := NewExec(execArgs)

	if err := execCmd.Run(); err != nil {
		return err
	}

	return nil
}
