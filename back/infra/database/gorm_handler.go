package database

import (
	"devport/domain/repo/sql"
	"devport/infra/database/gorm/repo"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormHandler struct {
	db *gorm.DB
}

func NewGormHandler(c *MysqlConfig) (*GormHandler, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FTokyo",
		c.user,
		c.password,
		c.host,
		c.port,
		c.database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60))
		},
	})

	if err != nil {
		return nil, err
	}

	return &GormHandler{db: db}, nil
}

func (h *GormHandler) UserRepository() sql.UserRepository {
	return repo.NewGormUserRepository(h.db)
}

// func (h *GormHandler) ExternalServiceUrlsRepository() sql.ExternalServiceUrlsRepository {
// 	return repo.NewGormExternalServiceRepository(h.db)
// }
