package main

import (
	"context"
	"errors"
	"io"
	"log"
	"log/slog"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	MainTimeout = time.Second
	ProcTimeout = 500 * time.Millisecond
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

func proc(_ context.Context) error {
	//
	// slogでは、出力先をハンドラにて指定できる。
	// 出力先には、 io.Writer を指定出来るので好きに調整出来る。
	//
	// ここに gopkg.in/natefinch/lumberjack.v2 などを
	// 設定することにより、ローリングも可能になる。
	//
	// 本サンプルでは、 os.Stdout と lumberjack.Logger を
	// io.MultiWriter で包んで slog.Handler に指定している。
	//

	var (
		level = &slog.LevelVar{}
		opt   = &slog.HandlerOptions{
			Level:       level,
			ReplaceAttr: replaceAttr,
		}
		writer, closeFn = newWriter()
		handler         = slog.NewJSONHandler(writer, opt)
		rootLogger      = slog.New(handler)
		logger          *slog.Logger
	)
	defer closeFn()

	logger = rootLogger.With("loop-count", 10)
	for i := range 10 {
		// 奇数番目のときだけログに出るように小細工
		level.Set(slog.LevelDebug)
		if i%2 == 0 {
			level.Set(slog.LevelInfo)
		}

		logger.Debug("helloworld", "i", i)
	}

	return nil
}

func newWriter() (io.Writer, func()) {
	var (
		writer1 = os.Stdout
		writer2 = &lumberjack.Logger{
			Filename:   "/tmp/try-golang/slog-example/app.log",
			MaxSize:    1,
			MaxBackups: 3,
			MaxAge:     7,
			Compress:   false,
		}

		writer = io.MultiWriter(writer1, writer2)
	)

	return writer, func() { writer2.Close() }
}

func replaceAttr(g []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey {
		return slog.Attr{}
	}

	return a
}
