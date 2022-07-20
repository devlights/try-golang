// スライス操作 (スライスヘッダの書き換え）はスレッドセーフでは無いというのを示すサンプルです。
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

type data struct {
	value string
}

func (me data) String() string {
	return me.value
}

func main() {
	var (
		verbose = flag.Bool("verbose", false, "verbose output")
	)
	flag.Parse()

	var (
		src = make([]data, 100)
		dst = make([]data, len(src))
	)

	for i := 0; i < len(src); i++ {
		src[i].value = strconv.Itoa(i)
	}

	wg := sync.WaitGroup{}
	for i, v := range src {
		wg.Add(1)
		go func(i int, v data) {
			defer wg.Done()

			var tmp data
			tmp.value = v.value

			dst[i] = tmp
		}(i, v)
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
