//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE
package db

import (
	"context"
	"devport/domain/model"
)

type BioImagesRepository interface {
	Create(context context.Context, userId string, fileName string) error
	Delete(context context.Context, u *model.User, fileName string) error
	DeleteAllByUserId(context context.Context, u *model.User) error
	FindByUserId(context context.Context, userId string) ([]string, error)
	IsExistsFileName(context context.Context, fileName string) (bool, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
