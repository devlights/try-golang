package times

import (
	"context"
	"time"

	"github.com/devlights/gomy/output"
)

// AfterFunc は、time.AfterFunc のサンプルです。
func AfterFunc() error {

	var (
		ctx, cxl = context.WithTimeout(context.Background(), 3*time.Second)
		timer    = time.AfterFunc(2*time.Second, func() {
			output.Stdoutl("time.AfterFunc", "timed out")
		})
	)

	defer cxl()

	// time.AfterFunc の戻り値で取得する *Timer は永遠に完了状態になることはない。
	// なので、<-timer.C とかすると永遠に待つことになる。
	// この *Timer は、途中でキャンセルしたい場合に利用するためにある。
LOOP:
	for {
		select {
		case <-ctx.Done():
			timer.Stop()
			break LOOP
		default:
			output.Stdoutl("loop", ".")
			time.Sleep(500 * time.Millisecond)
		}
	}

	return nil
}
