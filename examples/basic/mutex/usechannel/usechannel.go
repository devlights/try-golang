package usechannel

import (
	"context"
	"time"

	"github.com/devlights/gomy/output"
)

const (
	execCount = 10000
)

var (
	balance = make(chan int, 1)
	execCh  = make(chan struct{}, execCount*2)
)

func deposit(v int) {
	// 本来は親のContextを使ってselectしながら処理するべきであるが割愛
	curr := <-balance
	next := curr + v
	balance <- next

	execCh <- struct{}{}
}

func withdraw(v int) {
	// 本来は親のContextを使ってselectしながら処理するべきであるが割愛
	curr := <-balance
	next := curr - v
	balance <- next

	execCh <- struct{}{}
}

// UseChannel -- Mutexの代わりにチャネルを利用したサンプルです.
func UseChannel() error {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 100*time.Millisecond)
	)
	defer mainCxl()
	defer procCxl()

	// 最初の値を設定
	balance <- 1000

	// 10 引き出して 10 預けるというのを非同期で 10000 回繰り返し
	for i := 0; i < 10000; i++ {
		go withdraw(10)
		go deposit(10)
	}

	<-procCtx.Done()
	close(execCh)
	close(balance)

	var count int
	for range execCh {
		count++
	}

	output.Stdoutl("[execCount]", count)
	output.Stdoutl("[balance]", <-balance)

	return nil
}
