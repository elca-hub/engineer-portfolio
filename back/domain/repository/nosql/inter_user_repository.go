package nosql

import "devport/domain/model"

type UserRepository interface {
	StartSession(email *model.Email) (string, error)
	GetSession(token string) (*model.Email, error)
	DeleteSession(token string) error
}
