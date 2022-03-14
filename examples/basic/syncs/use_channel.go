package syncs

import (
	"sync"

	"github.com/devlights/gomy/output"
)

// UseChannel -- 値の同期をチャネルを使って実現しているサンプルです。
//
//   - sync/atomic.AddXXX のサンプルは atomic_add.go を参照。
//   - 同期なしのサンプルは no_sync.go を参照。
//
// # REFERENCES
//   - https://pkg.go.dev/sync@go1.17.8
//   - https://pkg.go.dev/sync/atomic@go1.17.8
//   - https://levelup.gitconnected.com/go-a-benchmark-to-compare-synchronization-techniques-ed73e118ec35
func UseChannel() error {
	var (
		ch = make(chan int32, 1)
		wg sync.WaitGroup
		fn = func(minus bool) {
			defer wg.Done()
			for i := 0; i < 50000; i++ {
				x := <-ch
				if minus {
					x--
				} else {
					x++
				}
				ch <- x
			}
		}
	)
	defer close(ch)

	for i := 0; i < 5; i++ {
		ch <- 0

		wg.Add(2)
		go fn(true)
		go fn(false)
		wg.Wait()

		output.Stdoutl("[x]", <-ch)
	}

	return nil
}
