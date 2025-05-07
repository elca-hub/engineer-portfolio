package router

import (
	"devport/adapter/logger"
	"devport/adapter/validator"
	"devport/infra/database"
	"devport/infra/email"
	"devport/infra/file_uploader"
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
	db database.SqlInter,
	nosqlDb database.NoSQLInter,
	validator validator.Validator,
	logger logger.Logger,
	email email.Email,
	fileUploader file_uploader.FileUploader,
) (Server, error) {
	switch instance {
	case InstanceGin:
		return NewGinServer(port, ctxTimeout, db, validator, logger, nosqlDb, email, fileUploader), nil
	default:
		return nil, fmt.Errorf("instance not exist")
	}
}
