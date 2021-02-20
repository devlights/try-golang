package signals

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/devlights/gomy/output"
)

// NotifyContext は、 Go 1.16 から追加された signal.NotifyContext のサンプルです.
//
// REFERENCES:
//   - https://golang.org/pkg/os/signal/#NotifyContext
func NotifyContext() error {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 5*time.Second)
		sigCtx, sigCxl   = signal.NotifyContext(procCtx, os.Interrupt)
	)

	defer mainCxl()
	defer procCxl()
	defer sigCxl()

	<-sigCtx.Done()
	sigCxl()

	switch sigCtx.Err() {
	case context.Canceled:
		output.Stdoutl("[Interrupt]", "Ctrl-C")
	case context.DeadlineExceeded:
		output.Stdoutl("[Timeout]", "procCtx.Done()")
	}

	return nil
}
