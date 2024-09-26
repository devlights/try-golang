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
		panic(err)
	}
}

func run(_ context.Context) error {
	//
	// slog.Info()のように明示的にロガーを生成せずに
	// 利用すると、内部でデフォルトのロガーが呼び出される。
	//
	// デフォルトロガーは以下の特徴を持つ。
	//    - 標準エラー出力に出力する
	//    - テキスト形式
	//    - ログレベルはInfo
	//
	// デフォルトロガーは、グローバル変数として実装されているため
	// 並行処理時の競合を避けるために内部で同期化されている。
	// そのため、高負荷な環境では若干のパフォーマンスオーバーヘッドが発生する可能性がある。
	//
	// デフォルトロガーは、アプリケーション全体で一貫したログ出力を簡単に実現できる反面、柔軟性に欠ける面がある。
	// 特定のモジュールや機能で異なるログ設定が必要な場合は、個別のロガーインスタンスを作成する方が良い。
	// (これは他の言語の場合でも同様)
	//
	slog.Debug("これは出力されない", "key", "value")
	slog.Info("これは出力される", "key", "value")

	//
	// デフォルトロガーの設定を変更するには
	//    slog.SetDefault()
	// を利用する
	//
	var (
		opt     = &slog.HandlerOptions{Level: slog.LevelDebug}
		handler = slog.NewJSONHandler(os.Stderr, opt)
		logger  = slog.New(handler)
	)
	slog.SetDefault(logger)

	slog.Debug("今度は出力される", "key", "value")
	slog.Info("これは出力される", "key", "value")

	return nil
}
