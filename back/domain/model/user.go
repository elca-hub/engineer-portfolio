//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package model

import (
	"errors"
	"fmt"
	"time"
)

const (
	MaxUserNameLen                = 50
	MaxUserSkillsLen              = 50
	MaxUserExternalServiceURLsLen = 5
	MaxOrganizationName           = 50
	MaxOccupationNameLen          = 50
	MaxPlaceLen                   = 50
)

type User struct {
	id                  string
	name                string
	birthday            time.Time
	age                 int
	email               *Email
	iconName            string
	headerIconName      string
	bioPath             string
	organizationName    string
	occupationName      string
	place               string
	createdAt           time.Time
	updatedAt           time.Time
	skills              []*Skill
	externalServiceURLs []*ExternalServiceUrl
}

func NewUser(
	id string,
	name string,
	birthDay time.Time,
	email *Email,
	iconName string,
	headerIconName string,
	bioPath string,
	organizationName string,
	occupationName string,
	place string,
	createdAt time.Time,
	updatedAt time.Time,
	skills []*Skill,
	externalServiceURLs []*ExternalServiceUrl,
) (*User, error) {
	excludeStrings := []string{" ", "@", "#", "$", "%", "&", "", "(", ")", "+", "=", "{", "}", "[", "]", "|", "\\", ":", ";", "\"", "'", "<", ">", ",", ".", "?", "/", "~", "`", "\n", "\t", "/", "?", "%", "#", "&", "=", "--", "/", "*/"}

	if len(name) > MaxUserNameLen {
		return nil, fmt.Errorf("名前「%s」は%d字を超過しています", name, MaxUserNameLen)
	}

	if len(name) == 0 {
		return nil, errors.New("名前は必ず入力してください")
	}

	for _, excludeString := range excludeStrings {
		for _, char := range name {
			if string(char) == excludeString {
				return nil, errors.New("名前に使用できない文字が含まれています")
			}
		}
	}

	if len(id) > MaxUserNameLen {
		return nil, fmt.Errorf("ID「%s」は%d字を超過しています", id, MaxUserNameLen)
	}

	if len(id) == 0 {
		return nil, errors.New("IDは必ず入力してください")
	}

	for _, excludeString := range excludeStrings {
		for _, char := range id {
			if string(char) == excludeString {
				return nil, errors.New("IDに使用できない文字が含まれています")
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

	if len(organizationName) > MaxOrganizationName {
		return nil, fmt.Errorf("組織名は%d字を超過しています", MaxOrganizationName)
	}

	if len(occupationName) > MaxOccupationNameLen {
		return nil, fmt.Errorf("職業名は%d字を超過しています", MaxOccupationNameLen)
	}

	if len(place) > MaxPlaceLen {
		return nil, fmt.Errorf("場所は%d字を超過しています", MaxPlaceLen)
	}

	if len(skills) > MaxUserSkillsLen {
		return nil, fmt.Errorf("スキルは%d個まで登録できます", MaxUserSkillsLen)
	}

	if len(externalServiceURLs) > MaxUserExternalServiceURLsLen {
		return nil, fmt.Errorf("外部サービスURLは%d個まで登録できます", MaxUserExternalServiceURLsLen)
	}

	return &User{
		id,
		name,
		birthDay,
		age,
		email,
		iconName,
		headerIconName,
		bioPath,
		organizationName,
		occupationName,
		place,
		createdAt,
		updatedAt,
		skills,
		externalServiceURLs,
	}, nil
}

func (u *User) SetIconName(iconName string) {
	u.iconName = iconName
}

func (u *User) ID() string {
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

func (u *User) Birthday() time.Time {
	return u.birthday
}

func (u *User) IconName() string {
	return u.iconName
}

func (u *User) HeaderIconName() string {
	return u.headerIconName
}

func (u *User) BioPath() string {
	return u.bioPath
}

func (u *User) OrganizationName() string {
	return u.organizationName
}

func (u *User) OccupationName() string {
	return u.occupationName
}

func (u *User) Place() string {
	return u.place
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) Skills() []*Skill {
	return u.skills
}

func (u *User) ExternalServiceURLs() []*ExternalServiceUrl {
	return u.externalServiceURLs
}
