package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"os"
	"time"
)

const (
	PasswordKey = "password"
)

func main() {
	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithTimeoutCause(rootCtx, 1*time.Second, errors.New("too slow"))
	)
	defer cxl()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(_ context.Context) error {
	var (
		opt = &slog.HandlerOptions{
			//
			// ReplaceAttr フィールドにカスタム関数を設定することで
			// 属性が出力される際の挙動をフックすることが出来る.
			//
			// 以下のシグネチャを持つ関数をセットする
			//   func([]string, slog.Attr) slog.Attr
			//
			// 例えば、パスワードなどの秘匿情報をログ出力してしまわないように
			// ここでマスキングしてしまうなど。
			//
			ReplaceAttr: replaceAttr,
		}
		writer     = os.Stdout
		handler    = slog.NewJSONHandler(writer, opt)
		rootLogger = slog.New(handler)
		logger     = rootLogger.With()
	)

	logger.Info("passwordフィールドはマスクされる", slog.String(PasswordKey, "12345"))

	return nil
}

func replaceAttr(group []string, a slog.Attr) slog.Attr {
	switch a.Key {
	case slog.TimeKey:
		a = slog.Attr{}
	case PasswordKey:
		a.Value = slog.StringValue("********")
	}

	return a
}
