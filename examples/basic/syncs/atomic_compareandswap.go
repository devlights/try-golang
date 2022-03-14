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
}
