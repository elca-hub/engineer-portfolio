//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE
package db

import (
	"context"
	"devport/domain/model"
)

type UserRepository interface {
	Create(context context.Context, u *model.User) error
	Exists(context context.Context, email *model.Email) (bool, error)
	ExistsByName(context context.Context, name string) (bool, error)
	ExistsById(context context.Context, id string) (bool, error)
	Update(context context.Context, u *model.User) error
	FindByEmail(context context.Context, email *model.Email) (*model.User, error)
	FindById(context context.Context, id string) (*model.User, error)
	FetchIconNamesAll(context context.Context) ([]*model.FileIconName, error)
	FetchHeaderNamesAll(context context.Context) ([]*model.FileIconName, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
