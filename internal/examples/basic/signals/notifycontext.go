package signals

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/devlights/gomy/output"
)

// NotifyWithContext は、 Go 1.16 から追加された signal.NotifyContext のサンプルです.
//
// REFERENCES:
//   - https://golang.org/pkg/os/signal/#NotifyContext
func NotifyWithContext() error {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 5*time.Second)
	)

	defer mainCxl()
	defer procCxl()

	sigCtx, sigStop := signal.NotifyContext(procCtx, os.Interrupt)
	defer sigStop()

	select {
	case <-sigCtx.Done():
		switch e := sigCtx.Err(); e {
		case context.Canceled:
			output.Stdoutl("[Interrupt]", "Ctrl-C")
		case context.DeadlineExceeded:
			output.Stdoutl("[Timeout]", "procCtx.Done()")
		default:
			return e
		}
	}

	return nil
}
