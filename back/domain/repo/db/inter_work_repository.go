package db

import (
	"context"
	"devport/domain/model"
	"devport/domain/model/group"
)

type WorkRepository interface {
	Create(ctx context.Context, userId string, work *model.Work) error
	Update(ctx context.Context, userId string, work *model.Work) error
	GetMaxSortIndex(ctx context.Context, userId string) (int, error)
	FindById(ctx context.Context, userId string, id string) (*model.Work, error)
	FindAll(ctx context.Context, userId string) (*group.WorkGroup, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
