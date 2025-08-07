package db

import (
	"context"
	"devport/domain/model"
)

type WorkHavingTagsRepository interface {
	DeleteByWorkId(ctx context.Context, workId string) error
	CreateByWorkId(ctx context.Context, workId string, tagIds []string) error
	Find(ctx context.Context, workId string) ([]*model.WorkTag, error)
	UpdateByWorkId(ctx context.Context, workId string, tagIds []string) error
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
