package ordone

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/devlights/gomy/chans"
	"github.com/devlights/gomy/enumerable"
	"github.com/devlights/gomy/output"
)

// MultiInput -- chans.OrDone() を利用して処理するサンプルです。（入力チャネルが複数の場合)
func MultiInput() error {
	// 型のエイリアス
	type (
		DoneCh <-chan struct{}
		InCh   chan<- interface{}
		OutCh  <-chan interface{}
	)

	// 並行起動するゴルーチンの数とタイムリミット
	var (
		timeLimit    = 1 * time.Second
		numGoroutine = runtime.NumCPU()
	)

	// コンテキストとチャネル関連
	var (
		rootCtx    = context.Background()
		mainCtx, _ = context.WithTimeout(rootCtx, timeLimit)
		srcCh      = make(chan interface{})
		doneChList = make([]<-chan struct{}, 0, numGoroutine)
	)

	// -------------------------------------------------------
	// [Send] 100ms 毎に srcCh にデータを送り続けるゴルーチン起動
	// -------------------------------------------------------
	doneChList = append(doneChList, func(done DoneCh, in InCh) DoneCh {
		var (
			out = make(chan struct{})
		)

		go func(done DoneCh, in InCh) {
			defer close(out)

			// 値を返すクロージャ生成
			fn := func() func() interface{} {
				r := enumerable.NewRange(0, 100)
				return func() interface{} {
					r.Next()
					return r.Current()
				}
			}()

			for v := range chans.RepeatFn(done, fn) {
				in <- v
				<-time.After(100 * time.Millisecond)
			}
		}(done, in)

		return out
	}(mainCtx.Done(), srcCh))

	// -------------------------------------------------------
	// [Recv] srcCh からデータを取得するゴルーチンを複数起動
	// -------------------------------------------------------
	for i := 0; i < numGoroutine; i++ {
		doneChList = append(doneChList, func(done DoneCh, ch OutCh, index int) DoneCh {
			var (
				out  = make(chan struct{})
				name = fmt.Sprintf("[goroutine-%02d]", index)
			)

			go func(done DoneCh, dataCh OutCh, name string) {
				defer close(out)

				for v := range chans.OrDone(done, dataCh) {
					output.Stdoutl(name, v)
				}
			}(done, ch, name)

			return out
		}(mainCtx.Done(), srcCh, i))
	}

	// -------------------------------------------------------
	// [Join] メインコンテキストが完了するまで待機
	// -------------------------------------------------------
	<-chans.WhenAll(doneChList...)

	return nil
}
