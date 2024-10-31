package main

import (
	"context"
	"sync"
)

// CyclicBarrier は、指定された数のgoroutineが特定のポイントで待ち合わせることができる同期プリミティブです.
// すべてのgoroutineが到達するか、コンテキストがキャンセルされるまでブロックします.
type (
	CyclicBarrier struct {
		parties int        // 待ち合わせが必要なgoroutineの数
		waiting int        // 現在待機中のgoroutine数
		mutex   sync.Mutex // 内部状態を保護するためのmutex
		cond    *sync.Cond // 条件変数
		ctx     context.Context
		cancel  context.CancelFunc
		barrier chan struct{} // バリアチャネル
	}
)

// NewCyclicBarrier は、新しいCyclicBarrierを作成します.
// partiesには、同期が必要なgoroutineの数を指定します.
func NewCyclicBarrier(parties int) *CyclicBarrier {
	if parties <= 0 {
		panic("parties must be greater than 0")
	}

	var (
		ctx, cancel = context.WithCancel(context.Background())
		barrier     = &CyclicBarrier{
			parties: parties,
			ctx:     ctx,
			cancel:  cancel,
			barrier: make(chan struct{}),
		}
	)
	barrier.cond = sync.NewCond(&barrier.mutex)

	return barrier
}

// Await は、他のgoroutineが到達するのを待機します.
// すべてのgoroutineが到達すると、バリアが解放され、カウンターがリセットされます.
// コンテキストがキャンセルされた場合はエラーを返します.
func (me *CyclicBarrier) Await() error {
	me.mutex.Lock()
	defer me.mutex.Unlock()

	// コンテキストが既にキャンセルされているかチェック
	if me.ctx.Err() != nil {
		return me.ctx.Err()
	}

	var (
		generation = me.barrier // 現在の世代を記録（バリア条件が満了した場合、次のチャネルに切り替わるため）
	)
	me.waiting++
	if me.waiting == me.parties {
		// 最後のgoroutineが到達
		me.waiting = 0

		close(me.barrier)
		me.barrier = make(chan struct{}) // 新しい世代のためのチャネルを作成
		me.cond.Broadcast()              // 待機解除

		return nil
	}

	// 他のgoroutineを待つ
	for generation == me.barrier && me.ctx.Err() == nil {
		me.cond.Wait()
	}

	if me.ctx.Err() != nil {
		return me.ctx.Err()
	}

	return nil
}

// Reset は、バリアをリセットし、待機中のすべてのgoroutineをキャンセルします.
func (me *CyclicBarrier) Reset() {
	me.mutex.Lock()
	defer me.mutex.Unlock()

	// 現在待機しているgoroutineを解除
	me.cancel()

	var (
		ctx, cancel = context.WithCancel(context.Background())
	)
	me.ctx = ctx
	me.cancel = cancel
	me.waiting = 0

	// 世代入れ替え
	close(me.barrier)
	me.barrier = make(chan struct{})
	me.cond.Broadcast()
}

// GetNumberWaiting は、現在待機中のgoroutineの数を返します.
func (me *CyclicBarrier) GetNumberWaiting() int {
	me.mutex.Lock()
	defer me.mutex.Unlock()

	return me.waiting
}

// GetParties は、同期に必要なgoroutineの数を返します.
func (me *CyclicBarrier) GetParties() int {
	return me.parties
}
