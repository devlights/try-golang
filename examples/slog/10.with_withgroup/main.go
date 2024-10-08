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
		ctx, cxl = context.WithTimeout(rootCtx, 2*time.Second)
	)
	defer cxl()

	if err := run(ctx); err != nil {
		slog.Error("FAILED", "err", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	var (
		level = &slog.LevelVar{}
		opt   = &slog.HandlerOptions{
			Level:       level,
			ReplaceAttr: composite(noTimeKey),
		}
		writer     = os.Stdout
		handler    = slog.NewJSONHandler(writer, opt)
		rootLogger = slog.New(handler)
		logger     = rootLogger.With()
	)

	// Logger.With()メソッドは、キーと値のペアを受け取り、それらの属性を持つ新しいロガーを返す.
	// 新たにWithメソッドを使って作成したロガーには、デフォルトでのこの属性が付与されて出力される。
	logger = logger.With("action", "login", slog.Time("timestamp", dt()))
	<-proc1(ctx, logger).Done()

	// Logger.WithGroup()メソッドは、Logger.With()メソッドのグループ版のようなもの
	logger = logger.WithGroup("user")
	<-proc2(ctx, logger).Done()

	return nil
}

func proc1(ctx context.Context, logger *slog.Logger) context.Context {
	ctx, cxl := context.WithCancelCause(ctx)
	go func() {
		defer cxl(nil)
		<-time.After(150 * time.Millisecond)
		logger.Info("LOGIN")
	}()

	return ctx
}

func proc2(ctx context.Context, logger *slog.Logger) context.Context {
	ctx, cxl := context.WithCancelCause(ctx)
	go func() {
		defer cxl(nil)
		logger.Info("DETAIL", slog.Int("id", 999), slog.String("name", "user-1"), slog.Bool("auth", true))
		<-time.After(300 * time.Millisecond)
	}()

	return ctx
}

func dt() time.Time {
	return time.Date(2024, time.December, 31, 0, 0, 0, 0, time.UTC)
}

func noTimeKey(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey {
		return slog.Attr{}
	}

	return a
}

func composite(fns ...func([]string, slog.Attr) slog.Attr) func([]string, slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		var lastReturn slog.Attr
		for _, fn := range fns {
			lastReturn = fn(groups, a)
		}

		return lastReturn
	}
}
