// デッドロックのサンプルです.
//
// 概要
//
//	実行すると以下のような結果となります。
//
//		$ go run cmd/deadlock/main.go
//		[v]                  0
//		[v]                  1
//		[v]                  2
//		[v]                  3
//		[v]                  4
//		fatal error: all goroutines are asleep - deadlock!
//		goroutine 1 [chan receive]:
//		main.main()
//		        /workspace/try-golang/cmd/deadlock/main.go:25 +0xa6
//		exit status 2
//
// 参考情報
//   - https://medium.com/@mertakar_22051/concurrency-in-golang-d49d2db1ed91
//   - https://stackoverflow.com/questions/54157836/a-simple-example-about-go-channel-with-deadlock-and-why
package main

import (
	"github.com/devlights/gomy/iter"
	"github.com/devlights/gomy/output"
)

func main() {
	var (
		ch = make(chan int)
	)

	go func() {
		for i := range iter.Range(5) {
			ch <- i
		}
	}()

	for range iter.Range(6) {
		v := <-ch
		output.Stdoutl("[v]", v)
	}

	output.Stdoutl("[done]")
}
