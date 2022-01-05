package logger

import (
	"github.com/rs/zerolog"
	"io"
	"os"
)

type AppLogger struct {
	Log *zerolog.Logger
}

func New(isDebug bool) *AppLogger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &AppLogger{Log: &logger}
}

func NewConsole(isDebug bool) *AppLogger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &AppLogger{Log: &logger}
}
func (l *AppLogger) Output(w io.Writer) zerolog.Logger {
	return l.Log.Output(w)
}
