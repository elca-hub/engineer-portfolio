package validation

import (
	"errors"
	"devport/adapter/validator"
)

var (
	errInvalidValidatorInstance = errors.New("invalid validator instance")
)

const (
	InstanceGoPlayground int = iota
)

func NewValidationFactory() (validator.Validator, error) {
	switch InstanceGoPlayground {
	case InstanceGoPlayground:
		return NewGoPlayground()
	default:
		return nil, errInvalidValidatorInstance
	}
}
