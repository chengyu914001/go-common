package log

import (
	"context"
)

func SetStrVal(ctx context.Context, vals ...[2]string) context.Context {
	loggerCtx := getLogger(ctx).With()
	for _, val := range vals {
		loggerCtx = loggerCtx.Str(val[0], val[1])
	}

	return context.WithValue(ctx, loggerCtxKeyType{}, loggerCtx.Logger())
}
