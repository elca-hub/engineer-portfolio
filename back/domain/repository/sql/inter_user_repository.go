//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package sql

import (
	"context"
	"devport/domain/model"
)

type UserRepository interface {
	Create(context context.Context, u *model.User) error
	Exists(context context.Context, email *model.Email) (bool, error)
	ExistsByName(context context.Context, name string) (bool, error)
	Update(context context.Context, u *model.User) error
	FindByEmail(context context.Context, email *model.Email) (*model.User, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
