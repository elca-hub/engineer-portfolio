package database

import (
	"database/sql"
	"fmt"

	"devport/domain/repo/db"
	"devport/infra/database/sqlboiler/repository"

	_ "github.com/go-sql-driver/mysql"
)

type SqlBoilerHandler struct {
	db *sql.DB
}

func NewSqlBoilerHandler(c *MysqlConfig) (*SqlBoilerHandler, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FTokyo",
		c.user,
		c.password,
		c.host,
		c.port,
		c.database,
	)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &SqlBoilerHandler{db: db}, nil
}

func (h *SqlBoilerHandler) UserRepository() db.UserRepository {
	return repository.NewSqlBoilerUserRepository(h.db)
}

func (h *SqlBoilerHandler) ExternalServiceUrlsRepository() db.ExternalServiceUrlsRepository {
	return repository.NewSqlBoilerExternalServiceUrlRepository(h.db)
}

func (h *SqlBoilerHandler) BioImagesRepository() db.BioImagesRepository {
	return repository.NewSqlBoilerBioImagesRepository(h.db)
}
