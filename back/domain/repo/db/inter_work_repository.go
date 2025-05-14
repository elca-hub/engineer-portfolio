//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE
package db

import (
	"context"
	"devport/domain/model"
	"time"
)

type WorkFindResponse struct {
	ID        string
	UserID    string
	Title     string
	SortIndex int
	GitHubURL string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WorkRepository interface {
	Create(context context.Context, w *model.Work) error
	FindById(context context.Context, id string) (*WorkFindResponse, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
