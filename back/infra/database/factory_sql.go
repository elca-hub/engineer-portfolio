package database

import (
	"devport/domain/repo/sql"
	"errors"
)

const (
	InstanceGormMySql int = iota
	InstanceSqlBoilerMySql
)

// 新しいリポジトリを作成したらここに追加する
type SqlInter interface {
	UserRepository() sql.UserRepository
	ExternalServiceUrlsRepository() sql.ExternalServiceUrlsRepository
	BioImagesRepository() sql.BioImagesRepository
}

func NewDatabaseSqlFactory(instance int) (SqlInter, error) {
	switch instance {
	// case InstanceGormMySql:
	// 	return NewGormHandler(NewMySQLConfig())
	case InstanceSqlBoilerMySql:
		return NewSqlBoilerHandler(NewMySQLConfig())
	default:
		return nil, errors.New("invalid instance")
	}
}
