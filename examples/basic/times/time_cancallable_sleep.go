package times

import (
	"context"
	"time"

	"github.com/devlights/gomy/output"
)

// CancellableSleep は、キャンセル可能なスリープ処理のサンプルです。
//
// time.Sleep() は、ブロックしてしまうためキャンセル可能な状態で
// スリープ処理を行いたい場合は、time.Tickerを利用して処理する。
//
// # REFERENCES
//   - https://pkg.go.dev/time@go1.22.2#NewTicker
func CancellableSleep() error {
	const (
		timeFormat = time.TimeOnly + ".000"
	)

	var (
		ctx, cxl = context.WithTimeout(context.Background(), 3500*time.Millisecond)
		ticker   = time.NewTimer(5 * time.Second)
	)
	defer cxl()
	defer ticker.Stop()

	output.Stdoutl("[begin]", time.Now().Format(timeFormat))
	defer func() { output.Stdoutl("[end  ]", time.Now().Format(timeFormat)) }()

	//
	// 以下は5秒間スリープする処理があるとして、3.5秒でキャンセルします。
	// time.Sleep(5 * time.Second) とすると 5秒間 ブロックされますが
	// context.Context と *time.Ticker.C を利用してselectで待機するようにしているため
	// 先に完了した方で抜けていきます。
	//
	select {
	case <-ctx.Done():
		output.Stdoutl("[cancel]", ctx.Err())
	case <-ticker.C:
		output.Stdoutl("[sleep]", "done")
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_cancellable_sleep

	   [Name] "time_cancellable_sleep"
	   [begin]              07:33:29.449
	   [cancel]             context deadline exceeded
	   [end  ]              07:33:32.951


	   [Elapsed] 3.502395005s
	*/

}
