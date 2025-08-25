package log

import (
	"github.com/chengyu914001/go-common/pkg/sysconst"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"os"
	"time"
)

var defaultLogger zerolog.Logger

type loggerCtxKeyType struct{}

func init() {
	switch sysconst.GetEnvMode() {
	case sysconst.ENV_MODE_LOCAL, sysconst.ENV_MODE_DEV:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case sysconst.ENV_MODE_PROD:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	writer := diode.NewWriter(os.Stdout, 32<<20, 10*time.Millisecond, nil)
	defaultLogger = zerolog.New(writer).
		With().
		Timestamp().
		Str("service", sysconst.GetServiceName()).
		Logger()
}
