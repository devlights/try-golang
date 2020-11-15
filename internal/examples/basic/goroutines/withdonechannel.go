package goroutines

import "github.com/devlights/gomy/output"

// WithDoneChannel -- doneチャネルを用いて待ち合わせを行うサンプルです.
func WithDoneChannel() error {
	done := func() <-chan struct{} {
		done := make(chan struct{})

		go func() {
			defer close(done)
			output.Stdoutl("[goroutine]", "This line is printed")
		}()

		return done
	}()

	<-done

	return nil
}
