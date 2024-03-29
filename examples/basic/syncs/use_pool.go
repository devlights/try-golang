package syncs

import (
	"sync"

	"github.com/devlights/gomy/output"
)

type _PooledObject struct {
	v int
}

func newPooledObject(v int) *_PooledObject {
	output.Stderrl("[New]", "Call newPooledObject()")
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
		NumItems = 20
	)

	var (
		pool = sync.Pool{
			New: func() any {
				return newPooledObject(0)
			},
		}
		ch   = make(chan int)
		done = make(chan struct{})
		wg   = sync.WaitGroup{}
	)

	wg.Add(NumItems)

	for i := 0; i < NumItems; i++ {
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

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: syncs_use_pool

	   [Name] "syncs_use_pool"
	   [New]                Call newPooledObject()
	   [output]             113
	   [output]             114
	   [New]                Call newPooledObject()
	   [output]             115
	   [New]                Call newPooledObject()
	   [output]             116
	   [output]             117
	   [output]             118
	   [output]             119
	   [output]             103
	   [output]             100
	   [output]             101
	   [output]             108
	   [output]             106
	   [output]             107
	   [output]             104
	   [output]             109
	   [output]             110
	   [output]             112
	   [output]             105
	   [output]             111
	   [New]                Call newPooledObject()
	   [output]             102


	   [Elapsed] 489.3µs
	*/

}
