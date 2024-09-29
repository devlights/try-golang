package main

import (
	"context"
	"log/slog"
	"os"
	"time"
)

func main() {
	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithCancel(rootCtx)
	)
	defer cxl()

	if err := run(ctx); err != nil {
		slog.Error("A fatal error occurred", "error", err)
		os.Exit(1)
	}
}

func run(_ context.Context) error {
	//
	// slogパッケージは、テキスト形式でのログ出力を簡単に実現できる。
	// ハンドラに slog.TextHandler を指定すれば良い。
	//
	// 注意点として、TextHandlerとJSONHandlerでは timeキー の出力値が少し異なる。
	//
	// - TextHandler: ベースは time.RFC3339Nano だが、ミリ秒までの精度で出力される
	// - JSONHandler: time.RFC3339Nano
	//
	// Goのソースコード go/src/log/slog/handler.go の appendRFC3339Millis() に
	// 説明が以下のように記載されている。以下、 go1.23.1 時点でのコメント。
	//
	// Format according to time.RFC3339Nano since it is highly optimized,
	// but truncate it to use millisecond resolution.
	// Unfortunately, that format trims trailing 0s, so add 1/10 millisecond
	// to guarantee that there are exactly 4 digits after the period.
	// (time.RFC3339Nanoは高度に最適化されているので、それにしたがってフォーマットするが、
	// ミリ秒の分解能を使うために切り捨てる。残念なことに、このフォーマットでは末尾の0が切り捨てられるので、
	// ピリオドの後に正確に4桁の数字があることを保証するために1/10ミリ秒を追加する。)
	//

	var (
		opt = &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
		handler = slog.NewTextHandler(os.Stdout, opt)
		logger  = slog.New(handler)
		ch      = make(chan int)
	)

	go func() {
		defer close(ch)
		for i := range 3 {
			ch <- i
			<-time.After(10 * time.Millisecond)
		}
	}()

	for i := range ch {
		logger.Debug("LOOP", "i", i)
	}

	return nil
}
