package main

import (
	"context"
	"log/slog"
)

// ContextHandler は、[context.Context]のキーをログの出力に加えるカスタムハンドラです。
type ContextHandler struct {
	slog.Handler
}

// NewContextHandler は、指定されたハンドラを元に [ContextHandler] を作成します。
func NewContextHandler(handler slog.Handler) *ContextHandler {
	return &ContextHandler{handler}
}

// Handle implements [slog.Handler.Handle].
func (me *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if message := ctxValue(ctx); message != "" {
		r.AddAttrs(slog.String("ctxkey", message))
	}

	return me.Handler.Handle(ctx, r)
}

// WithAttrs implements [slog.Handler.WithAttrs].
func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextHandler{h.Handler.WithAttrs(attrs)}
}

// WithGroup implements [slog.Handler.WithGroup].
func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return &ContextHandler{h.Handler.WithGroup(name)}
}
