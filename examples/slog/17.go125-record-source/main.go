package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

type (
	myHandler struct {
		handler slog.Handler
	}
)

func (me *myHandler) Handle(ctx context.Context, r slog.Record) error {
	var (
		// Go 1.25 から追加された
		// REF: https://pkg.go.dev/log/slog@go1.25.1#Record.Source
		source = r.Source()
	)
	if source == nil {
		return nil
	}

	var (
		f   = filepath.Base(source.File) // ファイルパス (そのままだとフルパス)
		fn  = source.Function            // 関数
		l   = source.Line                // 行
		msg = r.Message
	)
	fmt.Printf("%s:%s:%d\t%s\n", f, fn, l, msg)

	return nil
}

func (me *myHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return me.handler.Enabled(ctx, l)
}

func (me *myHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &myHandler{handler: me.handler.WithAttrs(attrs)}
}

func (me *myHandler) WithGroup(name string) slog.Handler {
	return &myHandler{handler: me.handler.WithGroup(name)}
}

func main() {
	if err := run(); err != nil {
		slog.Error("failed to run", "error", err)
		os.Exit(1)
	}
}

func run() error {
	var (
		v = &slog.LevelVar{}
		o = &slog.HandlerOptions{
			Level:     v,
			AddSource: true,
		}
		w = os.Stderr
		h = &myHandler{handler: slog.NewTextHandler(w, o)}
		l = slog.New(h)
	)
	l.Info("Start")
	{
		time.Sleep(1 * time.Second)
		v.Set(slog.LevelWarn)
	}
	l.Info("End")

	return nil
}
