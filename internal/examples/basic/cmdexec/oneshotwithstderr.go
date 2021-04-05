package cmdexec

import (
	"errors"
	"os/exec"
	"runtime"

	"github.com/devlights/gomy/output"
)

// OneShotWithStderr は、コマンドを一発実行して結果を取得するサンプルです。(標準エラー出力も含む)
//
// REFERENCES:
//   - https://golang.org/pkg/os/exec/#example_Cmd_CombinedOutput
//   - https://www.gnu.org/software/bash/manual/bash.html#Redirecting-Standard-Output-and-Standard-Error
//   - https://tldp.org/LDP/abs/html/io-redirection.html
func OneShotWithStderr() error {
	if runtime.GOOS == "windows" {
		return errors.New("this example cannot run on Windows, sorry")
	}

	const (
		Shell = "/bin/bash"
	)

	var (
		cmd *exec.Cmd // コマンド
		out []byte    // 実行結果
		err error     // 実行時エラー
	)

	//
	// (*Cmd).CombinedOutput() を使うと標準出力と標準エラー出力を結合した結果を取得できる
	// 以下では hello を標準出力に world を標準エラー出力に出力している
	//
	cmd = exec.Command(Shell, "-c", "echo hello; echo world 1>&2")
	out, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	output.Stdoutf("[CombinedOutput]", "\n%s", string(out))
	output.StdoutHr()

	//
	// Output だと、標準出力のみを対象とするので world は出力されない
	//
	cmd = exec.Command(Shell, "-c", "echo hello; echo world 1>&2")
	out, err = cmd.Output()
	if err != nil {
		return err
	}

	output.Stdoutf("[Output]", "\n%s", string(out))
	output.StdoutHr()

	return nil
}
