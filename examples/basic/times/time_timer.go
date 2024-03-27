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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_timer

	   [Name] "time_timer"
	   loop                 .
	   loop                 .
	   loop                 .
	   loop                 .
	   loop                 .
	   loop                 .
	   timer.C              timed out


	   [Elapsed] 3.004293004s
	*/

}
