package leak

import (
	"time"

	"github.com/devlights/gomy/output"
)

// ForgottenSender -- goroutineリークが発生するパターンのサンプルです。
//
// チャネルを作成し、チャネルの送信側がいないパターン。
// 受信側のgoroutineが永遠に待ち続けるので終了しません。
//
// 解決方法としては、送信側が適切に使い終わったチャネルを閉じること。
//
// REFERENCES:
//   - https://betterprogramming.pub/common-goroutine-leaks-that-you-should-avoid-fe12d12d6ee
func ForgottenSender() error {
	var (
		fn = func(ch <-chan int) {
			data := <-ch
			output.Stdoutl("[goroutine leak]", data)
		}
		ch = make(chan int)
	)

	go fn(ch)

	//
	// チャネルにデータを送信するものがいないので
	// 上のgoroutineはプロセスが起動中は永遠に終了しません。
	//
	time.Sleep(1 * time.Second)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: goroutines_leak_forgotten_sender

	   [Name] "goroutines_leak_forgotten_sender"


	   [Elapsed] 1.001033553s
	*/

}
