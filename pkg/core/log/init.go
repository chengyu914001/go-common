package log

import (
	"github.com/chengyu914001/go-common/pkg/core/sysconst"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"os"
	"runtime"
	"time"
)

func init() {
	setAllLoggerTimeFormat()
	initGlobalLogger()
	setGlobalLoggerLevel()
}

func setAllLoggerTimeFormat() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

var globalLogger zerolog.Logger

func initGlobalLogger() {
	buffer := runtime.NumCPU() << 12
	writer := diode.NewWriter(os.Stdout, buffer, time.Millisecond, func(missed int) {
		if missed > 0 {
			globalLogger.Error().Msgf("Dropped %d logs", missed)
		}
	})
	globalLogger = zerolog.New(writer).
		With().
		Timestamp().
		Str("service", sysconst.GetServiceName()).
		Logger()
}

func setGlobalLoggerLevel() {
	if sysconst.GetEnvMode() == sysconst.ENV_MODE_PROD {
		globalLogger = globalLogger.Level(zerolog.InfoLevel)
	} else {
		globalLogger = globalLogger.Level(zerolog.DebugLevel)
	}
}
