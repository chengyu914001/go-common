package log

import (
	"context"

	"os"

	"github.com/chengyu914001/go-common/pkg/constant"
	"github.com/rs/zerolog"
)

var defaultLogger zerolog.Logger

type loggerCtxKeyType struct{}

func init() {
	switch constant.GetEnvMode() {
	case constant.ENV_MODE_LOCAL:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case constant.ENV_MODE_DEV:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case constant.ENV_MODE_PROD:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}

	defaultLogger = zerolog.New(os.Stdout).With().Timestamp().Str("service", constant.GetServiceName()).Logger()
}

func SetDymanicValues(
	ctx context.Context,
	values map[string]string,
) context.Context {
	loggerCtx := getLogger(ctx).With()
	for k, v := range values {
		loggerCtx = loggerCtx.Str(k, v)
	}

	return context.WithValue(ctx, loggerCtxKeyType{}, loggerCtx.Logger())
}

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

func getLogger(ctx context.Context) zerolog.Logger {
	if logger := ctx.Value(loggerCtxKeyType{}); logger != nil {
		return logger.(zerolog.Logger)
	}

	return defaultLogger
}
