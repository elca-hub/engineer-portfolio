package database

import (
	"context"
	"devport/adapter/repository"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormHandler struct {
	db *gorm.DB
}

func NewGormHandler(c *MysqlConfig) (*GormHandler, error) {
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

	return &GormHandler{db: db}, nil
}

func (h GormHandler) BeginTx(ctx context.Context) (repository.Tx, error) {
	tx := h.db.WithContext(ctx).Begin()

	return newGormTx(tx), tx.Error
}

func (h GormHandler) Execute(ctx context.Context) *gorm.DB {
	return h.db.WithContext(ctx)
}

type gormTx struct {
	tx *gorm.DB
}

func newGormTx(tx *gorm.DB) repository.Tx {
	return gormTx{tx: tx}
}

func (t gormTx) Tx() *gorm.DB {
	return t.tx
}
