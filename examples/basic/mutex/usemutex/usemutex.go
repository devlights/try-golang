package usemutex

import (
	"sync"

	"github.com/devlights/gomy/output"
)

const (
	execCount = 10000
)

var (
	mutex   sync.Mutex
	balance = 1000
	countCh = make(chan struct{}, execCount*2)
)

func deposit(wg *sync.WaitGroup, v int) {
	defer wg.Done()

	mutex.Lock()
	defer mutex.Unlock()

	balance += v
	countCh <- struct{}{}
}

func withdraw(wg *sync.WaitGroup, v int) {
	defer wg.Done()

	mutex.Lock()
	defer mutex.Unlock()

	balance -= v
	countCh <- struct{}{}
}

// UseMutex -- NoMutexと同じ挙動で Mutex を使った版です.
func UseMutex() error {
	var (
		wg sync.WaitGroup
	)
	wg.Add(execCount * 2)

	// 10 引き出して 10 預けるというのを非同期で 10000 回繰り返し
	for i := 0; i < 10000; i++ {
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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: mutex_usemutex

	   [Name] "mutex_usemutex"
	   [execCount]          20000
	   [balance]            1000


	   [Elapsed] 11.125079ms
	*/

}
