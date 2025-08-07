package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"devport/domain/repo/db"
	"devport/infra/database/sqlboiler/repository"
	"devport/infra/database/sqlboiler/seed"

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

	if os.Getenv("GO_ENVIRONMENT") == "development" {
		// seedのデータを作成
		seed.CreateUserSeed(context.Background(), db)
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

func (h *SqlBoilerHandler) SkillsRepository() db.SkillsRepository {
	return repository.NewSqlBoilerSkillsRepository(h.db)
}

func (h *SqlBoilerHandler) CertificationsRepository() db.CertificationsRepository {
	return repository.NewSqlBoilerCertificationRepository(h.db)
}

func (h *SqlBoilerHandler) WorkHavingTagsRepository() db.WorkHavingTagsRepository {
	return repository.NewSqlboilerWorkHavingTagsRepository(h.db)
}

func (h *SqlBoilerHandler) WorkTagRepository() db.WorkTagRepository {
	return repository.NewSqlboilerWorkTagRepository(h.db)
}

func (h *SqlBoilerHandler) WorkUrlRepository() db.WorkUrlRepository {
	return repository.NewSqlboilerWorkUrlRepository(h.db)
}

func (h *SqlBoilerHandler) WorkRepository() db.WorkRepository {
	return repository.NewSqlboilerWorkRepository(h.db)
}

func (h *SqlBoilerHandler) WorkImagesRepository() db.WorkImagesRepository {
	return repository.NewSqlBoilerWorkImagesRepository(h.db)
}
