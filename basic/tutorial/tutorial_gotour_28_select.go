package tutorial

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func GoTourSelect() error {
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
	var (
		c1   = make(chan int)
		quit = make(chan struct{})
	)

	func1 := func(c chan<- int, q chan<- struct{}) {
		for i := 0; i < 5; i++ {
			c <- i
		}

		q <- struct{}{}
	}

	go func1(c1, quit)

loop:
	for {
		select {
		case x := <-c1:
			fmt.Println(x)
		case <-quit:
			break loop
		}
	}

	// ------------------------------------------------------------
	// タイムアウト付きで、かつ、動かしている非同期処理の終了を待機して抜ける
	// ------------------------------------------------------------
	var (
		wg          sync.WaitGroup
		c2          = make(chan int)
		c3          = make(chan int)
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		x, y        int
	)

	defer cancel()

	func2 := func(c chan<- int, q <-chan struct{}, wg *sync.WaitGroup, wait time.Duration, prefix string) {
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
	go func2(c2, ctx.Done(), &wg, 200*time.Millisecond, "f1")
	go func2(c3, ctx.Done(), &wg, 300*time.Millisecond, "f2")

loop2:
	for {
		select {
		case x = <-c2:
			fmt.Printf("%v <-f1\n", x)
		case y = <-c3:
			fmt.Printf("%v <-f2\n", x)
		case <-ctx.Done():
			wg.Wait()
			break loop2
		}
	}

	fmt.Printf("x:%v\ty:%v\n", x, y)

	// ------------------------------------------------------------
	// どの case も準備できていない場合は、defaultが実行される
	// ------------------------------------------------------------
	var (
		ctx2, cancel2 = context.WithTimeout(context.Background(), 2*time.Second)
	)

	defer cancel2()

loop3:
	for {
		select {
		case <-ctx2.Done():
			fmt.Println("終わり")
			break loop3
		default:
			fmt.Println("まだ、どのcaseも準備できていない")
			time.Sleep(500 * time.Millisecond)
		}
	}

	return nil
}
