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
					log.Printf("[%s] <<<", prefix)
					m.RLock()
					{
						log.Printf("[%s] %v", prefix, v)
					}
					m.RUnlock()
					log.Printf("[%s] >>>", prefix)

					select {
					case <-ctx.Done():
					case <-time.After(interval):
					}
				}()
			}
		}
		write = func(ctx context.Context, prefix string, interval time.Duration) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}

				func() {
					log.Printf("[%s] <<<", prefix)
					m.Lock()
					{
						old := v
						v++
						log.Printf("[%s] %v --> %v", prefix, old, v)
					}
					m.Unlock()
					log.Printf("[%s] >>>", prefix)

					select {
					case <-ctx.Done():
					case <-time.After(interval):
					}
				}()
			}
		}
	)

	ctx, cxl := context.WithTimeout(context.Background(), 2*time.Second)
	defer cxl()

	go read(ctx, "READ-1", 500*time.Millisecond)
	go read(ctx, "READ-2", 500*time.Millisecond)
	go write(ctx, "WRITE-1", 500*time.Millisecond)

	<-ctx.Done()
	log.Println("DONE")

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: syncs_use_rwmutex

	   [Name] "syncs_use_rwmutex"
	   05:57:00.082899 [WRITE-1] <<<
	   05:57:00.082972 [READ-2] <<<
	   05:57:00.083014 [WRITE-1] 0 --> 1
	   05:57:00.083022 [WRITE-1] >>>
	   05:57:00.083040 [READ-2] 1
	   05:57:00.083052 [READ-2] >>>
	   05:57:00.082938 [READ-1] <<<
	   05:57:00.083208 [READ-1] 1
	   05:57:00.083216 [READ-1] >>>
	   05:57:00.583383 [READ-1] <<<
	   05:57:00.583406 [READ-1] 1
	   05:57:00.583409 [READ-1] >>>
	   05:57:00.583422 [WRITE-1] <<<
	   05:57:00.583406 [READ-2] <<<
	   05:57:00.583490 [WRITE-1] 1 --> 2
	   05:57:00.583512 [WRITE-1] >>>
	   05:57:00.583518 [READ-2] 2
	   05:57:00.583523 [READ-2] >>>
	   05:57:01.083740 [READ-1] <<<
	   05:57:01.083786 [READ-1] 2
	   05:57:01.083801 [READ-1] >>>
	   05:57:01.083774 [READ-2] <<<
	   05:57:01.083829 [READ-2] 2
	   05:57:01.083833 [READ-2] >>>
	   05:57:01.083779 [WRITE-1] <<<
	   05:57:01.083837 [WRITE-1] 2 --> 3
	   05:57:01.083840 [WRITE-1] >>>
	   05:57:01.584523 [WRITE-1] <<<
	   05:57:01.584552 [READ-2] <<<
	   05:57:01.584553 [WRITE-1] 3 --> 4
	   05:57:01.584595 [WRITE-1] >>>
	   05:57:01.584607 [READ-2] 4
	   05:57:01.584617 [READ-2] >>>
	   05:57:01.584542 [READ-1] <<<
	   05:57:01.584623 [READ-1] 4
	   05:57:01.584629 [READ-1] >>>
	   05:57:02.083111 DONE


	   [Elapsed] 2.000292011s
	*/

}
