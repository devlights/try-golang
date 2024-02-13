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
//   - https://golang.org/os/signal/#NotifyContext
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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: signal_notify_context

	   [Name] "signal_notify_context"
	   ^Ctask: Signal received: "interrupt"
	   [Interrupt]          Ctrl-C


	   [Elapsed] 1.113577647s
	*/

}
