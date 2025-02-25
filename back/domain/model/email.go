package model

import (
	"errors"
	"regexp"
)

type Email struct {
	email string
}

func NewEmail(email string) (*Email, error) {

	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email) {
		return nil, errors.New("invalid email")
	}

	return &Email{email: email}, nil
}

func (e *Email) Email() string {
	return e.email
}
