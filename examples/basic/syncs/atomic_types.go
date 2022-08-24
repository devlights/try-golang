package syncs

import (
	"sync"
	"sync/atomic"

	"github.com/devlights/gomy/output"
)

// AtomicTypes -- Go 1.19 から追加された sync/atomic パッケージ内の型についてのサンプルです。
//
// Go 1.18 までの atomic.AddXXXX() を使ったサンプルは atomic_add.go を参照。
//
// # REFERENCES
//   - https://dev.to/emreodabas/quick-guide-go-119-features-1j40
//   - https://go.dev/doc/go1.19#atomic_types
func AtomicTypes() error {
	var (
		wg sync.WaitGroup
		x  atomic.Int32 // atomic.Int32 のゼロ値は0
		fn = func(minus bool) {
			defer wg.Done()
			for i := 0; i < 50000; i++ {
				if minus {
					x.Add(-1)
				} else {
					x.Add(1)
				}
			}
		}
	)

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go fn(true)
		go fn(false)
		wg.Wait()

		output.Stdoutl("[x]", x.Load())
	}

	return nil
}
