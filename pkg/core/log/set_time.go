package log

import (
	"context"
	"time"

	"github.com/rs/zerolog"
)

func SetDuration(ctx context.Context, key string, val time.Duration) context.Context {
	logger := getLogger(ctx)
	logger.UpdateContext(func(logCtx zerolog.Context) zerolog.Context {
		logCtx = logCtx.Dur(key, val)
		return logCtx
	})
	return context.WithValue(ctx, loggerCtxKeyType{}, logger)
}

func SetDurations(ctx context.Context, key string, vals []time.Duration) context.Context {
	logger := getLogger(ctx)
	logger.UpdateContext(func(logCtx zerolog.Context) zerolog.Context {
		logCtx = logCtx.Durs(key, vals)
		return logCtx
	})
	return context.WithValue(ctx, loggerCtxKeyType{}, logger)
}

func SetTime(ctx context.Context, key string, val time.Time) context.Context {
	logger := getLogger(ctx)
	logger.UpdateContext(func(logCtx zerolog.Context) zerolog.Context {
		logCtx = logCtx.Time(key, val)
		return logCtx
	})
	return context.WithValue(ctx, loggerCtxKeyType{}, logger)
}

func SetTimes(ctx context.Context, key string, vals []time.Time) context.Context {
	logger := getLogger(ctx)
	logger.UpdateContext(func(logCtx zerolog.Context) zerolog.Context {
		logCtx = logCtx.Times(key, vals)
		return logCtx
	})
	return context.WithValue(ctx, loggerCtxKeyType{}, logger)
}
