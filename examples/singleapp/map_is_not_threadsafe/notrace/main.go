// マップ操作 はスレッドセーフでは無いというのを示すサンプルです。
// 本サンプルはデータ競合が発生しません。
//
// REFERENCES:
//   - https://stackoverflow.com/questions/44152988/append-not-thread-safe
//   - https://stackoverflow.com/questions/49879322/can-i-concurrently-write-different-slice-elements
package main

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var (
		verbose = flag.Bool("verbose", false, "verbose output")
	)
	flag.Parse()

	var (
		mu  sync.Mutex
		src = make(map[string]bool)
		dst = make(map[string]bool)
	)

	for i := 0; i < 5; i++ {
		src[strconv.Itoa(i)] = true
	}

	wg := sync.WaitGroup{}
	for k, v := range src {
		wg.Add(1)
		go func(k string, v bool) {
			defer wg.Done()

			mu.Lock()
			dst[k] = v
			mu.Unlock()
		}(k, v)
	}

	wg.Wait()

	fmt.Printf("src-len=%d\tdst-len=%d\n", len(src), len(dst))

	if *verbose {
		fmt.Println("=========== SRC ===========")
		for _, v := range src {
			fmt.Println(v)
		}
		fmt.Println("=========== DST ===========")
		for _, v := range dst {
			fmt.Println(v)
		}
	}
}
