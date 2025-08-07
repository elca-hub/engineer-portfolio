//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE
package db

import (
	"context"
	"devport/domain/model"
)

type WorkTagRepository interface {
	Create(context context.Context, tag *model.WorkTag) error
	Exists(context context.Context, tagName string) (bool, error)
	Update(context context.Context, tag *model.WorkTag) error
	FindByName(context context.Context, name string) (*model.WorkTag, error)
	FindById(context context.Context, id string) (*model.WorkTag, error)
	Delete(context context.Context, id string) error
	DeleteNoUsed(context context.Context) error
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
