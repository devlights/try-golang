package goroutines

import (
	"context"
	"time"

	"github.com/devlights/gomy/output"
)

// WithContextTimeout -- context.Contextを用いてタイムアウト付きで待ち合わせを行うサンプルです
func WithContextTimeout() error {
	// 処理内で利用する共通関数
	var (
		iter = func(n int) []struct{} { return make([]struct{}, n) }
		now  = func() int64 { return time.Now().UTC().Unix() }
	)

	// コンテキスト定義
	var (
		rootCtx             = context.Background()
		mainCtx, mainCancel = context.WithTimeout(rootCtx, 2*time.Second)
		procCtx, procCancel = context.WithTimeout(mainCtx, 1*time.Second)
	)

	defer mainCancel()
	defer procCancel()

	// ---------------------------------------------------
	// 以下の仕様とする
	//   - アプリケーション全体の生存期間は２秒間
	//   - 非同期起動するゴルーチン全体の生存期間は１秒間
	// ---------------------------------------------------

	ctx := func(pCtx context.Context) context.Context {
		ctx, cancel := context.WithCancel(pCtx)

		go func() {
			defer cancel()

			for i := range iter(10) {
				select {
				case <-ctx.Done():
					output.Stdoutl("[ctx inside]", "done", now())
					return
				default:
				}

				output.Stdoutl("[goroutine]", i)
				time.Sleep(200 * time.Millisecond)
			}
		}()

		return ctx
	}(procCtx)

	// 待ち合わせしながら経過出力
	var (
		doneCtx, doneProc, doneMain = ctx.Done(), procCtx.Done(), mainCtx.Done()
	)

LOOP:
	for {
		select {
		case <-doneCtx:
			output.Stdoutl("[ctx]", "done", now())
			doneCtx = nil
		case <-doneProc:
			output.Stdoutl("[proc]", "done", now())
			doneProc = nil
		case <-doneMain:
			output.Stdoutl("[main]", "done", now())
			break LOOP
		}
	}

	return nil
}
