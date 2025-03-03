package log

import (
	"devport/adapter/logger"
	"errors"
)

const (
	InstanceSlog int = iota
	InstanceZap
)

var (
	errInvalidLoggerInstance = errors.New("invalid log instance")
)

func NewLoggerFactory(instance int) (logger.Logger, error) {
	switch instance {
	case InstanceSlog:
		return NewSlogLogger()
	case InstanceZap:
		return NewZapLogger()
	default:
		return nil, errInvalidLoggerInstance
	}
}
