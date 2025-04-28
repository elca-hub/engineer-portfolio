package repository

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type SQL interface {
	Execute(ctx context.Context) *gorm.DB
	BeginTx(ctx context.Context) (Tx, error)
	ScanRows(ctx context.Context, rows *sql.Rows, dest interface{}) error
}

type Tx interface {
	Tx() *gorm.DB
}
