package effectivego26

import (
	"fmt"
	"time"

	"github.com/devlights/gomy/chans"
	"github.com/devlights/gomy/output"
)

// Channels -- Effective Go - Channels の 内容についてのサンプルです。
func Channels() error {
	/*
		https://golang.org/doc/effective_go.html#channels

		チャネルにはバッファ無し (unbuffered) と バッファ有り (buffered) がある
		make(chan xxx) する際に、第２引数に１以上を指定するとバッファ有りとなる。

		バッファ無しのチャネルは主にゴルーチン間の通信に利用される。
		バッファ無しの場合、送受信ともに対向先が操作しないと次の値を処理できないからである。

		バッファ有りのチャネルはセマフォの方に利用することができる。
		例えば、同時に処理する最大数を制限する際などに指定数のバッファでチャネルを作成し
		それをゴルーチン内で、処理の開始時に値をチャネルに送信、処理の終了時に値を受信という風に
		することで、簡易的な導入制限を実施することが出来る。
	*/

	// -------------------------------------------------------------------
	// ゴルーチン間でデータを通信する
	// -------------------------------------------------------------------
	dataCh := make(chan int)

	doneProducer := func() <-chan struct{} {
		done := make(chan struct{})
		go func() {
			defer close(done)
			defer close(dataCh)

			dataCh <- 1
			dataCh <- 2
			dataCh <- 3
		}()

		return done
	}()

	doneConsumer := func() <-chan struct{} {
		done := make(chan struct{})
		go func() {
			defer close(done)
			for v := range dataCh {
				output.Stdoutl("consumer", v)
			}
		}()

		return done
	}()

	<-chans.WhenAll(doneProducer, doneConsumer)

	output.StdoutHr()

	// -------------------------------------------------------------------
	// バッファ無しチャネルは、待機シグナルのようにも利用できる
	// -------------------------------------------------------------------
	sigCh := make(chan struct{})
	go func() {
		output.Stdoutl("goroutine", "作業開始")
		// なにかの非同期処理をしたとする
		<-time.After(1 * time.Second)
		output.Stdoutl("goroutine", "作業終了")

		// 終わったことを通知
		// サンプルなので、ここに書いているがエラーが発生するかもしれないことを
		// 考慮して、本来は defer で書いておくべき
		close(sigCh)
	}()

	// 非同期処理をやっている間に別なことをしているとする
	output.Stdoutl("main", "他の作業開始")
	<-time.After(500 * time.Millisecond)
	output.Stdoutl("main", "他の作業終わり. 非同期処理完了を待つ")

	// 非同期処理の完了を待つ
	<-sigCh
	output.Stdoutl("main", "非同期処理が完了")

	output.StdoutHr()
	<-time.After(100 * time.Millisecond)

	// -------------------------------------------------------------------
	// バッファ有りチャネルで、一度に動く非同期処理の数を制限する
	// -------------------------------------------------------------------
	semCh := make(chan bool, 2) // 一度に動いて良い非同期処理を制限するためのチャネル
	dataCh2 := make(chan int)

	// データを投入
	go func() {
		defer close(dataCh2)

		for i := 0; i < 9; i++ {
			output.Stderrf("producer", "put: %d\n", i)
			dataCh2 <- i
		}
	}()

	// データを処理
	channels := make([]<-chan struct{}, 0)
	for i := 0; i < 9; i++ {
		name := fmt.Sprintf("goroutine-%02d", i)

		c := func(name string) <-chan struct{} {
			done := make(chan struct{})

			semCh <- true
			go func(name string) {
				defer close(done)
				output.Stderrf(name, "%v\n", <-dataCh2)

				<-semCh
			}(name)

			return done
		}(name)

		channels = append(channels, c)
	}

	<-chans.WhenAll(channels...)
	close(semCh)

	return nil
}
