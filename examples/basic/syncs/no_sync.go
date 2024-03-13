package syncs

import (
	"sync"

	"github.com/devlights/gomy/output"
)

// NoSync -- 同期なしで非同期処理をしているサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/sync@go1.17.8
//   - https://pkg.go.dev/sync/atomic@go1.17.8
//   - https://levelup.gitconnected.com/go-a-benchmark-to-compare-synchronization-techniques-ed73e118ec35
func NoSync() error {
	var (
		wg sync.WaitGroup
		x  int32 = 0
		fn       = func(minus bool) {
			defer wg.Done()
			for i := 0; i < 50000; i++ {
				if minus {
					x--
				} else {
					x++
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

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: syncs_no_sync

	   [Name] "syncs_no_sync"
	   [x]                  -7195
	   [x]                  3246
	   [x]                  15777
	   [x]                  12476
	   [x]                  15536


	   [Elapsed] 1.610119ms
	*/

}
