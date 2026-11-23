package goroutines

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// WithContextAfterFunc は、context.AfterFunc のサンプルです。
//
// context.AfterFunc は、当該のコンテキストがキャンセル(タイムアウトも)された場合に呼び出される関数を登録できます。
// 登録した際にすでにコンテキストがキャンセル済みの場合は即座に呼び出されます。
//
// AfterFuncに登録した関数も非同期で実行されます。
// なので、登録関数の完了を待って何かの処理を行いたい場合はちゃんと待機する仕組みが必要です。
//
// 複数回の関数登録は上書きされる訳ではなく、それぞれ個別に保持されます。
// つまり、2回同じコンテキストに対して関数を登録したら2回呼ばれます。
//
// AfterFunc呼び出しの戻り値である stop は、登録解除のための関数です。
// 呼び出すと登録を解除できます。stopの戻り値が true の場合は AfterFunc に登録した関数の実行が阻止できたという意味。
// false の場合は阻止出来なかった、つまり、登録した関数がすでに起動済みか登録解除済みということになります。
//
// # REFERENCES
//   - https://pkg.go.dev/context@go1.27.4#AfterFunc
func WithContextAfterFunc() error {
	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithCancel(rootCtx)
	)
	defer cxl()

	var (
		stopDone = make(chan struct{})
		stop     = context.AfterFunc(ctx, func() {
			log.Println("AfterFunc 呼び出し")
			close(stopDone)
		})
		wg sync.WaitGroup
	)
	wg.Go(func() {
		time.Sleep(500 * time.Millisecond)
		cxl()
	})
	wg.Go(func() {
		var (
			t = time.NewTicker(100 * time.Millisecond)
			f = func() { _, _ = fmt.Fprint(os.Stderr, ".") }
		)
		defer t.Stop()

		for {
			select {
			case <-t.C:
				f()
			case <-stopDone:
				return
			}
		}
	})

	time.Sleep(1 * time.Second)
	if stop() {
		// true の場合は登録した関数を実行前に登録解除したという意味
		close(stopDone)
	} else {
		// false の場合は既に関数が起動済みなので完了を待機
		<-stopDone
		log.Println("AfterFunc 登録した関数の完了を確認")
	}

	wg.Wait()

	return nil
}
