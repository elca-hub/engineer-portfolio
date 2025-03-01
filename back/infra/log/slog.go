package log

import (
	"devport/adapter/logger"
	"log/slog"
	"os"
)

type slogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger() (logger.Logger, error) {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return &slogLogger{logger: log}, nil
}

func (l *slogLogger) Infof(format string, args ...interface{}) {
	l.logger.Info(format, args...)
}

func (l *slogLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warn(format, args...)
}

func (l *slogLogger) Errorf(format string, args ...interface{}) {
	l.logger.Error(format, args...)
}

func (l *slogLogger) WithFields(keyValues logger.Fields) logger.Logger {
	for k, v := range keyValues {
		l.logger = l.logger.With(k, v)
	}

	return &slogLogger{logger: l.logger}
}

func (l *slogLogger) WithError(err error) logger.Logger {
	var log = l.logger.With(err.Error())
	return &slogLogger{logger: log}
}
