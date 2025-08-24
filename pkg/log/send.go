package log

import (
	"context"
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
