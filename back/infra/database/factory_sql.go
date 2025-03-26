package database

import (
	"devport/adapter/repository"
	"errors"
)

const (
	InstanceMySQL int = iota
)

func NewDatabaseSqlFactory(instance int) (repository.SQL, error) {
	switch instance {
	case InstanceMySQL:
		return NewGormHandler(NewMySQLConfig())
	default:
		return nil, errors.New("invalid instance")
	}
}
