package common

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

type ILogger interface {
	Debug(ctx context.Context) ILogEvent
	Info(ctx context.Context) ILogEvent
	Warn(ctx context.Context) ILogEvent
	Error(ctx context.Context) ILogEvent
}

type ILogEvent interface {
	Send()
	Msg(msg string)
	Msgf(msg string, v ...any)
	Err(err error) ILogEvent
	Strs(strs []struct {
		Key   string
		Value string
	}) ILogEvent
}

type logger struct {
	settings ISettings
	logger   zerolog.Logger
}

func NewLogger(settings ISettings) ILogger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	var zLogger zerolog.Logger

	if settings.IsDev() {
		zLogger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)
	} else {
		zLogger = zerolog.New(os.Stderr).Level(zerolog.InfoLevel)
	}

	return &logger{
		settings: settings,
		logger:   zLogger.With().Timestamp().Logger(),
	}
}

func (l *logger) Debug(ctx context.Context) ILogEvent {
	return l.newEvent(ctx, l.logger.Debug())
}

func (l *logger) Info(ctx context.Context) ILogEvent {
	return l.newEvent(ctx, l.logger.Info())
}

func (l *logger) Warn(ctx context.Context) ILogEvent {
	return l.newEvent(ctx, l.logger.Warn())
}

func (l *logger) Error(ctx context.Context) ILogEvent {
	return l.newEvent(ctx, l.logger.Error())
}

type logEvent struct {
	event *zerolog.Event
}

// add additional context to the event log
func (l *logger) newEvent(ctx context.Context, event *zerolog.Event) ILogEvent {
	event.Str("hostname", l.settings.Hostname())
	event.Str("appname", l.settings.Appname())

	requestId := ctx.Value(ContextKeyRequestId)
	if requestId != nil {
		event.Str("requestid", requestId.(string))
	}

	return &logEvent{
		event,
	}
}

func (e *logEvent) Send() {
	e.event.Send()
}

func (e *logEvent) Msg(msg string) {
	e.event.Msg(msg)
}

func (e *logEvent) Msgf(msg string, v ...any) {
	e.event.Msgf(msg, v...)
}

func (e *logEvent) Err(err error) ILogEvent {
	e.event.Stack().Err(err)
	return e
}

func (e *logEvent) Strs(strs []struct {
	Key   string
	Value string
}) ILogEvent {
	for _, str := range strs {
		e.event.Str(str.Key, str.Value)
	}
	return e
}
