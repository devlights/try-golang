package syncs

import (
	"sync"
	"sync/atomic"

	"github.com/devlights/gomy/output"
)

// CompareAndSwap -- sync/atomic.CompareAndSwap のサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/sync@go1.17.8
//   - https://pkg.go.dev/sync/atomic@go1.17.8
func CompareAndSwap() error {
	var (
		wg sync.WaitGroup
		x  int32 = 0
		fn       = func(minus bool) {
			defer wg.Done()

			results := make([]bool, 0)
			for i := 0; i < 50000; i++ {
				//
				// 値が変化していなければ、x-1 or x+1 を実施
				// atomic.CompareAndSwapXXX の結果が false になった場合は
				// ループに入った時点の値と実際の値変更までの 「ほんの僅か」 な時間の間で
				// 別のゴルーチンによって値が書き換えられたということ。
				//

				y := x
				if minus {
					results = append(results, atomic.CompareAndSwapInt32(&x, y, x-1))
				} else {
					results = append(results, atomic.CompareAndSwapInt32(&x, y, x+1))
				}
			}

			var tCount, fCount int
			for _, v := range results {
				if v {
					tCount++
				} else {
					fCount++
				}
			}

			output.Stdoutf("[results]", "minus=%-5v\tswapped:%d\tswapfail:%d\n", minus, tCount, fCount)
		}
	)

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go fn(true)
		go fn(false)
		wg.Wait()

		output.Stdoutl("[x]", x)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: syncs_atomic_compare_and_swap

	   [Name] "syncs_atomic_compare_and_swap"
	   [results]            minus=false        swapped:42078   swapfail:7922
	   [results]            minus=true         swapped:44293   swapfail:5707
	   [x]                  -2215
	   [results]            minus=true         swapped:40082   swapfail:9918
	   [results]            minus=false        swapped:44154   swapfail:5846
	   [x]                  1857
	   [results]            minus=true         swapped:33961   swapfail:16039
	   [results]            minus=false        swapped:37622   swapfail:12378
	   [x]                  5518
	   [results]            minus=true         swapped:36821   swapfail:13179
	   [results]            minus=false        swapped:40094   swapfail:9906
	   [x]                  8791
	   [results]            minus=false        swapped:40001   swapfail:9999
	   [results]            minus=true         swapped:38971   swapfail:11029
	   [x]                  9821


	   [Elapsed] 7.267749ms
	*/

}
