package database

import (
	"errors"
	"devport/domain/repository"
)

const (
	InstanceMySQL int = iota
)

func NewDatabaseSqlFactory(instance int) (repository.SQL, error) {
	switch instance {
	case InstanceMySQL:
		return NewMysqlHandler(NewMySQLConfig())
	default:
		return nil, errors.New("invalid instance")
	}
}
