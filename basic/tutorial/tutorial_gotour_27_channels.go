package tutorial

import (
	"fmt"
	"sync"
	"time"
)

func GoTourChannels() error {
	// ------------------------------------------------------------
	// Go言語のチャネル型(Channel)は、チャネルオペレータ(<-)を用いて値の送受信を
	// 行うためのストリーム（通り道）となる。
	//
	// 送受信用、送信用、受信用とそれぞれのチャネルを構築して利用できる.
	//
	// チャネルは、スライスやマップと同じく 組み込み関数 make() を用いて生成する
	//
	// 初期値はnil。 nil なチャネルにアクセスすると ランタイムエラーが発生するので注意
	//
	// チャネルは、　組み込み関数 close() を利用することにより閉じることが可能
	// 閉じたチャネルは、書き込みは出来ないが読み込みは行える
	// チャネルの性質を利用することにより、他の言語のシグナルと同じような動きを実現出来る
	// (専用の型が sync.Cond として用意されているがチャネルを使っても出来る)
	//
	// チャネルは、FIFOのキューのような構造をしている。通常、片方が準備できるまで
	// 送受信はブロックされる。データが送信されると受信できるようになり、その逆もしかり。
	//
	// チャネルは、 chan 型 という形で定義する。
	//   - ch := make(chan int)
	// 上は、intの値をやり取りできる送受信用のチャネルを生成している。
	//
	// 送信用と受信用も以下のようにして定義できる
	//   - sendCh chan<- int := ch
	//   - recvCh <-chan int := ch
	//
	// つまり、矢印の向きにデータが流れる
	//
	// チャネルと利用することにより、明確なロックや条件変数がなくても、goroutineの同期を
	// とることが出来る.
	//
	// チャネルは、forループで処理することも出来て、range チャネル とすることで
	// closeされるまで、データを受信し続けることも出来る.
	// chan + for-loop + select の組合せはとても良く使うパターン。
	//
	// チャネルは、バッファとして利用できる。バッファ付きのチャネルを作るには
	// makeの引数にバッファサイズを付与する。
	//   - make(chan int, 4)
	// バッファが詰まっている場合、送信がブロックされる。
	// バッファが空の場合、受信がブロックされる。
	// (C# の BlockingCollection みたいな感じ)
	//
	// チャネルが close されているかどうかは、以下の構文で判定できる
	// 		if v, ok := <-ch; !ok {
	// 			// closeされている
	// 		}
	// ------------------------------------------------------------
	var (
		ch1 = make(chan interface{})
	)

	run := func(wait time.Duration, prefix string, done chan<- interface{}) {
		if done != nil {
			defer func() {
				done <- struct{}{}
			}()
		}

		fmt.Printf("[%s] func begin\n", prefix)
		time.Sleep(wait)
		fmt.Printf("[%s] func end\n", prefix)
	}

	go run(2*time.Second, "async", ch1)
	run(1*time.Second, "sync", nil)

	// 待ち合わせ
	<-ch1

	// 送受信用のチャネルを作って、それを送信用と受信用にそれぞれ分けて使う
	var (
		calcCh            = make(chan int)
		sendCh chan<- int = calcCh
		recvCh <-chan int = calcCh
		data              = []int{7, 2, 8, -9, -4, 0}
		sum               = func(d []int, c chan<- int) {
			result := 0
			for _, v := range d {
				result += v
			}

			c <- result
		}
	)

	// ２つの goroutine で並行処理
	go sum(data[:len(data)/2], sendCh)
	go sum(data[len(data)/2:], sendCh)

	// 結果を取得
	x, y := <-recvCh, <-recvCh

	fmt.Printf("[sum] x:%v\ty:%v\tsum:%v\n", x, y, x+y)

	// チャネルとforループ
	var (
		loopCh  = make(chan int)
		putData = func(c chan int) {
			defer close(c) // 最後にclose()を呼ぶことでこのチャネルを閉じて、これ以上データが流れることはないことを通知
			for i := 0; i < 5; i++ {
				time.Sleep(1 * time.Second)
				fmt.Printf("[send] %v\n", i)
				c <- i
			}
		}
	)

	// データを作って流す
	go putData(loopCh)

	// 流れてくるデータを出力
	// このループは、対象となる ch が close されることにより break される
	for v := range loopCh {
		fmt.Printf("[recv] %v\n", v)
	}

	// チャネルの close を利用して、待機しているgoroutineを一気に開放する
	// sync.WaitGroupやsync.Cond使っても同じことできるが一応サンプルとして。
	var (
		wg    sync.WaitGroup
		ready = make(chan interface{})
		f     = func(c <-chan interface{}, wg *sync.WaitGroup, prefix string) {
			defer wg.Done()
			fmt.Printf("[%s] wait....\n", prefix)
			<-c
			fmt.Printf("[%s] awake\n", prefix)
		}
	)

	wg.Add(3)
	go f(ready, &wg, "f1")
	go f(ready, &wg, "f2")
	go f(ready, &wg, "f3")

	fmt.Println("close(f) after 2 seconds later")
	time.Sleep(2 * time.Second)
	close(ready)
	wg.Wait()

	// チャネルが close されているかどうかは以下で判定できる
	if _, ok := <-ready; !ok {
		fmt.Println("ch [ready] was closed")
	}

	return nil
}
