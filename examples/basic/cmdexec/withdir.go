package cmdexec

import (
	"errors"
	"os"
	"os/exec"
	"runtime"

	"github.com/devlights/gomy/output"
)

// WithDir -- *exec.Cmd 実行時にワーキングディレクトリを指定するサンプルです.
//
// REFERENCES
//   - https://dev.to/tobychui/quick-notes-for-go-os-exec-3ejg
func WithDir() error {
	if runtime.GOOS == "windows" {
		return errors.New("this example cannot run on Windows, sorry")
	}

	const (
		Shell = "/bin/bash"
	)

	var (
		cmd *exec.Cmd
		out []byte
		err error
	)

	output.Stdoutl("[cwd]", func() string { c, _ := os.Getwd(); return c }())

	//
	// プロセス実行時のワーキングディレクトリの指定なし
	//
	cmd = exec.Command(Shell, "-c", "pwd")

	out, err = cmd.Output()
	if err != nil {
		return err
	}

	output.Stdoutf("[no dir]", "%s", out)

	//
	// プロセス実行時のワーキングディレクトリの指定あり
	//
	cmd = exec.Command(Shell, "-c", "pwd")
	cmd.Dir = "/tmp"

	out, err = cmd.Output()
	if err != nil {
		return err
	}

	output.Stdoutf("[with dir]", "%s", out)

	return nil
}
