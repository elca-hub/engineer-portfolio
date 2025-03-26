package database

import (
	"devport/domain/repository"
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
