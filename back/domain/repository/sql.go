package repository

import (
	"context"
	"gorm.io/gorm"
)

type SQL interface {
	DB() *gorm.DB
	BeginTx(ctx context.Context) (Tx, error)
}

type Tx interface {
	Commit() error
	Rollback()
	Tx() *gorm.DB
}
