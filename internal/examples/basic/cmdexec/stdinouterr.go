package cmdexec

import (
	"bytes"
	"errors"
	"io"
	"os/exec"
	"runtime"
	"strings"

	"github.com/devlights/gomy/output"
)

// Stdinouterr は、標準入力・標準出力・標準エラー出力を指定してコマンドを実行するサンプルです。
//
// REFERENCES:
//   - https://golang.org/pkg/os/exec/#example_Command
func Stdinouterr() error {
	if runtime.GOOS == "windows" {
		return errors.New("this example cannot run on Windows, sorry")
	}

	const (
		Shell = "/bin/bash"
	)

	var (
		cmd *exec.Cmd                                     // コマンド
		fd0 io.Reader = strings.NewReader("hello\nworld") // 標準入力
		fd1 io.Writer = &bytes.Buffer{}                   // 標準出力
		fd2 io.Writer = &bytes.Buffer{}                   // 標準エラー出力
		err error                                         // エラー
	)

	cmd = exec.Command(Shell, "-c", "tr a-z A-Z; echo ..done.. 1>&2")

	// 標準入力・標準出力・標準エラー出力を設定
	cmd.Stdin = fd0
	cmd.Stdout = fd1
	cmd.Stderr = fd2

	// 実行
	err = cmd.Run()
	if err != nil {
		return err
	}

	output.Stdoutf("[Stdout]", "\n%s\n", fd1)
	output.Stdoutf("[Stderr]", "\n%s\n", fd2)

	return nil
}
