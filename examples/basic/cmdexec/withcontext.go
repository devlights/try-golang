package cmdexec

import (
	"context"
	"errors"
	"os/exec"
	"runtime"
	"time"

	"github.com/devlights/gomy/output"
)

// WithContext は、context.Context 付きでコマンドを実行するサンプルです。
func WithContext() error {
	if runtime.GOOS == "windows" {
		return errors.New("this example cannot run on Windows, sorry")
	}

	const (
		Shell = "/bin/bash"
	)

	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 500*time.Millisecond)
	)

	defer mainCxl()
	defer procCxl()

	var (
		cmd *exec.Cmd // コマンド
		err error     // エラー
	)

	//
	// コマンドは２秒かかるようにして実行するが、渡している context は 500ms でタイムアウトする
	//
	cmd = exec.CommandContext(procCtx, Shell, "-c", "sleep 2")
	err = cmd.Run()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.Is(procCtx.Err(), context.DeadlineExceeded) && errors.As(err, &exitErr) {
			output.Stdoutf("[timeout]", "%[1]v(%[1]T)\n", err)
			return nil
		} else {
			return err
		}
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: cmdexec_withcontext

	   [Name] "cmdexec_withcontext"
	   [timeout]            signal: killed(*exec.ExitError)


	   [Elapsed] 501.067449ms
	*/

}
