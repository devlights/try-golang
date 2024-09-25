package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	//
	// slogのデフォルト設定
	//
	// - JSON形式
	// - DEBUGレベルから出力
	// - ソースの情報は出力しない
	// - time属性は表示しない
	//
	var (
		opt = &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: false,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey {
					return slog.Attr{}
				}
				return a
			},
		}
		handler = slog.NewJSONHandler(os.Stdout, opt)
		logger  = slog.New(handler)
	)
	slog.SetDefault(logger)

	type (
		ctxKey string
	)
	var (
		rootCtx = context.Background()
		ctx     = context.WithValue(rootCtx, ctxKey("ctx-key"), "ctx-value")
	)
	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	//
	// 従来のlogパッケージと異なり、log/slogパッケージには
	// ログレベルが存在する. 存在するレベルは
	//
	// - Debug
	// - Info
	// - Warn
	// - Error
	//
	// また、各ログレベル毎のメソッドは
	// XXXContextというメソッドも持っている。
	//
	// これにより、コンテキストに含まれる情報をログに出力することが可能となっている。
	// しかし、コンテキストからの情報を扱うにはカスタムハンドラが必要となる。
	//
	// メッセージの後のキー/値は、いくつでも設定出来る (...any となっている)
	//
	const (
		msg = "hello world"
		k   = "key"
		v   = "value"
	)

	slog.Debug(msg, k, v)
	slog.DebugContext(ctx, msg, k, v)

	slog.Info(msg, k, v)
	slog.Warn(msg, k, v)
	slog.Error(msg, k, v)

	//
	// ログレベルを外から指定することができる汎用メソッドもある
	//
	slog.Log(ctx, slog.LevelInfo, msg, k, v)

	return nil
}
