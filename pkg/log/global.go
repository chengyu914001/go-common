package log

import (
	"context"

	"github.com/rs/zerolog"
)

func SetStrKetVals(ctx context.Context, keyVals ...[2]string) context.Context {
	logger := getLogger(ctx)
	logger.UpdateContext(func(logCtx zerolog.Context) zerolog.Context {
		for _, keyVal := range keyVals {
			logCtx = logCtx.Str(keyVal[0], keyVal[1])
		}
		return logCtx
	})
	return context.WithValue(ctx, loggerCtxKeyType{}, logger)
}

func Info(ctx context.Context, msg string) {
	logger := getLogger(ctx)
	logger.Info().Msg(msg)
}

func Warn(ctx context.Context, msg string) {
	logger := getLogger(ctx)
	logger.Warn().Msg(msg)
}

func Error(ctx context.Context, msg string) {
	logger := getLogger(ctx)
	logger.Error().Caller(1).Msg(msg)
}

type loggerCtxKeyType struct{}

func getLogger(ctx context.Context) zerolog.Logger {
	if logger := ctx.Value(loggerCtxKeyType{}); logger != nil {
		return logger.(zerolog.Logger)
	}

	return globalLogger
}
