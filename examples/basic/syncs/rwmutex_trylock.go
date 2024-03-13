package syncs

import (
	"context"
	"log"
	"sync"
	"time"
)

// RWMutexTryLock は、Go 1.18 で追加された RWMutex の TryLock() と TryRLock() のサンプルです。
//
// 使い方としては、Mutex.TryLock() と同じ。RLockのTry版がTryRLock, LockのTry版がTryLockとなる。
//
// # REFERENCES
//   - https://pkg.go.dev/sync@go1.22.1#RWMutex.TryLock
//   - https://pkg.go.dev/sync@go1.22.1#RWMutex.TryRLock
func RWMutexTryLock() error {
	//
	// 3つのゴルーチンがあり、２つは値の読み取り(RLock)、１つは値の書き込み(Lock)を行う。
	// それぞれ、ロックに入る前に TryRLock/TryLock により、ロック可能かどうかを確認する。
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
					for !m.TryRLock() {
						log.Printf("[%s] TryRLock=false", prefix)

						select {
						case <-ctx.Done():
							return
						case <-time.After(100 * time.Millisecond):
						}
					}

					log.Printf("[%s] %v", prefix, v)

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
					for !m.TryLock() {
						log.Printf("[%s] TryLock=false", prefix)

						select {
						case <-ctx.Done():
							return
						case <-time.After(100 * time.Millisecond):
						}
					}

					old := v
					v++
					log.Printf("[%s] %v --> %v", prefix, old, v)

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

	   ENTER EXAMPLE NAME: syncs_rwmutex_trylock

	   [Name] "syncs_rwmutex_trylock"
	   06:00:30.134950 [WRITE-1] <<<
	   06:00:30.135052 [WRITE-1] 0 --> 1
	   06:00:30.135059 [WRITE-1] >>>
	   06:00:30.134976 [READ-2] <<<
	   06:00:30.135079 [READ-2] 1
	   06:00:30.135085 [READ-2] >>>
	   06:00:30.134970 [READ-1] <<<
	   06:00:30.135264 [READ-1] 1
	   06:00:30.135268 [READ-1] >>>
	   06:00:30.635750 [READ-1] <<<
	   06:00:30.635773 [READ-1] 1
	   06:00:30.635776 [READ-1] >>>
	   06:00:30.635762 [READ-2] <<<
	   06:00:30.635803 [READ-2] 1
	   06:00:30.635812 [READ-2] >>>
	   06:00:30.635787 [WRITE-1] <<<
	   06:00:30.635855 [WRITE-1] 1 --> 2
	   06:00:30.635867 [WRITE-1] >>>
	   06:00:31.136388 [WRITE-1] <<<
	   06:00:31.136406 [READ-1] <<<
	   06:00:31.136411 [WRITE-1] 2 --> 3
	   06:00:31.136423 [WRITE-1] >>>
	   06:00:31.136422 [READ-2] <<<
	   06:00:31.136430 [READ-2] 3
	   06:00:31.136438 [READ-2] >>>
	   06:00:31.136415 [READ-1] TryRLock=false
	   06:00:31.236739 [READ-1] 3
	   06:00:31.236767 [READ-1] >>>
	   06:00:31.637314 [WRITE-1] <<<
	   06:00:31.637340 [READ-2] <<<
	   06:00:31.637367 [READ-2] TryRLock=false
	   06:00:31.637335 [WRITE-1] 3 --> 4
	   06:00:31.637405 [WRITE-1] >>>
	   06:00:31.737604 [READ-2] 4
	   06:00:31.737631 [READ-2] >>>
	   06:00:31.737631 [READ-1] <<<
	   06:00:31.737644 [READ-1] 4
	   06:00:31.737653 [READ-1] >>>
	   06:00:32.135034 DONE


	   [Elapsed] 2.000177751s
	*/

}
