package tutorial09

import (
	"fmt"
	"time"
)

// ForLoop は、 Tour of Go - For (https://tour.golang.org/flowcontrol/1) の サンプルです。
func ForLoop() error {
	// ------------------------------------------------------------
	// Go言語の for は、 通常の for ループ と foreach ループ の両方を担う.
	// 書き方は、他の言語とほぼ同様. continue も break も同様.
	// Go言語には while が存在せず、 for のみで記載する.
	// ------------------------------------------------------------
	var (
		arr [10]int
	)

	// for ループ
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	fmt.Println(arr)

	// while ループ
	var (
		limit = 10
		count = 1
	)

	for count < limit {
		count++
	}

	fmt.Println(count, limit)

	// foreach ループ
	for i, v := range arr {
		fmt.Printf("[%d]:%v\t", i, v)
	}

	fmt.Println("")

	// 無限ループ
	var (
		// 2秒後にタイムアウトするチャネル
		timeout = time.After(2 * time.Second)
	)

loop:
	for {
		select {
		case <-timeout:
			break loop
		}
	}

	fmt.Println("break infinite loop")

	return nil
}
