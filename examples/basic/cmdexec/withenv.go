package cmdexec

import (
	"errors"
	"os"
	"os/exec"
	"runtime"

	"github.com/devlights/gomy/output"
)

// WithEnv -- *exec.Cmd 実行時に追加の環境変数を指定するサンプルです.
//
// REFERENCES
//   - https://dev.to/tobychui/quick-notes-for-go-os-exec-3ejg
func WithEnv() error {
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

	//
	// 環境変数の追加なし
	//
	cmd = exec.Command(Shell, "-c", "env | grep LANG")

	out, err = cmd.Output()
	if err != nil {
		return err
	}

	output.Stdoutf("[no append]", "\n%s\n", out)
	output.StdoutHr()

	//
	// 環境変数の追加あり
	//
	cmd = exec.Command(Shell, "-c", "env | grep LANG")
	cmd.Env = append(os.Environ(), "LANG2=Japanese")

	out, err = cmd.Output()
	if err != nil {
		return err
	}

	output.Stdoutf("[append   ]", "\n%s\n", out)
	output.StdoutHr()

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: cmdexec_env

	   [Name] "cmdexec_env"
	   [no append]
	   LANG=en_US.UTF-8

	   --------------------------------------------------
	   [append   ]
	   LANG2=Japanese
	   LANG=en_US.UTF-8

	   --------------------------------------------------


	   [Elapsed] 6.235229ms
	*/

}
