package log

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

func SetDecimal(ctx context.Context, key string, val decimal.Decimal) context.Context {
	logger := getLogger(ctx)
	logger.UpdateContext(func(logCtx zerolog.Context) zerolog.Context {
		logCtx = logCtx.Str(key, val.String())
		return logCtx
	})
	return context.WithValue(ctx, loggerCtxKeyType{}, logger)
}

func SetDecimals(ctx context.Context, key string, vals []decimal.Decimal) context.Context {
	logger := getLogger(ctx)
	logger.UpdateContext(func(logCtx zerolog.Context) zerolog.Context {
		s := make([]string, len(vals))
		for i, v := range vals {
			s[i] = v.String()
		}
		logCtx = logCtx.Strs(key, s)
		return logCtx
	})
	return context.WithValue(ctx, loggerCtxKeyType{}, logger)
}
