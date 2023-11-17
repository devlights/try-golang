package cmdexec

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"runtime"
)

// Pipe は、 (*Cmd).StdinPipe,StdoutPipe,StderrPipeのサンプルです。
//
// REFERENCES:
//   - https://golang.org/os/exec/#example_Cmd_StdinPipe
func Pipe() error {
	if runtime.GOOS == "windows" {
		return errors.New("this example cannot run on Windows, sorry")
	}

	const (
		Shell = "/bin/bash"
	)

	var (
		cmd     *exec.Cmd      // コマンド
		fd0Pipe io.WriteCloser // 標準入力のパイプ
		fd1Pipe io.ReadCloser  // 標準出力のパイプ
		fd2Pipe io.ReadCloser  // 標準エラー出力のパイプ
	)

	// コマンド構築
	cmd = exec.Command(Shell, "-c", "tr a-z A-Z | sort; echo ...done... 1>&2")

	//
	// パイプを取得
	//   パイプを扱う場合は、それぞれを非同期で処理する必要がある。
	//
	fd0Pipe, _ = cmd.StdinPipe()
	fd1Pipe, _ = cmd.StdoutPipe()
	fd2Pipe, _ = cmd.StderrPipe()

	//
	// コマンド実行
	//   StdoutPipe または StderrPipe を利用する場合
	//   (*Cmd).Run() でコマンドを実行しない。
	//   (*Cmd).Start() で実行して (*Cmd).Wait() で待つようにする。
	//
	if err := cmd.Start(); err != nil {
		return err
	}

	// 標準入力のハンドリング
	go func() {
		//
		// 入力が完了したことを示すために明示的に Close する
		//
		defer fd0Pipe.Close()

		io.WriteString(fd0Pipe, "python\n")
		io.WriteString(fd0Pipe, "csharp\n")
		io.WriteString(fd0Pipe, "golang\n")
		io.WriteString(fd0Pipe, "java\n")
	}()

	// 標準出力のハンドリング
	go func() {
		//
		// 標準出力のパイプは (*Cmd).Wait() の呼び出しにて Close されるので
		// 通常呼ぶ必要はない。
		//
		scanner := bufio.NewScanner(fd1Pipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// 標準エラー出力のハンドリング
	go func() {
		//
		// 標準エラー出力のパイプは (*Cmd).Wait() の呼び出しにて Close されるので
		// 通常呼ぶ必要はない。
		//
		b, _ := io.ReadAll(fd2Pipe)
		fmt.Println(string(b))
	}()

	// コマンド終了待ち
	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: cmdexec_pipe

	   [Name] "cmdexec_pipe"
	   CSHARP
	   GOLANG
	   JAVA
	   PYTHON
	   ...done...



	   [Elapsed] 2.776069ms
	*/

}
