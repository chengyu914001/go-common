package log

import (
	"context"
	// "fmt"
	"github.com/chengyu914001/go-common/pkg/common"
	"github.com/rs/zerolog"
	"os"
)

var (
	defaultLogger zerolog.Logger
	logCtxKey     struct{}
)

func init() {
	zerolog.CallerSkipFrameCount = 3

	switch common.GetEnvMode() {
	case common.ENV_MODE_LOCAL:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case common.ENV_MODE_DEV:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case common.ENV_MODE_PROD:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}

	defaultLogger = zerolog.New(os.Stdout).With().Timestamp().Str("service", common.GetServiceName()).Logger()
}

func getLogger(ctx context.Context) zerolog.Logger {
	logger := defaultLogger
	if ctx.Value(logCtxKey) != nil {
		logger = ctx.Value(logCtxKey).(zerolog.Logger)
	}
	return logger
}

func SetTraceID(ctx context.Context, traceID string) context.Context {
	logger := getLogger(ctx).With().Str("trace_id", traceID).Logger()
	return context.WithValue(ctx, logCtxKey, logger)
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
	logger.Error().Caller().Msg(msg)
}
