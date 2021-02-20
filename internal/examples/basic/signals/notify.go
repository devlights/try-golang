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
//   - https://golang.org/pkg/os/signal/#example_Notify
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
}
