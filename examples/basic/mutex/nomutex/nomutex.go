package nomutex

import (
	"sync"

	"github.com/devlights/gomy/output"
)

const (
	execCount = 10000
)

var (
	balance = 1000
	countCh = make(chan struct{}, execCount*2)
)

func deposit(wg *sync.WaitGroup, v int) {
	defer wg.Done()

	balance += v
	countCh <- struct{}{}
}

func withdraw(wg *sync.WaitGroup, v int) {
	defer wg.Done()

	balance -= v
	countCh <- struct{}{}
}

// NoMutex -- Mutexを利用しない場合のサンプルです.
func NoMutex() error {
	var (
		wg sync.WaitGroup
	)
	wg.Add(execCount * 2)

	// 10 引き出して 10 預けるというのを非同期で 10000 回繰り返し
	for i := 0; i < execCount; i++ {
		go withdraw(&wg, 10)
		go deposit(&wg, 10)
	}

	wg.Wait()
	close(countCh)

	var count int
	for range countCh {
		count++
	}

	output.Stdoutl("[execCount]", count)
	output.Stdoutl("[balance]", balance)

	return nil
}
