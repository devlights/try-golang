package main

import (
	"flag"
	"fmt"
	"io"
	"runtime"
	"sync"
)

func main() {
	var (
		loopcnt = flag.Int("loop", 10000, "loop count")
		incnt   = flag.Int("inch", 0, "input ch buffer count, 0 is unbuffered")
		outcnt  = flag.Int("outch", 0, "output ch buffer count, 0 is unbuffered")
	)
	flag.Parse()

	var (
		wg         sync.WaitGroup
		numWorkers = runtime.GOMAXPROCS(0)
		in         = make(chan int, *incnt)
		out        = make(chan string, *outcnt)
	)

	go func() {
		defer close(in)
		for i := range *loopcnt {
			in <- i
		}
	}()

	fmt.Printf("numWorkers=%d\n", numWorkers)
	wg.Add(numWorkers)

	for i := range numWorkers {
		go func(id int) {
			defer wg.Done()

			for j := range in {
				out <- fmt.Sprintf("[%d] hello world [%d]", id, j)
			}
		}(i + 1)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for v := range out {
		fmt.Fprintln(io.Discard, v)
	}

	fmt.Println("done")
}
