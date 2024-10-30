package main

import (
	"sync"
	"sync/atomic"
)

// CountdownLatch は、C#-のCountdownEventやJavaのCountDownLatchと同様の機能を提供する構造体です.
type CountdownLatch struct {
	count int32
	mutex sync.Mutex
	cond  *sync.Cond
}

// NewCountdownLatch は、指定されたカウント数でCountdownLatchを初期化します.
func NewCountdownLatch(initialCount int) *CountdownLatch {
	if initialCount < 0 {
		panic("初期カウントは0以上である必要があります")
	}

	var (
		ce CountdownLatch
	)
	ce.count = int32(initialCount)
	ce.cond = sync.NewCond(&ce.mutex)

	return &ce
}

// Signal は、カウントを1減らします.
// 戻り値として、カウントダウンが満了したかどうかを返します.
func (me *CountdownLatch) Signal() bool {
	return me.SignalCount(1)
}

// SignalCount は、指定された数だけカウントを減らします.
// 戻り値として、カウントダウンが満了したかどうかを返します.
func (me *CountdownLatch) SignalCount(count int) bool {
	if count <= 0 {
		return false
	}

	me.mutex.Lock()
	defer me.mutex.Unlock()

	newCount := atomic.AddInt32(&me.count, -int32(count))
	if newCount <= 0 {
		me.cond.Broadcast()
		return true
	}

	return false
}

// Wait は、カウントが0になるまでブロックします.
func (me *CountdownLatch) Wait() {
	me.mutex.Lock()
	defer me.mutex.Unlock()

	for atomic.LoadInt32(&me.count) > 0 {
		me.cond.Wait()
	}
}

// CurrentCount は、現在のカウント値を返します.
func (me *CountdownLatch) CurrentCount() int {
	return int(atomic.LoadInt32(&me.count))
}

// Reset は、カウントを指定された値にリセットします.
func (me *CountdownLatch) Reset(count int) {
	if count < 0 {
		panic("リセットカウントは0以上である必要があります")
	}

	me.mutex.Lock()
	defer me.mutex.Unlock()

	atomic.StoreInt32(&me.count, int32(count))
}
