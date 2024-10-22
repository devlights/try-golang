package main

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type (
	counterHandler struct {
		slog.Handler
		count *int
		mu    *sync.Mutex
	}
)

func newHandler(handler slog.Handler) *counterHandler {
	var (
		count = 0
		mu    sync.Mutex
	)
	return &counterHandler{Handler: handler, count: &count, mu: &mu}
}

func (me *counterHandler) Handle(ctx context.Context, r slog.Record) error {
	me.mu.Lock()
	defer me.mu.Unlock()

	r.AddAttrs(slog.String("millis", time.Now().Format(".000")))
	r.AddAttrs(slog.Int("count", *me.count))
	*me.count++

	return me.Handler.Handle(ctx, r)
}

func (me *counterHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &counterHandler{Handler: me.Handler.WithAttrs(attrs), count: me.count, mu: me.mu}
}

func (me *counterHandler) WithGroup(name string) slog.Handler {
	return &counterHandler{Handler: me.Handler.WithGroup(name), count: me.count, mu: me.mu}
}

func main() {
	var (
		level = &slog.LevelVar{}
		opt   = &slog.HandlerOptions{
			Level:       level,
			ReplaceAttr: replaceAttr,
		}
		writer  = os.Stdout
		handler = newHandler(slog.NewTextHandler(writer, opt))
		logger  = slog.New(handler)

		wg             sync.WaitGroup
		goroutineCount = runtime.NumCPU() * 2
		loopCount      = 2
	)

	logger.Info("Start", "NumCPU", runtime.NumCPU(), "goroutineCount", goroutineCount)

	wg.Add(goroutineCount)
	for i := range goroutineCount {
		go func(logger *slog.Logger) {
			defer wg.Done()
			for i := range loopCount {
				logger.Info(strconv.Itoa(i))
			}
		}(logger.With("goroutine", i))
	}

	wg.Wait()
}

func replaceAttr(g []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey {
		return slog.Attr{}
	}

	return a
}
