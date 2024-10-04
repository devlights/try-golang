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
		slog.Error("A fatal error occurred", "err", err)
		os.Exit(1)
	}
}

func run(_ context.Context) error {
	//
	// slog.LevelVarを使用することで、実行時にログレベルを動的に変更できる.
	//
	// 以下、https://pkg.go.dev/log/slog@go1.23.2#hdr-Levels にある記載を抜粋
	//
	// > Setting it to a LevelVar allows the level to be varied dynamically.
	// > A LevelVar holds a Level and is safe to read or write from multiple goroutines.
	//
	// >> LevelVar を設定することで、レベルを動的に変化させることができます。
	// >> LevelVar は Level を保持し、複数のゴルーチンから安全に読み書きできます。
	//

	var (
		level = &slog.LevelVar{}
		opt   = &slog.HandlerOptions{
			Level:       level,
			ReplaceAttr: noTimeKey,
		}
		handler    = slog.NewTextHandler(os.Stdout, opt)
		rootLogger = slog.New(handler)
		logger     *slog.Logger
	)

	// デフォルトのレベルはINFO。なのでDEBUGは出力されない
	logger = rootLogger.With("idx", "1")
	logger.Debug("this is DEBUG level message")
	logger.Info("this is INFO level message")

	// レベルを変更
	//   現実的な使い方だと、コマンドライン引数で -debug を受け取った場合に
	//   DEBUGレベルを有効にするなどの使い方が出来る。
	level.Set(slog.LevelDebug)

	// デフォルトのレベルがDEBUGに変更された
	logger = rootLogger.With("idx", "2")
	logger.Debug("this is DEBUG level message")
	logger.Info("this is INFO level message")

	return nil
}

func noTimeKey(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey {
		return slog.Attr{}
	}

	return a
}
