package common

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

func NewLogger(settings *Settings) *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var logger zerolog.Logger

	if settings.IsDev() {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		logger = zerolog.New(os.Stderr)

	}

	return &Logger{
		logger: logger.With().Timestamp().Logger(),
	}
}

func (logger Logger) Info(msg string) {
	logger.logger.Info().Msg(msg)
}

func (logger Logger) Infof(msg string, v ...interface{}) {
	logger.logger.Info().Msgf(msg, v...)
}

func (logger Logger) Error(msg string) {
	logger.logger.Error().Msg(msg)
}

func (logger Logger) Errorf(msg string, v ...interface{}) {
	logger.logger.Error().Msgf(msg, v...)
}

func (logger Logger) Err(err error, msg string) {
	logger.logger.Error().Err(err).Msg(msg)
}

func (logger Logger) Errf(err error, msg string, v ...interface{}) {
	logger.logger.Error().Err(err).Msgf(msg, v...)
}

func (logger Logger) Debug(msg string) {
	logger.logger.Debug().Msg(msg)
}

func (logger Logger) Debugf(msg string, v ...interface{}) {
	logger.logger.Debug().Msgf(msg, v...)
}
