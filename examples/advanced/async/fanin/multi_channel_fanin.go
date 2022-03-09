package fanin

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/devlights/gomy/chans"
	"github.com/devlights/gomy/output"
)

// MultiChannelFanIn -- chans.FanIn() を利用して処理するサンプルです。（入力チャネルが複数の場合)
func MultiChannelFanIn() error {
	var (
		numGoroutine   = 5                      // 並行起動するデータ取得ゴルーチンの数
		takeCount      = 2                      // データ取得チャネル毎にいくつデータを取得するかの数
		dataInInterval = 100 * time.Millisecond // データ投入時のインターバル
	)

	// コンテキストを生成
	var (
		rootCtx         = context.Background()
		mainCtx, cancel = context.WithCancel(rootCtx)
	)

	// 入力元となるチャネルと各ゴルーチンの終了判定チャネル
	var (
		srcCh      = make(chan interface{})
		doneChList = make([]<-chan struct{}, 0)
		takeChList = make([]<-chan interface{}, 0, numGoroutine)
	)

	defer close(srcCh)

	// srcCh に データを投入していくゴルーチン起動
	doneChList = append(doneChList, gen(mainCtx.Done(), srcCh, dataInInterval))

	// srcCh から takeCount個 取り出すチャネルを複数生成
	for i := 0; i < numGoroutine; i++ {
		takeDoneCh, takeCh := take(mainCtx.Done(), srcCh, i+1, takeCount)

		doneChList = append(doneChList, takeDoneCh)
		takeChList = append(takeChList, takeCh)
	}

	// 複数の取得チャネルを纏めてしまって、出力 (出力順序は問わない)
	for v := range chans.FanIn(mainCtx.Done(), takeChList...) {
		output.Stdoutl("[main]", v)
	}

	cancel()
	<-chans.WhenAll(doneChList...)

	return nil
}

func take(done <-chan struct{}, in <-chan interface{}, index int, takeCount int) (<-chan struct{}, <-chan interface{}) {
	terminated := make(chan struct{})
	out := make(chan interface{})
	go func() {
		// defer output.Stdoutf("take-goroutine", "%02d END\n", index)
		defer close(terminated)
		defer close(out)
		for v := range chans.Take(done, in, takeCount) {
			out <- fmt.Sprintf("[take-%02d] %v", index, v)
		}
	}()
	return terminated, out
}

func gen(done <-chan struct{}, in chan<- interface{}, interval time.Duration) <-chan struct{} {
	out := make(chan struct{})
	go func() {
		// defer output.Stdoutl("gen-goroutine", "END")
		defer close(out)

		randInt := func() interface{} {
			return rand.Int()
		}

		for v := range chans.RepeatFn(done, randInt) {
			in <- v
			<-time.After(interval)
		}
	}()
	return out
}
