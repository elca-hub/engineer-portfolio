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

	if len(name) > MaxNameLen {
		return nil, fmt.Errorf("名前「%s」は%d字を超過しています", name, MaxNameLen)
	}

	if len(name) == 0 {
		return nil, errors.New("the name must not be empty")
	}

	for _, excludeString := range excludeStrings {
		for _, char := range name {
			if string(char) == excludeString {
				return nil, errors.New("名前に使用できない文字が含まれています")
			}
		}
	}

	nowDate := time.Now()

	if birthDay.After(nowDate) {
		return nil, errors.New("誕生日は未来の日付を指定できません")
	}

	// ageは満何歳かを計算する
	age := nowDate.Year() - birthDay.Year()

	if nowDate.Month() < birthDay.Month() || (nowDate.Month() == birthDay.Month() && nowDate.Day() < birthDay.Day()) {
		age--
	}

	if age < 0 {
		return nil, errors.New("年齢は0歳以上である必要があります")
	}

	if email == nil {
		return nil, errors.New("メールアドレスが指定されていません")
	}

	if password == nil {
		return nil, errors.New("パスワードが指定されていません")
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

func (u *User) Birthday() time.Time {
	return u.birthday
}

func (u *User) UpdateEmailVerification(emailVerification int) {
	u.emailVerification = emailVerification
}
