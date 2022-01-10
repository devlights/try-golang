package nomutex

import (
	"context"
	"time"

	"github.com/devlights/gomy/output"
)

const (
	execCount = 10000
)

var (
	balance = 1000
	execCh  = make(chan struct{}, execCount*2)
)

func deposit(v int) {
	balance += v
	execCh <- struct{}{}
}

func withdraw(v int) {
	balance -= v
	execCh <- struct{}{}
}

// NoMutex -- Mutexを利用しない場合のサンプルです.
func NoMutex() error {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 100*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	// 10 引き出して 10 預けるというのを非同期で 10000 回繰り返し
	for i := 0; i < execCount; i++ {
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
