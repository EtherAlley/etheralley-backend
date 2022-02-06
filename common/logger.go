package common

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

type ILogger interface {
	Info(ctx context.Context, msg string)
	Infof(ctx context.Context, msg string, v ...interface{})
	Error(ctx context.Context, msg string)
	Errorf(ctx context.Context, msg string, v ...interface{})
	Err(ctx context.Context, err error, msg string)
	Errf(ctx context.Context, err error, msg string, v ...interface{})
	Debug(ctx context.Context, msg string)
	Debugf(ctx context.Context, msg string, v ...interface{})
	Event(ctx context.Context, strs []struct {
		Key   string
		Value string
	})
}

type logger struct {
	settings ISettings
	logger   zerolog.Logger
}

func NewLogger(settings ISettings) ILogger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var zLogger zerolog.Logger

	if settings.IsDev() {
		zLogger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
		zLogger = zLogger.Level(zerolog.DebugLevel)
	} else {
		zLogger = zerolog.New(os.Stderr)
		zLogger = zLogger.Level(zerolog.InfoLevel)
	}

	return &logger{
		settings: settings,
		logger:   zLogger.With().Timestamp().Logger(),
	}
}

// add additional context to the event log
func (l *logger) addContext(ctx context.Context, event *zerolog.Event) {
	requestId := ctx.Value(ContextKeyRequestId)
	if requestId != nil {
		event.Str("requestid", requestId.(string))
	}
}

func (l *logger) Info(ctx context.Context, msg string) {
	event := l.logger.Info()
	l.addContext(ctx, event)
	event.Msg(msg)
}

func (l *logger) Infof(ctx context.Context, msg string, v ...interface{}) {
	event := l.logger.Info()
	l.addContext(ctx, event)
	event.Msgf(msg, v...)
}

func (l *logger) Error(ctx context.Context, msg string) {
	event := l.logger.Error()
	l.addContext(ctx, event)
	event.Msg(msg)
}

func (l *logger) Errorf(ctx context.Context, msg string, v ...interface{}) {
	event := l.logger.Error()
	l.addContext(ctx, event)
	event.Msgf(msg, v...)
}

func (l *logger) Err(ctx context.Context, err error, msg string) {
	event := l.logger.Error()
	l.addContext(ctx, event)
	event.Stack().Err(err).Msg(msg)
}

func (l *logger) Errf(ctx context.Context, err error, msg string, v ...interface{}) {
	event := l.logger.Error()
	l.addContext(ctx, event)
	event.Stack().Err(err).Msgf(msg, v...)
}

func (l *logger) Debug(ctx context.Context, msg string) {
	event := l.logger.Debug()
	l.addContext(ctx, event)
	event.Msg(msg)
}

func (l *logger) Debugf(ctx context.Context, msg string, v ...interface{}) {
	event := l.logger.Debug()
	l.addContext(ctx, event)
	event.Msgf(msg, v...)
}

func (l *logger) Event(ctx context.Context, strs []struct {
	Key   string
	Value string
}) {
	event := l.logger.Info()
	l.addContext(ctx, event)
	for _, str := range strs {
		event.Str(str.Key, str.Value)
	}
	event.Msg("event")
}
