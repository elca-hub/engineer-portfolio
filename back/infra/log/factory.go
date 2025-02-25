package log

import (
	"devport/adapter/logger"
	"errors"
)

const (
	InstanceSlog int = iota
)

var (
	errInvalidLoggerInstance = errors.New("invalid log instance")
)

func NewLoggerFactory(instance int) (logger.Logger, error) {
	switch instance {
	case InstanceSlog:
		return NewSlogLogger()
	default:
		return nil, errInvalidLoggerInstance
	}
}
