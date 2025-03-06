package database

import (
	"devport/domain/repository"
	"devport/domain/repository/sql"
	"devport/infra/database/gorm/gorm_model"
	gormrepository "devport/infra/database/gorm/repository"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RepositoryConfig struct {
	db *gorm.DB
}

func NewRepositoryConfig(db *gorm.DB) *RepositoryConfig {
	return &RepositoryConfig{
		db: db,
	}
}

func (c *RepositoryConfig) UserRepository() sql.UserRepository {
	return gormrepository.NewGormUserRepository(c.db)
}

func NewMysqlHandler(c *MysqlConfig) (repository.SQL, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		c.user,
		c.password,
		c.host,
		c.port,
		c.database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&gorm_model.User{})
	if err != nil {
		return nil, err
	}

	return NewRepositoryConfig(db), nil
}
