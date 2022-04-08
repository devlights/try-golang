package cmdexec

import (
	"os/exec"

	"github.com/devlights/gomy/output"
)

// WithSlice -- *exec.Cmd 実行時にスライスの値をコマンドの引数で指定するサンプルです.
//
// REFERENCES
//  - https://dev.to/tobychui/quick-notes-for-go-os-exec-3ejg
func WithSlice() error {
	var (
		cmd *exec.Cmd
		out []byte
		err error
	)

	var (
		p = []string{
			"hello",
			"world",
			"こんにちわ",
			"世界",
		}
	)

	cmd = exec.Command("echo", p...)

	out, err = cmd.CombinedOutput()
	if err != nil {
		output.Stdoutf("[cmd error]", "%s", out)
		return err
	}

	output.Stdoutf("[cmd]", "%s", out)

	return nil
}
