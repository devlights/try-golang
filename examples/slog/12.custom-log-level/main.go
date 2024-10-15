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
	ProcTimeout = 10 * time.Millisecond
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
		mainCtx, mainCxl = context.WithTimeoutCause(rootCtx, MainTimeout, ErrMainTooSlow)
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

const (
	LevelTrace   = slog.Level(-8)
	LevelVerbose = slog.Level(-2)
	LevelFatal   = slog.Level(12)
)

func proc(ctx context.Context) error {
	//
	// カスタムログレベル
	//
	// slogパッケージでは、デフォルトのログレベル（Debug、Info、Warn、Error）に加えて
	// カスタムログレベルを定義することが出来る。
	//
	// カスタムログレベルの作成は、slog.Level型を使用して行う。
	// 既定のログレベルの値は https://pkg.go.dev/log/slog@go1.23.2#Level を参照。
	//
	// - Debug: -4
	// - Info : 0
	// - Warn : 4
	// - Error: 8
	//
	// カスタムログレベルを作成するには、これらの既存のレベルの間の値を選択する。
	//
	// カスタムログレベルをそのままログ出力すると "DEBUG-4" のような表示となるため
	// 文字列表現を出力したい場合は、適時変換する関数などを用意して処理する。
	//

	var (
		level = slog.LevelDebug
		opt   = &slog.HandlerOptions{
			Level:       level,
			ReplaceAttr: replaceAttr,
		}
		writer     = os.Stdout
		handler    = slog.NewJSONHandler(writer, opt)
		rootLogger = slog.New(handler)
		logger     = rootLogger.With()
	)

	//
	// カスタムログレベルを利用する場合は slog.Log() メソッドで出力する
	//
	logger.Log(ctx, LevelTrace, "TRACE MESSAGE")
	logger.Log(ctx, LevelVerbose, "VERBOSE MESSAGE")
	logger.Log(ctx, slog.LevelDebug, "Debug MESSAGE")
	logger.Log(ctx, slog.LevelInfo, "INFO MESSAGE")
	logger.Log(ctx, LevelFatal, "FATAL MESSAGE")

	return nil
}

func replaceAttr(group []string, a slog.Attr) slog.Attr {
	switch a.Key {
	case slog.TimeKey:
		return slog.Attr{}
	case slog.LevelKey:
		a.Value = slog.StringValue(l2s(a.Value.Any().(slog.Level)))
	}

	return a
}

func l2s(l slog.Level) string {
	switch l {
	case LevelTrace:
		return "TRACE"
	case LevelVerbose:
		return "VERBOSE"
	case LevelFatal:
		return "FATAL"
	default:
		return l.String()
	}
}
