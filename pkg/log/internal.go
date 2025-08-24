package log

import (
	"context"

	"github.com/rs/zerolog"
)

func getLogger(ctx context.Context) zerolog.Logger {
	if logger := ctx.Value(loggerCtxKeyType{}); logger != nil {
		return logger.(zerolog.Logger)
	}

	return defaultLogger
}
