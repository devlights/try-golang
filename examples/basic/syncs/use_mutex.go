package syncs

import (
	"sync"

	"github.com/devlights/gomy/output"
)

// UseMutex -- sync.Mutex のサンプルです。
//
//   - sync/atomic.AddXXX のサンプルは atomic_add.go を参照。
//   - 同期なしのサンプルは no_sync.go を参照。
//   - チャネルを使ったサンプルは use_channel.go を参照。
//
// # REFERENCES
//   - https://pkg.go.dev/sync@go1.17.8
//   - https://pkg.go.dev/sync/atomic@go1.17.8
//   - https://levelup.gitconnected.com/go-a-benchmark-to-compare-synchronization-techniques-ed73e118ec35
func UseMutex() error {
	var (
		wg sync.WaitGroup
		m  sync.Mutex
		x  int32 = 0
		fn       = func(minus bool) {
			defer wg.Done()
			for i := 0; i < 50000; i++ {
				m.Lock()
				if minus {
					x--
				} else {
					x++
				}
				m.Unlock()
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
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: syncs_use_mutex

	   [Name] "syncs_use_mutex"
	   [x]                  0
	   [x]                  0
	   [x]                  0
	   [x]                  0
	   [x]                  0


	   [Elapsed] 4.79704ms
	*/

}
