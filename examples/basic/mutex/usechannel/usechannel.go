package usechannel

import (
	"sync"

	"github.com/devlights/gomy/output"
)

const (
	execCount = 10000
)

var (
	balance = make(chan int, 1)
	countCh = make(chan struct{}, execCount*2)
)

func deposit(wg *sync.WaitGroup, v int) {
	defer wg.Done()

	balance <- <-balance + v
	countCh <- struct{}{}
}

func withdraw(wg *sync.WaitGroup, v int) {
	defer wg.Done()

	balance <- <-balance - v
	countCh <- struct{}{}
}

// UseChannel -- Mutexの代わりにチャネルを利用したサンプルです.
func UseChannel() error {
	var (
		wg sync.WaitGroup
	)
	wg.Add(execCount * 2)

	// 最初の値を設定
	balance <- 1000

	// 10 引き出して 10 預けるというのを非同期で 10000 回繰り返し
	for i := 0; i < 10000; i++ {
		go withdraw(&wg, 10)
		go deposit(&wg, 10)
	}

	wg.Wait()
	close(balance)
	close(countCh)

	var count int
	for range countCh {
		count++
	}

	output.Stdoutl("[execCount]", count)
	output.Stdoutl("[balance]", <-balance)

	return nil
}
