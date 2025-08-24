package log

import (
	"context"
)

func SetStrVal(ctx context.Context, key, value string) context.Context {
	loggerCtx := getLogger(ctx).With()
	loggerCtx = loggerCtx.Str(key, value)

	return context.WithValue(ctx, loggerCtxKeyType{}, loggerCtx.Logger())
}
