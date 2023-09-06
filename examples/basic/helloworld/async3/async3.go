package async3

import (
	"log"
	"math/rand"
	"time"
)

var (
	count = 20
	delay = func() time.Duration {
		return time.Duration(rand.Intn(200)) * time.Millisecond
	}
)

// Async3 -- HelloWorld 非同期版 (3)
func Async3() error {
	log.SetFlags(0)

	var (
		hello = newRunner("hello", count, delay)
		world = newRunner("world", count, delay)
	)

	hello.run()
	world.run()

	var (
		h   string
		w   string
		hOk = true
		wOk = true
	)

	for {
		select {
		case h, hOk = <-hello.ch:
			if !hOk {
				break
			}
			log.Println(h)
		case w, wOk = <-world.ch:
			if !wOk {
				break
			}
			log.Println(w)
		}

		if !hOk && !wOk {
			break
		}
	}

	return nil
}
