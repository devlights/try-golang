package main

import (
	"context"
	"log/slog"
	"os"
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
	// slogパッケージは、JSON形式でのログ出力を簡単に実現できる。
	// ハンドラに slog.JSONHandler を指定すれば良い。
	//
	// 注意点として、TextHandlerとJSONHandlerでは timeキー の出力値が少し異なる。
	//
	// - TextHandler: ベースは time.RFC3339Nano だが、ミリ秒までの精度で出力される
	// - JSONHandler: time.RFC3339Nano
	//

	var (
		opt = &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
		handler = slog.NewJSONHandler(os.Stdout, opt)
		logger  = slog.New(handler)
		ch      = make(chan int)
	)

	go func() {
		defer close(ch)
		for i := range 3 {
			ch <- i
		}
	}()

	for i := range ch {
		logger.Debug("LOOP", "i", i)
	}

	return nil
}
