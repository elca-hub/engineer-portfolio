package database

import (
	"devport/domain/repo/sql"
	"errors"
)

const (
	InstanceGormMySql int = iota
	InstanceSqlBoilerMySql
)

type SqlInter interface {
	UserRepository() sql.UserRepository
	ExternalServiceUrlsRepository() sql.ExternalServiceUrlsRepository
}

func NewDatabaseSqlFactory(instance int) (SqlInter, error) {
	switch instance {
	case InstanceGormMySql:
		return NewGormHandler(NewMySQLConfig())
	case InstanceSqlBoilerMySql:
		return NewSqlBoilerHandler(NewMySQLConfig())
	default:
		return nil, errors.New("invalid instance")
	}
}
