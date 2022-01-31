package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// TimeSub は、 time.Sub() のサンプルです.
func TimeSub() error {
	var (
		ch1  = make(chan time.Time)
		ch2  = make(chan time.Time)
		exec = func(ch chan<- time.Time, delay time.Duration) {
			defer close(ch)
			time.Sleep(delay)
			ch <- time.Now()
		}
	)

	go exec(ch1, 100*time.Millisecond)
	go exec(ch2, 150*time.Millisecond)

	var (
		t1, t2  time.Time
		elapsed time.Duration
	)
LOOP:
	for {
		select {
		case t1 = <-ch1:
			ch1 = nil
		case t2 = <-ch2:
			ch2 = nil
		}

		if ch1 == nil && ch2 == nil {
			elapsed = t2.Sub(t1)
			break LOOP
		}
	}

	output.Stdoutl("[time.Sub]", elapsed)

	return nil
}
