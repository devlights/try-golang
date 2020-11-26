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
}
