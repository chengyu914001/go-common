package log

import (
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
