package log

import (
	"context"

	"os"

	"github.com/chengyu914001/go-common/pkg/sysconst"
	"github.com/rs/zerolog"
)

var defaultLogger zerolog.Logger

type loggerCtxKeyType struct{}

func init() {
	switch sysconst.GetEnvMode() {
	case sysconst.ENV_MODE_LOCAL:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case sysconst.ENV_MODE_DEV:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case sysconst.ENV_MODE_PROD:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}

	defaultLogger = zerolog.New(os.Stdout).With().Timestamp().Str("service", sysconst.GetServiceName()).Logger()
}

func SetStrVal(ctx context.Context, key, value string) context.Context {
	loggerCtx := getLogger(ctx).With()
	loggerCtx = loggerCtx.Str(key, value)

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
