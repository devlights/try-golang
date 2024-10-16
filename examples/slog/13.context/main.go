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
	MainTimeout = time.Second
	ProcTimeout = 50 * time.Millisecond
)

var (
	ErrMainTooSlow = errors.New("(MAIN) TOO SLOW")
	ErrProcTooSlow = errors.New("(PROC) TOO SLOW")
)

func init() {
	log.SetFlags(0)
}

func main() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithTimeoutCause(rootCtx, time.Second, ErrMainTooSlow)
		procCtx          = run(mainCtx)
		err              error
	)
	defer mainCxl()

	select {
	case <-mainCtx.Done():
		err = context.Cause(mainCtx)
	case <-procCtx.Done():
		if err = context.Cause(procCtx); errors.Is(err, context.Canceled) {
			err = nil
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}

func run(pCtx context.Context) context.Context {
	var (
		ctx, cxl = context.WithCancelCause(pCtx)
	)

	go func() {
		cxl(proc(ctx))
	}()
	go func() {
		<-time.After(ProcTimeout)
		cxl(ErrProcTooSlow)
	}()

	return ctx
}

func proc(pCtx context.Context) error {
	//
	// slog には、context.Context を受け取るログ出力メソッドも存在する。
	//
	// - DebugContext
	// - InfoContext
	// - WarnContext
	// - ErrorContext
	// - Log
	//
	// 上記のメソッドには、context.Context を指定できるが
	// 指定しただけだと context.Context の中に設定されている
	// キー/値 は、何も出力されない。
	//
	// 自前で設定した context.Context のキー/値 をログに出力するには
	// カスタムハンドラを作成する必要がある。
	//
	// カスタムハンドラを作成する場合の注意点として
	// Handleメソッドのみをオーバーライドして利用しているサンプルが
	// 結構あるが、これだと slog.Logger.With() や slog.Logger.WithGroup() などで
	// 情報を追加したロガーを利用する場合に、元となるハンドラが利用されてしまい
	// context.Contextの情報が出力されなくなってしまうことに注意が必要。
	// slog.Handler.WithAttrs(), slog.Handler.WithGroup() もオーバーライドしておく。
	//

	var (
		ctx = setCtxValue(pCtx, "helloworld")

		level = slog.LevelInfo
		opt   = &slog.HandlerOptions{
			Level:       level,
			ReplaceAttr: replaceAttr,
		}
		writer     = os.Stdout
		handler    = NewContextHandler(slog.NewJSONHandler(writer, opt))
		rootLogger = slog.New(handler)
		logger     = rootLogger.With("id", 1)
	)

	logger.DebugContext(ctx, "Call DebugContext")
	logger.InfoContext(ctx, "Call InfoContext")
	logger.WarnContext(ctx, "Call WarnContext")
	logger.ErrorContext(ctx, "Call ErrorContext")

	return nil
}

func replaceAttr(group []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey {
		return slog.Attr{}
	}

	return a
}
