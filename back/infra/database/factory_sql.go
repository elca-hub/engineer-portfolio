package database

import (
	"devport/domain/repo/db"
	"errors"
)

const (
	InstanceGormMySql int = iota
	InstanceSqlBoilerMySql
)

// 新しいリポジトリを作成したらここに追加する
type SqlInter interface {
	UserRepository() db.UserRepository
	ExternalServiceUrlsRepository() db.ExternalServiceUrlsRepository
	BioImagesRepository() db.BioImagesRepository
	SkillsRepository() db.SkillsRepository
	CertificationsRepository() db.CertificationsRepository
}

func NewDatabaseSqlFactory(instance int) (SqlInter, error) {
	switch instance {
	case InstanceSqlBoilerMySql:
		return NewSqlBoilerHandler(NewMySQLConfig())
	default:
		return nil, errors.New("invalid instance")
	}
}
