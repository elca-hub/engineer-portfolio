//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package model

import (
	"errors"
	"fmt"
	"time"
)

const (
	Unconfirmed int = iota
	InConfirmation
	Confirmed
	MaxNameLen = 50
)

type User struct {
	id                UUID
	name              string
	birthday          time.Time
	age               int
	email             *Email
	password          *HashedPassword
	createdAt         time.Time
	updatedAt         time.Time
	emailVerification int
}

func NewUser(
	id UUID,
	name string,
	birthDay time.Time,
	email *Email,
	password *HashedPassword,
	createdAt time.Time,
	updatedAt time.Time,
	emailVerification int,
) (*User, error) {
	excludeStrings := []string{" ", "@", "#", "$", "%", "&", "", "(", ")", "+", "=", "{", "}", "[", "]", "|", "\\", ":", ";", "\"", "'", "<", ">", ",", ".", "?", "/", "~", "`", "\n", "\t", "/", "?", "%", "#", "&", "=", "--", "/", "*/"}

	if len(name) >= MaxNameLen {
		return nil, errors.New(fmt.Sprintf("The name '%s' exceeds the maximum length of %d characters.", name, MaxNameLen))
	}

	if len(name) == 0 {
		return nil, errors.New("the name must not be empty")
	}

	for _, excludeString := range excludeStrings {
		for _, char := range name {
			if string(char) == excludeString {
				return nil, errors.New("the name must not contain any special characters")
			}
		}
	}

	nowDate := time.Now()

	if birthDay.After(nowDate) {
		return nil, errors.New("the birthday must be less than the current date")
	}

	// ageは満何歳かを計算する
	age := nowDate.Year() - birthDay.Year()

	if nowDate.Month() < birthDay.Month() || (nowDate.Month() == birthDay.Month() && nowDate.Day() < birthDay.Day()) {
		age--
	}

	if age < 0 {
		return nil, errors.New("the age must be greater than or equal to 0")
	}

	if email == nil {
		return nil, errors.New("the email must not be nil")
	}

	if password == nil {
		return nil, errors.New("the password must not be nil")
	}

	return &User{
		id,
		name,
		birthDay,
		age,
		email,
		password,
		createdAt,
		updatedAt,
		emailVerification,
	}, nil
}

func (u *User) ID() UUID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Age() int {
	return u.age
}

func (u *User) Email() *Email {
	return u.email
}

func (u *User) Password() *HashedPassword {
	return u.password
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) EmailVerification() int {
	return u.emailVerification
}

func (u *User) UpdateEmailVerification(emailVerification int) {
	u.emailVerification = emailVerification
}
