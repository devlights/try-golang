package leak

import (
	"errors"

	"github.com/devlights/gomy/output"
)

// SenderAfterErrorCheck -- goroutineリークが発生するパターンのサンプルです。
//
// チャネルの送受信の実装があるが、内部の処理結果によっては送信側がいなくなってしまうパターン。
// 受信側のgoroutineが永遠に待ち続けるので終了しません。
//
// 解決方法としては、送信側が適切に使い終わったチャネルを閉じること。
//
// REFERENCES:
//   - https://betterprogramming.pub/common-goroutine-leaks-that-you-should-avoid-fe12d12d6ee
func SenderAfterErrorCheck() error {
	var (
		ch   = make(chan int)
		proc = func() bool {
			return false
		}
		fn = func(ch <-chan int) {
			data := <-ch
			output.Stdoutl("[recv]", data)
		}
	)

	go fn(ch)

	if !proc() {
		return errors.New("this is dummy error")
	}

	//
	// 上でエラーが発生した場合、以下は処理されない。
	// なので、上で起動しているgoroutineは永遠に受信待機することになる。
	//
	ch <- 1

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: goroutines_leak_sender_after_error_check

	   [Name] "goroutines_leak_sender_after_error_check"


	   [Elapsed] 20.16µs
	   [Error] this is dummy error (goroutines_leak_sender_after_error_check)
	*/

}
