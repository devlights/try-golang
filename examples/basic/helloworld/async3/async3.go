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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: helloworld_async3

	   [Name] "helloworld_async3"
	   0:world (88ms)
	   0:hello (120ms)
	   1:hello (68ms)
	   1:world (179ms)
	   2:world (85ms)
	   2:hello (195ms)
	   3:hello (34ms)
	   3:world (85ms)
	   4:hello (48ms)
	   4:world (34ms)
	   5:hello (68ms)
	   5:world (113ms)
	   6:hello (56ms)
	   7:hello (1ms)
	   6:world (26ms)
	   8:hello (195ms)
	   9:hello (18ms)
	   7:world (194ms)
	   8:world (69ms)
	   10:hello (91ms)
	   9:world (32ms)
	   10:world (6ms)
	   11:world (41ms)
	   11:hello (128ms)
	   12:world (79ms)
	   12:hello (55ms)
	   13:hello (36ms)
	   14:hello (0s)
	   15:hello (55ms)
	   13:world (161ms)
	   16:hello (169ms)
	   14:world (180ms)
	   15:world (101ms)
	   16:world (1ms)
	   17:hello (137ms)
	   17:world (16ms)
	   18:hello (103ms)
	   18:world (99ms)
	   19:world (17ms)
	   19:hello (198ms)


	   [Elapsed] 1.783459428s
	*/

}
