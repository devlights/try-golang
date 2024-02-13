package signals

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/devlights/gomy/output"
)

// Notify は、 signal.Notify のサンプルです.
//
// REFERENCES:
//   - https://golang.org/os/signal/#example_Notify
func Notify() error {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 5*time.Second)
	)

	defer mainCxl()
	defer procCxl()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	defer close(sigCh)

	select {
	case <-procCtx.Done():
		output.Stdoutl("[Timeout]", "procCtx.Done()")
	case <-sigCh:
		output.Stdoutl("[Interrupt]", "Ctrl-C")
	}

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: signal_notify

	   [Name] "signal_notify"
	   ^C[Interrupt]          Ctrl-C
	   task: Signal received: "interrupt"


	   [Elapsed] 1.771348899s
	*/

}
