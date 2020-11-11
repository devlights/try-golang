package times

import (
	"fmt"
	"time"
)

// TimeAfter は、 time.After() のサンプルです.
func TimeAfter() error {
	// ------------------------------------------------------------
	// time.After(d Duration) <-chan Time
	//
	// 指定した Duration 後に着火するチャネルを返してくれる.
	// 所定時間が経過したらタイムアウトするような処理を書く際に便利.
	// ------------------------------------------------------------

	// ------------------------------------------------------------
	// 3秒後に受信するチャネルを作って3秒間処理して終了
	// ------------------------------------------------------------
	var (
		timeUp = time.After(3 * time.Second)
	)

loop1:
	for {
		select {
		case <-timeUp:
			fmt.Println("timed up")
			break loop1
		default:
		}

		fmt.Println("processing...")
		time.Sleep(1 * time.Second)
	}

	// ------------------------------------------------------------
	// 値を生産する goroutine を起動し
	// 3秒後に受信するチャネルを作って3秒間処理して終了
	// タイムアップしたら、goroutine 側も停止させる
	//
	// 現実的には、 goroutine を内部で起動する関数の戻り値が chan struct{}と
	// することはあまりなくて、 Result みたいなデータ構造を別途定義して、その中に
	// エラーかどうかなどの情報を詰めて返すチャネルを返すのが多い。
	//   例: <-chan Result
	//
	// type Result struct {
	// 		HasError bool
	// 		Cause error
	// }
	// ------------------------------------------------------------
	var (
		queue   = make(chan int)
		done    = make(chan struct{})
		timeUp2 = time.After(3 * time.Second)
	)

	defer close(queue)

	producer := func(queue chan<- int, done <-chan struct{}) <-chan struct{} {
		var (
			terminated = make(chan struct{})
		)

		go func(done <-chan struct{}) {
			defer close(terminated)

			var (
				i int
			)

		loopGorouine:
			for {
				select {
				case <-done:
					break loopGorouine
				default:
					queue <- i
					i++
					time.Sleep(1 * time.Second)
				}
			}
		}(done)

		return terminated
	}

	terminated := producer(queue, done)

loop2:
	for {
		select {
		case <-timeUp2:
			fmt.Println("timed up")
			close(done)
			break loop2
		case x := <-queue:
			fmt.Printf("[recv] %d\n", x)
		}
	}

	<-terminated
	fmt.Println("goroutine terminated")

	close(queue)
	fmt.Println("close queue")

	return nil
}
