package leak

import (
	"context"
	"time"

	"github.com/devlights/gomy/output"
)

// AbandonedReceiver -- goroutineリークが発生するパターンのサンプルです。
//
// チャネルの送受信の実装があるが、タイミングによっては送信側がいなくなってしまうパターン。
// 受信側のgoroutineが永遠に待ち続けるので終了しません。
//
// 解決方法としては、送信側が適切に使い終わったチャネルを閉じること。
//
// REFERENCES:
//   - https://betterprogramming.pub/common-goroutine-leaks-that-you-should-avoid-fe12d12d6ee
func AbandonedReceiver() error {
	var (
		ctx, cxl = context.WithTimeout(context.Background(), 10*time.Millisecond)
		ch       = make(chan int)
		iowait   = func() {
			time.Sleep(1 * time.Second)
		}
		fn = func(ch <-chan int) {
			iowait()
			data := <-ch
			output.Stdoutl("[recv]", data)
		}
	)
	defer cxl()

	go fn(ch)

	select {
	case ch <- 1:
		output.Stdoutl("[send]", 1)
	case <-ctx.Done():
	}

	//
	// チャネルにデータを送信するものがいなくなるので
	// 上のgoroutineはプロセスが起動中は永遠に終了しません。
	//
	time.Sleep(1 * time.Second)

	return nil
}
