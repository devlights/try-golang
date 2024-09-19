package main

import (
	"errors"
	"io"
	"log"
	"time"

	"github.com/devlights/try-golang/examples/singleapp/dev_shm/shm"
)

func main() {
	log.SetFlags(0)

	if err := run(); err != nil {
		log.Panic(err)
	}
}

func run() error {
	mw, err := shm.New("test2")
	if err != nil {
		return err
	}
	defer mw.Close()

	mr, err := shm.Open("test2")
	if err != nil {
		return err
	}
	defer mr.Close()

	ch := make(chan int)
	go func() {
		defer close(ch)

		for i := range 5 {
			if _, err := mw.Write([]byte("hello")); err != nil {
				log.Printf("hoge %s", err)
				break
			}

			ch <- i
			time.Sleep(1 * time.Second)
		}
	}()

	done := make(chan struct{})
	go func() {
		defer close(done)

		buf := make([]byte, 1<<6)
		for range ch {
			clear(buf)

			n, err := mr.Read(buf)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}

				log.Println(err)
				break
			}

			log.Printf("%s", buf[:n])
		}
	}()

	<-done

	return nil
}
