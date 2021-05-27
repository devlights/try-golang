// Package binarysemaphore は、チャネルでバイナリセマフォを表現しています.
package binarysemaphore

import (
	chansem "github.com/devlights/try-golang/examples/basic/goroutines/chansemaphore"
)

type (
	impl chan struct{}
)

var _ chansem.Semaphore = (impl)(nil)

// New は、バイナリセマフォを生成して返します.
func New() chansem.Semaphore {
	ch := make(chan struct{}, 1)
	return impl(ch)
}

func (me impl) Acquire() {
	me <- struct{}{}
}

func (me impl) Release() {
	<-me
}
