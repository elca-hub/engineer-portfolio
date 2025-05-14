//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE
package db

import (
	"context"
	"devport/domain/model"
)

type BioImagesRepository interface {
	Create(context context.Context, u *model.User, fileName *model.FileIconName) error
	Delete(context context.Context, u *model.User, fileName *model.FileIconName) error
	DeleteAllByUserId(context context.Context, u *model.User) error
	FindByUserId(context context.Context, u *model.User) ([]*model.FileIconName, error)
	IsExistsFileName(context context.Context, fileName *model.FileIconName) (bool, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
