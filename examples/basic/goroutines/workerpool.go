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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: goroutines_workerpool

	   [Name] "goroutines_workerpool"
	   [v]                  fib(31)                 1346269    (12.336158ms)
	   [v]                  fib(32)                 2178309    (21.715015ms)
	   [v]                  fib(33)                 3524578    (80.962574ms)
	   [v]                  fib(34)                 5702887    (100.01898ms)
	   [v]                  fib(35)                 9227465    (197.626281ms)
	   [v]                  fib(36)                14930352    (308.72339ms)
	   [v]                  fib(37)                24157817    (412.192279ms)
	   [v]                  fib(38)                39088169    (611.75433ms)
	   [v]                  fib(39)                63245986    (906.788411ms)
	   [v]                  fib(40)               102334155    (1.282063608s)
	   [v]                  fib(41)               165580141    (1.830469841s)
	   [v]                  fib(42)               267914296    (2.537388482s)
	   [v]                  fib(43)               433494437    (3.453410792s)
	   [v]                  fib(44)               701408733    (5.059109086s)
	   [v]                  fib(45)              1134903170    (7.570757194s)


	   [Elapsed] 7.570912494s
	*/

}

func __fib(n int) int {
	if n <= 1 {
		return n
	}
	return __fib(n-1) + __fib(n-2)
}
