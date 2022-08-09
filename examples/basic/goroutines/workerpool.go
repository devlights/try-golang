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
//   - https://medium.com/@mertakar_22051/concurrency-in-golang-d49d2db1ed91
func WorkerPool() error {
	const (
		numItems = 45
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

	var (
		// 元値となる値を投入
		gen = func(out chan<- int) {
			defer close(out)
			for i := range iter.Range(numItems) {
				out <- i + 1
			}
		}
		// フィボナッチ数を算出
		calc = func(in <-chan int, out chan<- item) {
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
		}
		// 結果をフィルタリング
		filter = func(in <-chan item, out chan<- item) {
			defer close(out)
			for v := range in {
				if v.v >= 1000000 {
					out <- v
				}
			}
		}
	)

	// 仕事開始
	go gen(jobs)
	go calc(jobs, ret1)
	go filter(ret1, ret2)

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
