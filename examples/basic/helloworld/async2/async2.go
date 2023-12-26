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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: helloworld_async2

	   [Name] "helloworld_async2"
	   0:world (66ms)
	   0:hello (189ms)
	   1:world (126ms)
	   2:world (113ms)
	   1:hello (143ms)
	   2:hello (16ms)
	   3:hello (7ms)
	   4:hello (90ms)
	   3:world (146ms)
	   5:hello (20ms)
	   4:world (44ms)
	   6:hello (54ms)
	   5:world (130ms)
	   7:hello (146ms)
	   6:world (108ms)
	   8:hello (89ms)
	   9:hello (152ms)
	   10:hello (1ms)
	   7:world (184ms)
	   11:hello (60ms)
	   12:hello (65ms)
	   8:world (178ms)
	   9:world (38ms)
	   13:hello (198ms)
	   14:hello (71ms)
	   10:world (190ms)
	   15:hello (134ms)
	   11:world (115ms)
	   12:world (44ms)
	   16:hello (120ms)
	   13:world (178ms)
	   17:hello (156ms)
	   18:hello (1ms)
	   14:world (160ms)
	   19:hello (117ms)
	   15:world (151ms)
	   16:world (91ms)
	   17:world (25ms)
	   18:world (131ms)
	   19:world (181ms)


	   [Elapsed] 2.408240175s
	*/

}
