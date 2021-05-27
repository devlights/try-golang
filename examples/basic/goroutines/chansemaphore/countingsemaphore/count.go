// Package countingsemaphore は、チャネルで計数セマフォを表現しています.
package countingsemaphore

import (
	chansem "github.com/devlights/try-golang/examples/basic/goroutines/chansemaphore"
)

type (
	impl chan struct{}
)

var _ chansem.Semaphore = (impl)(nil)

// New は、指定された計数で処理するセマフォを生成して返します.
func New(n int) chansem.Semaphore {
	ch := make(chan struct{}, n)
	return impl(ch)
}

func (me impl) Acquire() {
	me <- struct{}{}
}

func (me impl) Release() {
	<-me
}
