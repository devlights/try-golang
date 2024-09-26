package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
)

type (
	ctxKey int

	Args struct {
		json bool
	}
)

var (
	key1 = ctxKey(1)
	args Args
)

func init() {
	flag.BoolVar(&args.json, "json", false, "output with json-layout")
}

func main() {
	flag.Parse()

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
	if args.json {
		slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	}

	var (
		rootCtx = context.Background()
		ctx     = context.WithValue(rootCtx, key1, "ctx-value-1")
	)

	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	//
	// slog は、「構造化ログ」に対応している。
	//
	// 構造化ログとは、主にJSONなどの機械的に読み取り可能な形でパラメータを付加した形式のログのこと。
	// このようなログは、付加したパラメータなどを動的にパースしてフィルタ・加工することで、
	// 容易にログを絞り込んだり可視化に利用しやすいというメリットがある。
	//
	// REF: https://blog.cybozu.io/entry/2024/08/07/080000
	//

	//
	// slogでは、メッセージの後に キーと値 のペアを使用して追加情報を付与出来る
	// 最初に指定するメッセージは msg というキーが付与される
	//
	slog.Info("メッセージ", "ctx-key", ctx.Value(key1), "key", "value", "user", os.Getenv("USER"), "pwd", os.Getenv("PWD"))

	return nil
}
