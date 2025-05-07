package database

import (
	"devport/domain/repo/nosql"
	"errors"
)

const (
	InstanceRedis int = iota
)

type NoSQLInter interface {
	UserRepository() nosql.UserRepository
}

func NewDatabaseNoSqlFactory(instance int) (NoSQLInter, error) {
	switch instance {
	case InstanceRedis:
		return NewRedisHandler(NewMyNoSQLConfig())
	default:
		return nil, errors.New("invalid instance")
	}
}
