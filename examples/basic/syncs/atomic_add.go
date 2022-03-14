package syncs

import (
	"sync"
	"sync/atomic"

	"github.com/devlights/gomy/output"
)

// AtomicAdd -- atomic.AddXXX のサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/sync@go1.17.8
//   - https://pkg.go.dev/sync/atomic@go1.17.8
//   - https://levelup.gitconnected.com/go-a-benchmark-to-compare-synchronization-techniques-ed73e118ec35
func AtomicAdd() error {
	var (
		wg sync.WaitGroup
		x  int32 = 0
		fn       = func(minus bool) {
			defer wg.Done()
			for i := 0; i < 50000; i++ {
				if minus {
					atomic.AddInt32(&x, -1)
				} else {
					atomic.AddInt32(&x, 1)
				}
			}
		}
	)

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go fn(true)
		go fn(false)
		wg.Wait()

		output.Stdoutl("[x]", x)
	}

	return nil
}
