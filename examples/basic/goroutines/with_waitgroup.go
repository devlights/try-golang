package goroutines

import (
	"sync"

	"github.com/devlights/gomy/output"
)

// WithWaitGroup -- sync.WaitGroupを用いて待ち合わせを行うパターンです.
func WithWaitGroup() error {
	var (
		wg sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		output.Stdoutl("[goroutine]", "This line is printed")
	}()

	wg.Wait()

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: goroutines_with_waitgroup

	   [Name] "goroutines_with_waitgroup"
	   [goroutine]          This line is printed


	   [Elapsed] 80.92µs
	*/

}
