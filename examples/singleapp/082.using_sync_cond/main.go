package main

import (
	"flag"
	"log"
	"sync"
	"time"
)

var (
	useSignal1   = flag.Bool("signal", false, "use cond.Signal (one time)")
	useSignal2   = flag.Bool("signal2", false, "use cond.Signal (multiple times)")
	useBroadCast = flag.Bool("broadcast", false, "use cond.Broadcast")
)

func init() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	switch {
	case *useSignal1:
		if err := signal1(); err != nil {
			return err
		}
	case *useSignal2:
		if err := signal2(); err != nil {
			return err
		}
	case *useBroadCast:
		if err := broadcast(); err != nil {
			return err
		}
	}

	return nil
}

func signal1() error {
	var (
		mu   sync.Mutex
		cond = sync.NewCond(&mu)
		done = false
	)

	go func() {
		// 時間のかかる処理を模擬
		log.Println("時間のかかる処理 開始")
		{
			time.Sleep(3 * time.Second)
		}
		log.Println("時間のかかる処理 完了")

		mu.Lock()
		{
			// 完了を通知
			log.Println("完了を通知 (Cond.Signal)")
			done = true
			cond.Signal()
		}
		mu.Unlock()
	}()

	mu.Lock() // Cond.Wait()は待機に入る前にロックを解除するので事前にロックしておく必要がある。
	for !done {
		// シグナルが通知されるまで待機
		log.Println("待機 開始 (Cond.Wait)")
		cond.Wait()
		log.Println("待機 終了 (Cond.Wait)")
	}
	mu.Unlock() // Cond.Wait()は待機から復帰した際にロックを獲得しているので解除する必要がある。

	return nil
}

func signal2() error {
	var (
		mu   sync.Mutex
		cond = sync.NewCond(&mu)
		wg   sync.WaitGroup
	)
	wg.Add(3)

	// 1人が処理を先行して行っていて、それが終わったら
	// 後続の3人が動き出すというシナリオとする。

	for i := 0; i < 3; i++ {
		i := i
		go func() {
			defer wg.Done()

			mu.Lock()
			{
				log.Printf("ワーカー%02d Wait", i)
				cond.Wait()
				log.Printf("ワーカー%02d Wake Up", i)
			}
			mu.Unlock()

			log.Printf("ワーカー%02d Start", i)
			time.Sleep(1 * time.Second)
			log.Printf("ワーカー%02d End", i)
		}()
	}

	// 時間のかかる処理を模擬
	log.Println("先行処理 開始")
	{
		time.Sleep(3 * time.Second)
	}
	log.Println("先行処理 完了")

	for i := 0; i < 3; i++ {
		mu.Lock()
		{
			log.Println("完了を通知")
			cond.Signal()
		}
		mu.Unlock()

		time.Sleep(200 * time.Millisecond)
	}

	wg.Wait()

	return nil
}

func broadcast() error {
	var (
		mu   sync.Mutex
		cond = sync.NewCond(&mu)
		wg   sync.WaitGroup
	)
	wg.Add(3)

	// 1人が処理を先行して行っていて、それが終わったら
	// 後続の3人が動き出すというシナリオとする。

	for i := 0; i < 3; i++ {
		i := i
		go func() {
			defer wg.Done()

			mu.Lock()
			{
				log.Printf("ワーカー%02d Wait", i)
				cond.Wait()
				log.Printf("ワーカー%02d Wake Up", i)
			}
			mu.Unlock()

			log.Printf("ワーカー%02d Start", i)
			time.Sleep(1 * time.Second)
			log.Printf("ワーカー%02d End", i)
		}()
	}

	// 時間のかかる処理を模擬
	log.Println("先行処理 開始")
	{
		time.Sleep(3 * time.Second)
	}
	log.Println("先行処理 完了")

	mu.Lock()
	{
		log.Println("完了をブロードキャスト")
		cond.Broadcast()
	}
	mu.Unlock()

	wg.Wait()

	return nil
}
