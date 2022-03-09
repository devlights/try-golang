package leak

import (
	"time"

	"github.com/devlights/gomy/output"
)

// ForgottenReceiver -- goroutineリークが発生するパターンのサンプルです。
//
// チャネルを作成し、チャネルの受信側がいないパターン。
// 送信側のgoroutineが永遠に待ち続けるので終了しません。
//
// 解決方法としては、Bufferedなチャネルを使うこと。
//
// REFERENCES:
//   - https://betterprogramming.pub/common-goroutine-leaks-that-you-should-avoid-fe12d12d6ee
func ForgottenReceiver() error {
	var (
		fn = func(ch chan<- int) {
			ch <- 1
			output.Stdoutl("[goroutine leak]", 1)
		}
		ch = make(chan int)
	)

	go fn(ch)

	//
	// チャネルのデータを受信するものがいないので
	// 上のgoroutineはプロセスが起動中は永遠に終了しません。
	//
	time.Sleep(1 * time.Second)

	return nil
}
