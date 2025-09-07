package log

import (
	"github.com/chengyu914001/go-common/pkg/core/sysconst"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"os"
	"runtime"
	"strconv"
	"time"
)

func init() {
	setAllLoggerTime()
	initGlobalLogger()
	setGlobalLoggerLevel()
}

func setAllLoggerTime() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}
}

var globalLogger zerolog.Logger

func initGlobalLogger() {
	writer := diode.NewWriter(
		os.Stderr,
		runtime.NumCPU()<<12,
		time.Millisecond,
		func(missed int) {
			os.Stderr.WriteString("Drop " + strconv.Itoa(missed) + " logs\n")
		},
	)
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
