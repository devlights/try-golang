package usemutex

import (
	"context"
	"sync"
	"time"

	"github.com/devlights/gomy/output"
)

const (
	execCount = 10000
)

var (
	mutex   sync.Mutex
	balance = 1000
	execCh  = make(chan struct{}, execCount*2)
)

func deposit(v int) {
	mutex.Lock()
	defer mutex.Unlock()

	balance += v
	execCh <- struct{}{}
}

func withdraw(v int) {
	mutex.Lock()
	defer mutex.Unlock()

	balance -= v
	execCh <- struct{}{}
}

// UseMutex -- NoMutexと同じ挙動で Mutex を使った版です.
func UseMutex() error {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 100*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	// 10 引き出して 10 預けるというのを非同期で 10000 回繰り返し
	for i := 0; i < 10000; i++ {
		go withdraw(10)
		go deposit(10)
	}

	<-procCtx.Done()
	close(execCh)

	var count int
	for range execCh {
		count++
	}

	output.Stdoutl("[execCount]", count)
	output.Stdoutl("[balance]", balance)

	return nil
}
