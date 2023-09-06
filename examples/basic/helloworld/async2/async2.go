package async2

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// Async2 -- HelloWorld 非同期版 (2)
func Async2() error {
	log.SetFlags(0)

	var (
		wg    = new(sync.WaitGroup)
		ch    = make(chan string)
		delay = func() time.Duration {
			return time.Duration(rand.Intn(200)) * time.Millisecond
		}
	)

	var (
		hello   = newRunner("hello", wg, ch, delay)
		world   = newRunner("world", wg, ch, delay)
		closer  = newCloser(wg, ch)
		printer = newPrinter(ch)
	)

	wg.Add(2)

	go hello.run()
	go world.run()
	go closer.run()

	printer.run()

	return nil
}
