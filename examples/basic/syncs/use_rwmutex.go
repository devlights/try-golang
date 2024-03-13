package syncs

import (
	"context"
	"log"
	"sync"
	"time"
)

// UseRWMutex は、sync.RWMutex のサンプルです。
//
// RWMutexは、読み取りと書き込みで別々にロックを取れるMutex。
//
// ドキュメントには以下のように記載されている。
//
// > A RWMutex is a reader/writer mutual exclusion lock. The lock can be held by an arbitrary number of readers or a single writer.
// The zero value for a RWMutex is an unlocked mutex.
//
// > (RWMutexは、リーダ/ライタ相互排他ロックである。このロックは、任意の数のリー ダーまたは1人のライターが保持することができる。
// RWMutexのゼロ値は、アンロックされたミューテックスである。)
//
// > If any goroutine calls Lock while the lock is already held by one or more readers,
// concurrent calls to RLock will block until the writer has acquired (and released) the lock,
// to ensure that the lock eventually becomes available to the writer.
// Note that this prohibits recursive read-locking.
//
// > ロックがすでに1つ以上のリーダによって保持されている間に、いずれかのゴルーチンがLockを呼び出すと、
// RLockの同時呼び出しは、ライタがロックを獲得（および解放）するまでブロックされ、ロックが最終的にライタが利用できるようになる。
// これは再帰的な読み取りロックを禁止していることに注意してください。
//
// RLockが読み取り用、Lockが書き込み用となる。RLockは複数のゴルーチンが取れるが、Lockは排他ロックとなる。
//
// # REFERENCES
//   - https://pkg.go.dev/sync@go1.22.1#RWMutex
func UseRWMutex() error {
	//
	// 3つのゴルーチンがあり、２つは値の読み取り(RLock)、１つは値の書き込み(Lock)を行う。
	//

	log.SetFlags(log.Lmicroseconds)

	var (
		v    int64
		m    sync.RWMutex
		read = func(ctx context.Context, prefix string, interval time.Duration) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}

				func() {
					m.RLock()
					defer m.RUnlock()

					log.Printf("[%s] %v", prefix, v)

					select {
					case <-ctx.Done():
					case <-time.After(interval):
					}
				}()
			}
		}
		write = func(ctx context.Context, prefix string, interval time.Duration) {
			for {
				// 1秒間、RLockできる猶予を与える
				select {
				case <-ctx.Done():
					return
				case <-time.After(1 * time.Second):
				}

				// ここから interval の間は排他ロックとなる
				func() {
					m.Lock()
					defer m.Unlock()

					old := v
					v++

					log.Printf("[%s] %v --> %v", prefix, old, v)

					select {
					case <-ctx.Done():
					case <-time.After(interval):
					}
				}()
			}
		}
	)

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	go read(ctx, "READ-1", 500*time.Millisecond)
	go read(ctx, "READ-2", 250*time.Millisecond)
	go write(ctx, "WRITE-1", 2*time.Second)

	<-ctx.Done()
	log.Println("DONE")

	return nil
}
