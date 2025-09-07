package log

import (
	"context"

	"github.com/rs/zerolog"
)

func SetStr(ctx context.Context, key string, val string) context.Context {
	logger := getLogger(ctx)
	logger.UpdateContext(func(logCtx zerolog.Context) zerolog.Context {
		logCtx = logCtx.Str(key, val)
		return logCtx
	})
	return context.WithValue(ctx, loggerCtxKeyType{}, logger)
}

func SetStrs(ctx context.Context, key string, vals []string) context.Context {
	logger := getLogger(ctx)
	logger.UpdateContext(func(logCtx zerolog.Context) zerolog.Context {
		logCtx = logCtx.Strs(key, vals)
		return logCtx
	})
	return context.WithValue(ctx, loggerCtxKeyType{}, logger)
}
