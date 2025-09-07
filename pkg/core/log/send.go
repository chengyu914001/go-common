package log

import (
	"context"

	"github.com/rs/zerolog"
)

func Debug(ctx context.Context, msg string) {
	logger := getLogger(ctx)
	logger.Debug().Msg(msg)
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
