package tutorial

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Select は、 Tour of Go - Select (https://tour.golang.org/concurrency/5) の サンプルです。
func Select() error {
	// ------------------------------------------------------------
	// select は、複数の goroutine を操作するために利用できる構文
	// switch と似ているが違うので注意。
	//
	// 通常、非同期処理にて複数のデータソースを準備が出来るまでブロックし
	// 準備が出来たものから実行していくというのはとても面倒な処理となる。
	//
	// Go言語のselectは、その負担をとても楽にしてくれる.
	//
	// 他の言語で同様の処理を記載するときと同じで、複数のcaseが同時に
	// 準備可能な場合、実行される case はランタイムによりランダムに選択される.
	//
	// よく利用するのは、メインでデータを扱うチャネルとタイムアウトを扱うチャネルの組合せとか
	// 所定条件で処理を打ち切るためのチャネルとの組合せなどで利用したりする。
	//
	// 書き方としては、基本は for の無限ループの中にselectを配置して
	// 特定条件で break するためのチャネルを case に入れておくというのが多い
	//
	// どの case も準備できていない場合は default case が実行される
	// ブロックせずに送受信する場合は、 default case を使う
	// ------------------------------------------------------------

	// ------------------------------------------------------------
	// 値を送信してくれる非同期処理を処理しながら、終わりを検知して抜ける
	// ------------------------------------------------------------
	channel1()

	// ------------------------------------------------------------
	// タイムアウト付きで、かつ、動かしている非同期処理の終了を待機して抜ける
	// (context.WithTimeout版)
	// ------------------------------------------------------------
	channel2()

	// ------------------------------------------------------------
	// タイムアウト付きで、かつ、動かしている非同期処理の終了を待機して抜ける
	// (time.After版)
	// ------------------------------------------------------------
	channel3()

	// ------------------------------------------------------------
	// どの case も準備できていない場合は、defaultが実行される
	// ------------------------------------------------------------
	channel4()

	// ------------------------------------------------------------
	// 非同期な処理を外から停止シグナル用チャネルを渡して停止させる
	// ------------------------------------------------------------
	channel5()

	return nil
}

// channel5 は、非同期な処理を外から停止シグナル用チャネルを渡して停止させるサンプルです.
func channel5() {
	type (
		nothing struct{}
	)

	var (
		done       = make(chan nothing)
		terminated <-chan nothing
	)

	fn := func(done <-chan nothing) <-chan nothing {
		var (
			terminated = make(chan nothing)
		)

		go func() {
			defer fmt.Println("fn exited")
			defer close(terminated)

			for {
				select {
				case <-done:
					return
				default:
				}

				fmt.Println("processing...")
				time.Sleep(1 * time.Second)
			}
		}()

		return terminated
	}

	terminated = fn(done)

	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	<-terminated
	fmt.Println("main done")
}

// channel4 は、どの case も準備できていない場合は、defaultが実行される場合のサンプルです.
func channel4() {
	var (
		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	)

	defer cancel()

loop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("終わり")
			break loop
		default:
			fmt.Println("まだ、どのcaseも準備できていない")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// channel3　は、タイムアウト付きで、かつ、動かしている非同期処理の終了を待機して抜ける場合のサンプルです.
// (time.After版)
func channel3() {
	var (
		wg       sync.WaitGroup
		c1       = make(chan int)
		c2       = make(chan int)
		quitX    = time.After(2 * time.Second)
		quitY    = time.After(2 * time.Second)
		quitMain = time.After(2 * time.Second)
		x, y     int
	)

	fn := func(c chan<- int, q <-chan time.Time, wg *sync.WaitGroup, wait time.Duration, prefix string) {
		defer wg.Done()

		for i := 0; true; i++ {
			select {
			case c <- i:
				fmt.Printf("%s<- %v\n", prefix, i)
			case <-q:
				fmt.Printf("%s ****** end ******\n", prefix)
				return
			}

			time.Sleep(wait)
		}
	}

	wg.Add(2)
	go fn(c1, quitX, &wg, 600*time.Millisecond, "f1")
	go fn(c2, quitY, &wg, 800*time.Millisecond, "f2")

loop:
	for {
		select {
		case x = <-c1:
			fmt.Printf("%v <-f1\n", x)
		case y = <-c2:
			fmt.Printf("%v <-f2\n", y)
		case <-quitMain:
			wg.Wait()
			break loop
		}
	}

	fmt.Printf("x:%v\ty:%v\n", x, y)
}

// channel2　は、タイムアウト付きで、かつ、動かしている非同期処理の終了を待機して抜ける場合のサンプルです.
// (context.WithTimeout版)
func channel2() {
	var (
		wg          sync.WaitGroup
		c1          = make(chan int)
		c2          = make(chan int)
		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		x, y        int
	)

	defer cancel()

	fn := func(c chan<- int, q <-chan struct{}, wg *sync.WaitGroup, wait time.Duration, prefix string) {
		defer wg.Done()

		for i := 0; true; i++ {
			select {
			case c <- i:
				fmt.Printf("%s<- %v\n", prefix, i)
			case <-q:
				fmt.Printf("%s ****** end ******\n", prefix)
				return
			}

			time.Sleep(wait)
		}
	}

	wg.Add(2)

	go fn(c1, ctx.Done(), &wg, 600*time.Millisecond, "f1")
	go fn(c2, ctx.Done(), &wg, 800*time.Millisecond, "f2")

loop:
	for {
		select {
		case x = <-c1:
			fmt.Printf("%v <-f1\n", x)
		case y = <-c2:
			fmt.Printf("%v <-f2\n", y)
		case <-ctx.Done():
			wg.Wait()
			break loop
		}
	}

	fmt.Printf("x:%v\ty:%v\n", x, y)
}

// channel1 は、値を送信してくれる非同期処理を処理しながら、終わりを検知して抜ける場合のサンプルです.
func channel1() {
	var (
		c1   = make(chan int)
		quit = make(chan struct{})
	)

	fn := func(c chan<- int, q chan<- struct{}) {
		for i := 0; i < 5; i++ {
			c <- i
		}

		q <- struct{}{}
	}

	go fn(c1, quit)

loop:
	for {
		select {
		case x := <-c1:
			fmt.Println(x)
		case <-quit:
			break loop
		}
	}
}
