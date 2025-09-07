package log

import (
	"context"

	"github.com/rs/zerolog"
)

func SetFloat(ctx context.Context, key string, val float64) context.Context {
	logger := getLogger(ctx)
	logger.UpdateContext(func(logCtx zerolog.Context) zerolog.Context {
		logCtx = logCtx.Float64(key, val)
		return logCtx
	})
	return context.WithValue(ctx, loggerCtxKeyType{}, logger)
}

func SetFloats(ctx context.Context, key string, vals []float64) context.Context {
	logger := getLogger(ctx)
	logger.UpdateContext(func(logCtx zerolog.Context) zerolog.Context {
		logCtx = logCtx.Floats64(key, vals)
		return logCtx
	})
	return context.WithValue(ctx, loggerCtxKeyType{}, logger)
}

func SetInt(ctx context.Context, key string, val int64) context.Context {
	logger := getLogger(ctx)
	logger.UpdateContext(func(logCtx zerolog.Context) zerolog.Context {
		logCtx = logCtx.Int64(key, val)
		return logCtx
	})
	return context.WithValue(ctx, loggerCtxKeyType{}, logger)
}

func SetInts(ctx context.Context, key string, vals []int64) context.Context {
	logger := getLogger(ctx)
	logger.UpdateContext(func(logCtx zerolog.Context) zerolog.Context {
		logCtx = logCtx.Ints64(key, vals)
		return logCtx
	})
	return context.WithValue(ctx, loggerCtxKeyType{}, logger)
}
