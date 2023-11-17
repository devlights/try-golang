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

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: cmdexec_oneshot

	   [Name] "cmdexec_oneshot"
	   pwd                  /workspace/try-golang

	   ls -l
	   total 14608
	   drwxr-xr-x  2 gitpod gitpod       61 Nov 17 01:36 builder
	   drwxr-xr-x  2 gitpod gitpod       36 Nov 17 01:36 cmd
	   -rw-r--r--  1 gitpod gitpod      637 Nov 17 01:36 Dockerfile
	   drwxr-xr-x 13 gitpod gitpod      176 Nov 17 01:36 examples
	   -rw-r--r--  1 gitpod gitpod      406 Nov 17 01:36 go.mod
	   -rw-r--r--  1 gitpod gitpod     3027 Nov 17 01:36 go.sum
	   -rw-r--r--  1 gitpod gitpod     1071 Nov 17 01:36 LICENSE
	   -rw-r--r--  1 gitpod gitpod       91 Nov 17 01:36 main.go
	   drwxr-xr-x  2 gitpod gitpod       76 Nov 17 01:36 mapping
	   -rw-r--r--  1 gitpod gitpod     3445 Nov 17 01:36 README.md
	   -rw-r--r--  1 gitpod gitpod      614 Nov 17 01:36 revive.toml
	   drwxr-xr-x  2 gitpod gitpod      165 Nov 17 01:36 runner
	   -rw-r--r--  1 gitpod gitpod      100 Nov 17 01:36 Taskfile_linux.yml
	   -rw-r--r--  1 gitpod gitpod       75 Nov 17 01:36 Taskfile_windows.yml
	   -rw-r--r--  1 gitpod gitpod     1411 Nov 17 01:36 Taskfile.yml
	   -rwxr-xr-x  1 gitpod gitpod 14914552 Nov 17 01:37 try-golang
	   --------------------------------------------------
	   ls -l go.*
	   -rw-r--r-- 1 gitpod gitpod  406 Nov 17 01:36 go.mod
	   -rw-r--r-- 1 gitpod gitpod 3027 Nov 17 01:36 go.sum
	   --------------------------------------------------
	   ls -l | tail -n 3
	   -rw-r--r--  1 gitpod gitpod       75 Nov 17 01:36 Taskfile_windows.yml
	   -rw-r--r--  1 gitpod gitpod     1411 Nov 17 01:36 Taskfile.yml
	   -rwxr-xr-x  1 gitpod gitpod 14914552 Nov 17 01:37 try-golang
	   --------------------------------------------------
	   last line            -rwxr-xr-x  1 gitpod gitpod 14914552 Nov 17 01:37 try-golang


	   [Elapsed] 10.481578ms
	*/

}
