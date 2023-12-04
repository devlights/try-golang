package syncs

import (
	"sync"

	"github.com/devlights/gomy/output"
)

type _PooledObject struct {
	v int
}

func NewPooledObject(v int) *_PooledObject {
	output.Stderrl("[New]", "Call NewPooledObject()")
	return &_PooledObject{v: v}
}

func (me *_PooledObject) Reset() {
	me.v = 0
}

func (me *_PooledObject) Set(v int) {
	me.v = v + 100
}

func (me *_PooledObject) Value() int {
	return me.v
}

// UsePool は、sync.Poolのサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/sync@go1.21.4#Pool
func UsePool() error {
	//
	// sync.Pool を利用する際の注意点 (go doc より引用)
	//
	// > Any item stored in the Pool may be removed automatically at any time without notification.
	// > If the Pool holds the only reference when this happens, the item might be deallocated.
	//
	// プールに保持されているオブジェクトは通知無しに自動削除される可能性がある。
	// その際に、プールがそのオブジェクトを参照する唯一であれば、メモリから
	// 開放される可能性がある。
	//
	// > A Pool must not be copied after first use.
	//
	// プールは、最初に使用した後はコピーしてはならない。
	//
	// > The Pool's New function should generally only return pointer types,
	// > since a pointer can be put into the return interface value without an allocation
	//
	// ポインタは割り当てなしで戻りインターフェイス値に入れることができるため、
	// プールの New 関数は通常、ポインタ型のみを返す必要がある.
	//

	const (
		NUM_ITEMS = 20
	)

	var (
		pool = sync.Pool{
			New: func() any {
				return NewPooledObject(0)
			},
		}
		ch   = make(chan int)
		done = make(chan struct{})
		wg   = sync.WaitGroup{}
	)

	wg.Add(NUM_ITEMS)

	for i := 0; i < NUM_ITEMS; i++ {
		go func(i int, ch chan<- int) {
			defer wg.Done()

			// プールから取得
			o := pool.Get().(*_PooledObject)

			o.Reset()
			o.Set(i)
			v := o.Value()

			// 使い終わったらプールに戻す
			pool.Put(o)

			ch <- v
		}(i, ch)
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	go func(done chan<- struct{}, ch <-chan int) {
		defer close(done)

		for v := range ch {
			output.Stderrl("[output]", v)
		}
	}(done, ch)

	<-done

	return nil
}
