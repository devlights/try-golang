package syncs

import (
	"sync"
	"time"

	"github.com/devlights/gomy/output"
)

// UseCondBroadcast -- sync.Cond.Broadcast() のサンプルです。
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
func UseCondBroadcast() error {
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

	// ブロードキャストを発行して、即座に全員に起きて処理してもらう
	time.Sleep(1 * time.Second)
	c.Broadcast()

	wg.Wait()

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: syncs_use_cond_broadcast

	   [Name] "syncs_use_cond_broadcast"
	   [begin ]             4
	   [begin ]             3
	   [begin ]             0
	   [begin ]             2
	   [begin ]             1
	   [signal]             1: h
	   [signal]             3: e
	   [signal]             4: l
	   [signal]             0: l
	   [signal]             2: o


	   [Elapsed] 1.000272946s
	*/

}
