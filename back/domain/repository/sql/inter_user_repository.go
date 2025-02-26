//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package sql

import "devport/domain/model"

type UserRepository interface {
	Create(u *model.User) error
	Exists(email *model.Email) (bool, error)
	Update(u *model.User) error
	FindByEmail(email *model.Email) (*model.User, error)
	FetchInConfirmationUsers() ([]*model.User, error)
}
