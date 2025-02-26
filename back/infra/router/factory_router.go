package router

import (
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/domain/repository"
	"devport/infra/email"
	"fmt"
	"time"
)

type Server interface {
	Listen()
}

type Port int64

const (
	InstanceGin int = iota
)

func NewWebServerFactory(
	instance int,
	port Port,
	ctxTimeout time.Duration,
	db repository.SQL,
	nosqlDb repository.NoSQL,
	validator validator.Validator,
	logger logger.Logger,
	email email.Email,
) (Server, error) {
	switch instance {
	case InstanceGin:
		return NewGinServer(port, ctxTimeout, db, validator, logger, nosqlDb, email), nil
	default:
		return nil, fmt.Errorf("instance not exist")
	}
}
