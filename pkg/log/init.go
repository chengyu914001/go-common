package log

import (
	"github.com/chengyu914001/go-common/pkg/sysconst"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"os"
	"time"
)

func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	initGlobalLogger()
	setGlobalLoggerLevel()
}

var globalLogger zerolog.Logger

func initGlobalLogger() {
	writer := diode.NewWriter(os.Stdout, 1<<16, 10*time.Millisecond, nil)
	globalLogger = zerolog.New(writer).
		With().
		Timestamp().
		Str("service", sysconst.GetServiceName()).
		Logger()
}

func setGlobalLoggerLevel() {
	if sysconst.GetEnvMode() == sysconst.ENV_MODE_PROD {
		globalLogger = globalLogger.Level(zerolog.WarnLevel)
	} else {
		globalLogger = globalLogger.Level(zerolog.InfoLevel)
	}
}
