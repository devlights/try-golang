package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// TickAndTicker -- time.Tick と time.NewTicker の利用シーンの違いについてのサンプルです。
func TickAndTicker() error {
	// -------------------------------------------------
	// time.Tick と time.NewTicker の使い分け
	//
	// time.Tick は、以下の定義となっている。
	//   func Tick(d time.Duration) <-chan Time
	// 受信専用のチャネルを返しているので、内部で goroutine を
	// 起動してチャネル経由で値を返してきている。
	// 受信専用のチャネルであるので、このチャネルをユーザ側で
	// クローズすることは出来ない。なので、Tickを呼び出した際に
	// 生成される goroutine は止まることが無い。
	// 止まるタイミングがなく、ずっと動いている goroutine は
	// メインゴルーチン以外は goroutine leak していると考える。
	//
	// なので、time.Tick のドキュメントには以下のように記載されている。
	//   While Tick is useful for clients that have no need to shut down the Ticker,
	//   be aware that without a way to shut it down the underlying
	//   Ticker cannot be recovered by the garbage collector; it "leaks".
	//
	// time.Tick で生成される goroutine は終了しないので
	// アプリケーションの生存期間と同じ時間生存できるタイミングで
	// 利用する場合は便利である。
	//
	// それ以外のケース、例えば 特定の時間枠で処理する goroutineの
	// 中で利用したい場合は、time.NewTicker で明示的に time.Ticker を
	// 生成して利用するべき。time.Tickerには Stop メソッドが用意されている
	// ので、それを呼び出すと内部リソースが開放される。
	// (time.Ticker.C のチャネルはクローズされないことに注意)
	// -------------------------------------------------

	// 一時的な処理時間で動作するゴルーチン
	done := func() <-chan struct{} {
		done := make(chan struct{})
		go func() {
			defer close(done)

			// このような場合は time.Tick ではなく time.NewTicker を使うべき
			ticker := time.NewTicker(500 * time.Millisecond)
			timeout := time.After(2 * time.Second)
			defer ticker.Stop()

		LOOP:
			for {
				select {
				case t := <-ticker.C:
					output.Stdoutl("[goroutine] ", t.UTC().Unix())
				case <-timeout:
					break LOOP
				}
			}
		}()

		return done
	}()

	// ここはメインゴルーチン
	// ここで処理が終わるまでインターバルする場合などに利用する場合は
	// time.Tick は便利（アプリがそのまま終了するので goroutine leak は問題にならない)
	var (
		tick    <-chan time.Time
		timeout = time.After(5 * time.Second)
	)

LOOP:
	for {
		select {
		case <-done:
			// 非同期処理が終わったのでメインの出力に切り替え
			// 再びこのチャネルが select で選択されないように nil を設定しておく
			//lint:ignore SA1015 サンプルなのでOK
			tick = time.Tick(500 * time.Millisecond)
			done = nil

			output.Stdoutl("[main     ]", "goroutine end.")
		case t := <-tick:
			output.Stdoutl("[main     ]", t.UTC().Unix())
		case <-timeout:
			break LOOP
		}
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_tick_and_ticker

	   [Name] "time_tick_and_ticker"
	   [goroutine]          1711519193
	   [goroutine]          1711519193
	   [goroutine]          1711519194
	   [goroutine]          1711519194
	   [main     ]          goroutine end.
	   [main     ]          1711519195
	   [main     ]          1711519195
	   [main     ]          1711519196
	   [main     ]          1711519196
	   [main     ]          1711519197


	   [Elapsed] 5.000150923s
	*/

}
