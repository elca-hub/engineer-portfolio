//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package nosql

import "devport/domain/model"

type UserRepository interface {
	StartSession(email *model.Email) (string, error)
	GetSession(token string) (*model.Email, error)
	DeleteSession(token string) error
	AddConfirmationCode(email *model.Email, code int64) error
	GetConfirmationCode(email *model.Email) (int64, error)
}
