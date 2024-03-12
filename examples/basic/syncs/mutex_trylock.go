package syncs

import (
	"sync"
	"time"

	"github.com/devlights/gomy/output"
)

// MutexTryLock は、Go 1.18 で追加された mutex.TryLock() についてのサンプルです。
//
// > TryLock tries to lock m and reports whether it succeeded.
//
// > (TryLockはmをロックしようとし、成功したかどうかを報告します。)
//
// # REFERENCES
//
//   - https://pkg.go.dev/sync@go1.22.1#Mutex.TryLock
func MutexTryLock() error {
	const (
		WAIT_SECS = 5
	)

	var (
		m         sync.Mutex
		start     = make(chan struct{})
		lock5secs = func() {
			m.Lock()
			defer m.Unlock()

			close(start)
			time.Sleep(WAIT_SECS * time.Second)
		}
		printStatus = func() {
			output.Stdoutf("[TryLock]", "%s\tTryLock=%v\n", time.Now().Format(time.TimeOnly), m.TryLock())
		}
	)

	// 5秒間 mutex をロックし続ける
	go lock5secs()
	<-start

	for range WAIT_SECS {
		printStatus()
		time.Sleep(1 * time.Second)
	}

	// ここは mutex がアンロックされているのでロックが取れる
	printStatus()

	return nil
}
