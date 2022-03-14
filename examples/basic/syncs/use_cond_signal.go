package syncs

import (
	"sync"
	"time"

	"github.com/devlights/gomy/output"
)

// UseCondSignal -- sync.Cond.Signal() のサンプルです。
//
//   - sync/atomic.AddXXX のサンプルは atomic_add.go を参照。
//   - 同期なしのサンプルは no_sync.go を参照。
//   - チャネルを使ったサンプルは use_channel.go を参照。
//
// # REFERENCES
//   - https://pkg.go.dev/sync@go1.17.8
//   - https://pkg.go.dev/sync/atomic@go1.17.8
//   - https://bit.ly/3i4X1ov
//   - https://mattn.kaoriya.net/software/lang/go/20140625223125.htm
func UseCondSignal() error {
	var (
		l  = len("hello")
		wg = sync.WaitGroup{}
		mu = sync.Mutex{}
		c  = sync.NewCond(&mu)
		ch = make(chan string, l)
	)

	// producer
	ch <- "h"
	ch <- "e"
	ch <- "l"
	ch <- "l"
	ch <- "o"

	// consumer
	wg.Add(l)
	for i := 0; i < l; i++ {
		go func(i int) {
			defer wg.Done()
			defer c.L.Unlock()
			
			output.Stderrf("[begin ]", "%d\n", i)
			c.L.Lock()
			c.Wait()
			output.Stderrf("[signal]", "%d: %v\n", i, <-ch)
		}(i)
	}

	// 1秒ごとに Signal() を発行し、順に起きていくことを確認
	for i := 0; i < l; i++ {
		time.Sleep(1 * time.Second)
		c.Signal()
	}

	wg.Wait()

	return nil
}
