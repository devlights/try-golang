package goroutines

import (
	"sync"
	"time"

	"github.com/devlights/gomy/iter"
	"github.com/devlights/gomy/output"
)

// WorkerPool -- Worker Pool パターンのサンプルです.
//
// REFENRECES
// 		- https://medium.com/@mertakar_22051/concurrency-in-golang-d49d2db1ed91
func WorkerPool() error {
	const (
		numItems = 40
	)

	type (
		item struct {
			i int
			v int
			t time.Duration
		}
	)

	var (
		jobs = make(chan int, numItems)
		ret1 = make(chan item, numItems)
		ret2 = make(chan item, 1)
	)

	// 元値となる値を投入
	go func(out chan<- int) {
		defer close(out)
		for i := range iter.Range(numItems) {
			out <- i + 1
		}
	}(jobs)

	// フィボナッチ数を算出
	go func(in <-chan int, out chan<- item) {
		defer close(out)
		wg := sync.WaitGroup{}
		for v := range in {
			wg.Add(1)
			go func(n int) {
				defer wg.Done()

				s := time.Now()
				v := __fib(n)

				out <- item{n, v, time.Since(s)}
			}(v)
		}
		wg.Wait()
	}(jobs, ret1)

	// 結果をフィルタリング
	go func(in <-chan item, out chan<- item) {
		defer close(out)
		for v := range in {
			if v.v >= 1000000 {
				out <- v
			}
		}
	}(ret1, ret2)

	// 結果出力
	for v := range ret2 {
		output.Stdoutf("[v]", "fib(%v)\t%+20v\t(%v)\n", v.i, v.v, v.t)
	}

	return nil
}

func __fib(n int) int {
	if n <= 1 {
		return n
	}
	return __fib(n-1) + __fib(n-2)
}
