//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package sql

import (
	"context"
	"devport/domain/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(tx *gorm.DB, u *model.User) error
	Exists(tx *gorm.DB, email *model.Email) (bool, error)
	ExistsByName(tx *gorm.DB, name string) (bool, error)
	Update(tx *gorm.DB, u *model.User) error
	FindByEmail(tx *gorm.DB, email *model.Email) (*model.User, error)
	WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error
}
