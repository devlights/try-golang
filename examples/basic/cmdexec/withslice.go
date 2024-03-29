package cmdexec

import (
	"fmt"
	"os/exec"

	"github.com/devlights/gomy/output"
)

// WithSlice -- *exec.Cmd 実行時にスライスの値をコマンドの引数で指定するサンプルです.
//
// REFERENCES
//   - https://dev.to/tobychui/quick-notes-for-go-os-exec-3ejg
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
		return fmt.Errorf("%w (%s)", err, out)
	}

	output.Stdoutf("[cmd]", "%s", out)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: cmdexec_slice

	   [Name] "cmdexec_slice"
	   [cmd]                hello world こんにちわ 世界


	   [Elapsed] 5.243239ms
	*/

}
