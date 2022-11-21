package leak

import (
	"context"
	"time"

	"github.com/devlights/gomy/output"
)

// AbandonedSender -- goroutineリークが発生するパターンのサンプルです。
//
// チャネルの送受信の実装があるが、タイミングによっては受信側がいなくなってしまうパターン。
// 送信側のgoroutineが永遠に待ち続けるので終了しません。
//
// 解決方法としては、Bufferedなチャネルを使うこと。
//
// REFERENCES:
//   - https://betterprogramming.pub/common-goroutine-leaks-that-you-should-avoid-fe12d12d6ee
func AbandonedSender() error {
	var (
		ctx, cxl = context.WithTimeout(context.Background(), 10*time.Millisecond)
		ch       = make(chan int)
		iowait   = func() {
			time.Sleep(1 * time.Second)
		}
		fn = func(ch chan<- int) {
			iowait()
			ch <- 1
			output.Stdoutl("[send]", 1)
		}
	)
	defer cxl()

	go fn(ch)

	select {
	case v := <-ch:
		output.Stdoutl("[recv]", v)
	case <-ctx.Done():
	}

	//
	// チャネルからデータを受信するものがいなくなるので
	// 上のgoroutineはプロセスが起動中は永遠に終了しません。
	//
	time.Sleep(1 * time.Second)

	return nil
}
