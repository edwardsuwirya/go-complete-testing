package logger

import (
	"github.com/rs/zerolog"
	"os"
)

type ErrorSubject string

const (
	ErrorDeliverySubject   ErrorSubject = "DELIVERY"
	ErrorRepositorySubject ErrorSubject = "REPOSITORY"
	CommonSubject          ErrorSubject = "SUBJECT"
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
func F(err error, message string) {
	Log.Fatal().Err(err).Msg(message)
}

func I(message string) {
	Log.Info().Msg(message)
}
func IS(title string, subject []string, message string) {
	if title == "" {
		Log.Info().Strs(string(CommonSubject), subject).Msg(message)
	} else {
		Log.Info().Strs(title, subject).Msg(message)
	}

}
func D(message string) {
	Log.Debug().Msg(message)
}
func ER(err error, message string) {
	Log.Error().Err(err).Msg(message)
}
func ES(err error, title string, subject []string, message string) {
	if title == "" {
		Log.Error().Err(err).Strs(string(CommonSubject), subject).Msg(message)
	} else {
		Log.Error().Err(err).Strs(title, subject).Msg(message)
	}

}
