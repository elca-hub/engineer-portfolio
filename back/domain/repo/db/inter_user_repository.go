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
	FindByEmail(context context.Context, email *model.Email, bio *model.Bio) (*model.User, error)
	FindById(context context.Context, id string, bio *model.Bio) (*model.User, error)
	FetchIconNamesAll(context context.Context) ([]string, error)
	FetchHeaderNamesAll(context context.Context) ([]string, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
