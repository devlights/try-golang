package main

import (
	"io"
	"log"
	"sync"
	"time"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(io.Discard)

	if err := run(); err != nil {
		panic(err)
	}
}

// この関数は最適化が有効な場合、直接 v*2 になる可能性があり
// インライン化が有効な場合、インライン化される可能性がある
func calc(v int) int {
	v1 := v
	v2 := 2
	return v1 * v2
}

func run() error {
	const (
		COUNT   = 10000000
		WORKERS = 4
	)
	var (
		ch = make(chan int)
		wg sync.WaitGroup
	)

	// producer
	wg.Add(1)
	go func(ch chan<- int) {
		defer wg.Done()
		defer close(ch)

		for i := range COUNT {
			ch <- calc(i)
		}
	}(ch)

	time.Sleep(10 * time.Millisecond)

	// consumer
	for range WORKERS {
		wg.Add(1)
		go func(ch <-chan int) {
			for v := range ch {
				log.Println(v)
			}
		}(ch)
	}

	wg.Done()

	return nil
}
