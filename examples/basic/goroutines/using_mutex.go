package goroutines

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/devlights/gomy/chans"
)

// UsingMutex -- sync.Mutex を利用したサンプルです
func UsingMutex() error {
	// sync.Mutex は 排他制御 を行う際に利用できる.
	// カウントが 1 のセマフォと動きは似ているが、セマフォと違いミューテックスは
	// 所有権の概念を持つ。つまり、ロックを獲得したスレッドが、そのミューテックスを
	// 所有しているとなり、開放は所有者のみが行える.
	// セマフォには所有権の概念がないので、獲得したスレッドと別のスレッドで開放することも出来る
	const (
		goroutineCount = 5
	)

	var (
		iter = func(n int) []struct{} { return make([]struct{}, n) }
	)

	var (
		mainLog = log.New(os.Stdout, "[main] ", 0)
		gLog    = log.New(os.Stderr, "[goroutine] ", 0)
	)

	var (
		mu sync.Mutex
	)

	mainLog.Println("start")

	total := 0
	dones := make([]<-chan struct{}, 0, goroutineCount)
	for i := range iter(goroutineCount) {
		dones = append(dones, func(no int) <-chan struct{} {
			ch := make(chan struct{})

			go func() {
				defer func() {
					close(ch)
					gLog.Println(no, "done")
				}()

				mu.Lock()
				gLog.Println(no, "acquire mutex lock")

				// -------------------------------------
				// critical section
				// -------------------------------------
				gLog.Println(no, "processing...")
				total++

				gLog.Println(no, "release mutex lock")
				mu.Unlock()

				time.Sleep(10 * time.Millisecond)
			}()

			return ch
		}(i))
	}

	<-chans.WhenAll(dones...)

	mainLog.Printf("total=%d", total)
	mainLog.Println("done")

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: goroutines_using_mutex

	   [Name] "goroutines_using_mutex"
	   [main] start
	   [goroutine] 1 acquire mutex lock
	   [goroutine] 1 processing...
	   [goroutine] 1 release mutex lock
	   [goroutine] 0 acquire mutex lock
	   [goroutine] 0 processing...
	   [goroutine] 0 release mutex lock
	   [goroutine] 3 acquire mutex lock
	   [goroutine] 3 processing...
	   [goroutine] 3 release mutex lock
	   [goroutine] 4 acquire mutex lock
	   [goroutine] 4 processing...
	   [goroutine] 4 release mutex lock
	   [goroutine] 2 acquire mutex lock
	   [goroutine] 2 processing...
	   [goroutine] 2 release mutex lock
	   [goroutine] 2 done
	   [goroutine] 1 done
	   [goroutine] 0 done
	   [goroutine] 4 done
	   [main] total=5
	   [goroutine] 3 done
	   [main] done


	   [Elapsed] 10.969168ms
	*/

}
