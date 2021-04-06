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
		cmd *exec.Cmd
		out []byte
		err error
	)

	cmd = exec.CommandContext(procCtx, Shell, "-c", "sleep 1; echo ...done...")
	out, err = cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.Is(procCtx.Err(), context.DeadlineExceeded) && errors.As(err, &exitErr) {
			output.Stdoutf("[timeout]", "%[1]v(%[1]T)\n", err)
			return nil
		} else {
			return err
		}
	}

	output.Stdoutl("[output]", string(out))

	return nil
}
