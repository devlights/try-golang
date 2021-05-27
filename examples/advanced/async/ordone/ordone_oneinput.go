package ordone

import (
	"context"
	"time"

	"github.com/devlights/gomy/chans"
	"github.com/devlights/gomy/output"
)

// OneInput -- chans.OrDone() を利用して処理するサンプルです。（入力チャネルが一つの場合)
func OneInput() error {
	// 3 秒後に終了するコンテキストを生成
	var (
		rootCtx         = context.Background()
		mainCtx, cancel = context.WithTimeout(rootCtx, 3*time.Second)
		inCh            = make(chan interface{})
	)

	defer cancel()
	defer close(inCh)

	// 1 秒毎にデータを送りつづけるチャネル生成
	subChDone := func(done <-chan struct{}, in chan<- interface{}) <-chan struct{} {
		out := make(chan struct{})
		go func() {
			defer close(out)

			for v := range chans.Repeat(done, 1) {
				in <- v
				<-time.After(1 * time.Second)
			}
		}()
		return out
	}(mainCtx.Done(), inCh)

	// データを出力
	//   この出力を実施している間にコンテキストのタイムアウトを迎えるため
	//   自動的にOrDone()の内部で終了判定となり、ループを抜けることになる。
	for v := range chans.OrDone(mainCtx.Done(), inCh) {
		output.Stdoutl("[main]", v)
	}

	// 各ゴルーチンの終了を待機
	<-chans.WhenAll(mainCtx.Done(), subChDone)

	return nil
}
