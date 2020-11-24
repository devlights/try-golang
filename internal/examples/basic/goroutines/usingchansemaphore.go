package goroutines

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/internal/examples/basic/goroutines/chansemaphore"
	"github.com/devlights/try-golang/internal/examples/basic/goroutines/chansemaphore/binarysemaphore"
	"github.com/devlights/try-golang/internal/examples/basic/goroutines/chansemaphore/countingsemaphore"
)

// UsingChanSemaphore -- チャネルでセマフォの動作を行わせるサンプルです.
func UsingChanSemaphore() error {
	// バイナリセマフォを使ったサンプル
	output.StdoutHr()
	binsem()

	// 計数セマフォを使ったサンプル
	output.StdoutHr()
	cntsem()

	return nil
}

func cntsem() {
	// ユーティリティ
	var (
		iter = func(n int) []struct{} { return make([]struct{}, n) }
		now  = func() int64 { return time.Now().UTC().Unix() }
	)

	// ロガー
	var (
		mainLog = log.New(os.Stdout, "[main][cntsem] ", 0)
		semLog  = log.New(os.Stderr, ">>> [semaphore] >>> ", 0)
		gLog    = log.New(ioutil.Discard, "[goroutine] >>> >>> ", 0)
	)

	var (
		wg  sync.WaitGroup
		sem chansemaphore.Semaphore
	)

	sem = countingsemaphore.New(3)
	mainLog.Println("start", now())

	for i := range iter(5) {
		wg.Add(1)

		go func(no int) {
			defer wg.Done()

			gLog.Println(no, "start", now())
			sem.Acquire()

			semLog.Println(no, "acquire", now())
			gLog.Println(no, "processing....", now())

			time.Sleep(1 * time.Second)

			semLog.Println(no, "release", now())
			sem.Release()

			gLog.Println(no, "end  ", now())
		}(i)
	}

	wg.Wait()
	mainLog.Println("end  ", now())
}

func binsem() {
	// ユーティリティ
	var (
		iter = func(n int) []struct{} { return make([]struct{}, n) }
		now  = func() int64 { return time.Now().UTC().Unix() }
	)

	// ロガー
	var (
		mainLog = log.New(os.Stdout, "[main][binsem] ", 0)
		semLog  = log.New(os.Stderr, ">>> [semaphore] >>> ", 0)
		gLog    = log.New(ioutil.Discard, "[goroutine] >>> >>> ", 0)
	)

	var (
		wg  sync.WaitGroup
		sem chansemaphore.Semaphore
	)

	sem = binarysemaphore.New()
	mainLog.Println("start", now())

	for i := range iter(3) {
		wg.Add(1)

		go func(no int) {
			defer wg.Done()

			gLog.Println(no, "start", now())
			sem.Acquire()

			semLog.Println(no, "acquire", now())
			gLog.Println(no, "processing....", now())

			time.Sleep(1 * time.Second)

			semLog.Println(no, "release", now())
			sem.Release()

			gLog.Println(no, "end  ", now())
		}(i)
	}

	wg.Wait()
	mainLog.Println("end  ", now())
}
