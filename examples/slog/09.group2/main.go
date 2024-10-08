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
		slog.Error("Error", "err", err)
		os.Exit(1)
	}
}

func run(_ context.Context) error {
	//
	// slogパッケージにおけるグループの概念は、関連する属性をまとめて階層構造を作成するための強力な機能である。
	// グループを使用することで、ログの可読性が向上し、構造化されたデータの管理が容易になる。
	//
	// グループの利点は以下の通り
	//   - 構造化: 関連する情報を論理的にグループ化出来る
	//   - 可読性: ログの階層構造が明確になり、情報の関連性が理解しやすくなる
	//   - フィルタリング: グループ名を使用して特定の情報を容易に抽出できる
	//
	// 特に JSON 形式で構造化ログとして出力した場合に扱いやすい。
	//
	// 基本的に利用するのは
	//   - slog.With()
	//   - slog.WithGroup()
	//   - slog.Group() と {slog.Int(), slog.String(), slog.Time() など}
	// となる。
	//
	// グループを付与すると、HandlerOptions.ReplaceAttr()の第一引数 groups に
	// そのキーが所属するグループ情報が設定される。
	//
	var (
		level = &slog.LevelVar{}
		opt   = &slog.HandlerOptions{
			Level:       level,
			ReplaceAttr: composite(noTimeKey),
		}
		handler    = slog.NewJSONHandler(os.Stdout, opt)
		rootLogger = slog.New(handler)
		logger     = rootLogger.With()
	)

	// ログ出力時に、slog.Group()関数を使って属性をグループ化
	logger.Info("Operation",
		slog.Group("group-1",
			slog.Int("i", 999),
			slog.Bool("b", false),
			slog.Time("t", dt()),
			slog.String("s", "hello world"),
		),
		slog.Group("action",
			slog.String("type", "login"),
			slog.Time("timestamp", dt()),
		),
	)

	return nil
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
