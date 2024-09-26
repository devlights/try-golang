package main

import (
	"context"
	"log"
	"log/slog"
	"os"
)

type (
	SlogWriter struct {
		logger *slog.Logger
		level  slog.Level
	}
)

func (me *SlogWriter) Write(p []byte) (int, error) {
	me.logger.Log(context.Background(), me.level, string(p))
	return len(p), nil
}

func main() {
	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithCancel(rootCtx)
	)
	defer cxl()

	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(_ context.Context) error {
	//
	// slogパッケージは、従来のlogパッケージと共存できるように設計されている。
	// これにより、既存のコードベースを段階的に移行したり、サードパーティライブラリとの互換性を維持したりすることが可能。
	//
	// リンクロガーを利用する場合は、slog.NewLogLogger() を用いる。
	// この *log.Logger 経由でログ出力すると、内部でslogを通じて出力されることになる。
	//

	var (
		opt        = &slog.HandlerOptions{Level: slog.LevelDebug}
		handler    = slog.NewTextHandler(os.Stdout, opt)
		logger     = slog.New(handler)
		linkLogger = slog.NewLogLogger(handler, slog.LevelInfo)
	)

	logger.Debug("[SLOG] HELLO WORLD")
	linkLogger.Println("[LOG ] HELLO WORLD")

	//
	// 上とは別のやり方として、io.Writerを実装した状態にして
	// log.New() で渡して利用するというやり方も出来る。
	//
	// ただし、この場合、出力先として指定しているので
	// 通常の log.Print() の処理と通り、メッセージの末尾に 改行(\n) が追加される。
	//

	var (
		writer      = &SlogWriter{logger, slog.LevelDebug}
		linkLogger2 = log.New(writer, "", 0)
	)

	linkLogger2.Print("[LOG2] HELLO WORLD")

	return nil
}
