package common

import (
	"os"

	"github.com/rs/zerolog"
)

type ILogger interface {
	Info(msg string)
	Infof(msg string, v ...interface{})
	Error(msg string)
	Errorf(msg string, v ...interface{})
	Err(err error, msg string)
	Errf(err error, msg string, v ...interface{})
	Debug(msg string)
	Debugf(msg string, v ...interface{})
}

type logger struct {
	logger zerolog.Logger
}

func NewLogger(settings ISettings) ILogger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var zLogger zerolog.Logger

	if settings.IsDev() {
		zLogger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		zLogger = zerolog.New(os.Stderr)

	}

	return &logger{
		logger: zLogger.With().Timestamp().Logger(),
	}
}

func (l logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l logger) Infof(msg string, v ...interface{}) {
	l.logger.Info().Msgf(msg, v...)
}

func (l logger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

func (l logger) Errorf(msg string, v ...interface{}) {
	l.logger.Error().Msgf(msg, v...)
}

func (l logger) Err(err error, msg string) {
	l.logger.Error().Err(err).Msg(msg)
}

func (l logger) Errf(err error, msg string, v ...interface{}) {
	l.logger.Error().Err(err).Msgf(msg, v...)
}

func (l logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l logger) Debugf(msg string, v ...interface{}) {
	l.logger.Debug().Msgf(msg, v...)
}
