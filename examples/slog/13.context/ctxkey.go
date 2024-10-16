package main

import "context"

type ctxKey struct{}

var (
	msgKey = ctxKey{}
)

func ctxValue(ctx context.Context) string {
	return ctx.Value(msgKey).(string)
}

func setCtxValue(ctx context.Context, message string) context.Context {
	return context.WithValue(ctx, msgKey, message)
}
