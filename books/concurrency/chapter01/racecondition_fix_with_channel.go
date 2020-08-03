package chapter01

import (
	"github.com/devlights/gomy/output"
)

// RaceConditionFixWithChannel -- チャネルを使い競合状態を回避するサンプルです。
//
// バッファ無しのchannel を利用することで、処理の流れを決定的にし、非同期処理が存在するが、常に同じ結果になるようにしています。
func RaceConditionFixWithChannel() error {
	var (
		inCh  = make(chan int)
		outCh = make(chan int)
	)

	go func() {
		outCh <- <-inCh + 1
	}()

	go func() {
		inCh <- 0
		close(inCh)
		close(outCh)
	}()

	output.Stdoutf("[result]", "value is %d\n", <-outCh)

	return nil
}
