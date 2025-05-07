package database

import (
	"database/sql"
	"fmt"

	sql_inter "devport/domain/repo/sql"
	"devport/infra/database/sqlboiler/repository"
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

func (h *SqlBoilerHandler) UserRepository() sql_inter.UserRepository {
	return repository.NewSqlBoilerUserRepository(h.db)
}

func (h *SqlBoilerHandler) ExternalServiceUrlsRepository() sql_inter.ExternalServiceUrlsRepository {
	return repository.NewSqlBoilerExternalServiceUrlRepository(h.db)
}
