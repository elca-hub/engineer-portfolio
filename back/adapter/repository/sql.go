package repository

import (
	"context"
	"gorm.io/gorm"
)

type SQL interface {
	Execute(ctx context.Context) *gorm.DB
	BeginTx(ctx context.Context) (Tx, error)
}

type Tx interface {
	Tx() *gorm.DB
}
