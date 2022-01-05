package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var (
	Log *zerolog.Logger
)

func New(isDebug bool) {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	Log = &logger
}
