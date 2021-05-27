package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// Timer は、time.NewTimer のサンプルです。
func Timer() error {
	var (
		timer = time.NewTimer(3 * time.Second)
		done  = func(limit time.Duration) <-chan struct{} {
			ch := make(chan struct{})
			go func() {
				defer close(ch)
				time.Sleep(limit)
			}()
			return ch
		}(5 * time.Second)
	)

LOOP:
	for {
		select {
		case <-timer.C:
			output.Stdoutl("timer.C", "timed out")
			break LOOP
		case <-done:
			output.Stdoutl("done", "proc done")
			break LOOP
		default:
			output.Stdoutl("loop", ".")
			time.Sleep(500 * time.Millisecond)
		}
	}

	return nil
}
