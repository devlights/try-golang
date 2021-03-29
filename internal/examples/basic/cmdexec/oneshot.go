package cmdexec

import (
	"bufio"
	"bytes"
	"errors"
	"os/exec"
	"runtime"

	"github.com/devlights/gomy/output"
)

// OneShot は、コマンドを一発実行して結果を取得するサンプルです.
//
// REFERENCES:
//   - https://stackoverflow.com/questions/19847594/how-to-reliably-detect-os-platform-in-go
//   - https://github.com/devlights/try-golang/issues/87
//   - https://stackoverflow.com/questions/31467153/golang-failed-exec-command-that-works-in-terminal
//   - https://github.com/github/hub/blob/2e002395b6a23fd2f51b9ed46e7d7581acd9dbd1/cmd/cmd.go#L40
func OneShot() error {
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
	// シェルの展開が必要ない場合は以下のようにそのまま指定して実行できる
	//

	out, _ = exec.Command("pwd").Output()
	output.Stdoutl("pwd", string(out))

	cmd = exec.Command("ls", "-l")
	out, err = cmd.Output()
	if err != nil {
		return err
	}

	output.Stdoutf("ls -l", "\n%s", string(out))
	output.StdoutHr()

	//
	// シェルの展開が必要な場合は sh -c または bash -c のようにシェル起動後にコマンド実行してもらう
	//

	cmd = exec.Command(Shell, "-c", "ls -l go.*")
	out, err = cmd.Output()
	if err != nil {
		return err
	}

	output.Stdoutf("ls -l go.*", "\n%s", string(out))
	output.StdoutHr()

	//
	// シェル起動後の実行は 普段のコマンド実行 と変わりない。パイプも指定できる
	//

	cmd = exec.Command(Shell, "-c", "ls -l | tail -n 3")
	out, err = cmd.Output()
	if err != nil {
		return err
	}

	output.Stdoutf("ls -l | tail -n 3", "\n%s", string(out))
	output.StderrHr()

	//
	// 結果は []byte で取得できているので、後で好きに加工できる
	//

	var (
		scanner  = bufio.NewScanner(bytes.NewReader(out))
		lastline string
	)

	for scanner.Scan() {
		lastline = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	output.Stdoutl("last line", lastline)

	return nil
}
