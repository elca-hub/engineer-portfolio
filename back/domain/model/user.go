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

func updateBirthdayLogic(birthday time.Time) (time.Time, int, error) {
	nowDate := time.Now()

	if birthday.After(nowDate) {
		return time.Time{}, 0, errors.New("誕生日は未来の日付を指定できません")
	}

	// ageは満何歳かを計算する
	age := nowDate.Year() - birthday.Year()

	if nowDate.Month() < birthday.Month() || (nowDate.Month() == birthday.Month() && nowDate.Day() < birthday.Day()) {
		age--
	}

	if age < 0 {
		return time.Time{}, 0, errors.New("年齢は0歳以上である必要があります")
	}

	return birthday, age, nil
}

func updateNameLogic(name string) (string, error) {
	excludeStrings := []string{" ", "@", "#", "$", "%", "&", "", "(", ")", "+", "=", "{", "}", "[", "]", "|", "\\", ":", ";", "\"", "'", "<", ">", ",", ".", "?", "/", "~", "`", "\n", "\t", "/", "?", "%", "#", "&", "=", "--", "/", "*/"}

	if len(name) > MaxUserNameLen {
		return "", fmt.Errorf("名前は%d字を超過しています", MaxUserNameLen)
	}

	if len(name) == 0 {
		return "", errors.New("名前は必ず入力してください")
	}

	for _, excludeString := range excludeStrings {
		for _, char := range name {
			if string(char) == excludeString {
				return "", errors.New("名前に使用できない文字が含まれています")
			}
		}
	}

	return name, nil
}

func updateIDLogic(id string) (string, error) {
	excludeStrings := []string{" ", "@", "#", "$", "%", "&", "", "(", ")", "+", "=", "{", "}", "[", "]", "|", "\\", ":", ";", "\"", "'", "<", ">", ",", ".", "?", "/", "~", "`", "\n", "\t", "/", "?", "%", "#", "&", "=", "--", "/", "*/"}
	if len(id) > MaxUserNameLen {
		return "", fmt.Errorf("IDは%d字を超過しています", MaxUserNameLen)
	}

	if len(id) == 0 {
		return "", errors.New("IDは必ず入力してください")
	}

	for _, excludeString := range excludeStrings {
		for _, char := range id {
			if string(char) == excludeString {
				return "", errors.New("IDに使用できない文字が含まれています")
			}
		}
	}

	return id, nil
}

func updateOccupationNameLogic(occupationName string) (string, error) {
	if len(occupationName) > MaxOccupationNameLen {
		return "", fmt.Errorf("職業名は%d字を超過しています", MaxOccupationNameLen)
	}

	return occupationName, nil
}

func updateOrganizationNameLogic(organizationName string) (string, error) {
	if len(organizationName) > MaxOrganizationName {
		return "", fmt.Errorf("組織名は%d字を超過しています", MaxOrganizationName)
	}

	return organizationName, nil
}

func updatePlaceLogic(place string) (string, error) {
	if len(place) > MaxPlaceLen {
		return "", fmt.Errorf("場所は%d字を超過しています", MaxPlaceLen)
	}

	return place, nil
}

func updateSkillsLogic(skills []*Skill) ([]*Skill, error) {
	if len(skills) > MaxUserSkillsLen {
		return nil, fmt.Errorf("スキルは%d個まで登録できます", MaxUserSkillsLen)
	}

	return skills, nil
}

func updateExternalServiceURLsLogic(externalServiceURLs []*ExternalServiceUrl) ([]*ExternalServiceUrl, error) {
	if len(externalServiceURLs) > MaxUserExternalServiceURLsLen {
		return nil, fmt.Errorf("外部サービスURLは%d個まで登録できます", MaxUserExternalServiceURLsLen)
	}

	return externalServiceURLs, nil
}

func NewUser(
	id string,
	name string,
	birthday time.Time,
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
	name, err := updateNameLogic(name)

	if err != nil {
		return nil, err
	}

	id, err = updateIDLogic(id)

	if err != nil {
		return nil, err
	}

	birthday, age, err := updateBirthdayLogic(birthday)

	if err != nil {
		return nil, err
	}

	if email == nil {
		return nil, errors.New("メールアドレスが指定されていません")
	}

	organizationName, err = updateOrganizationNameLogic(organizationName)

	if err != nil {
		return nil, err
	}

	occupationName, err = updateOccupationNameLogic(occupationName)

	if err != nil {
		return nil, err
	}

	place, err = updatePlaceLogic(place)

	if err != nil {
		return nil, err
	}

	skills, err = updateSkillsLogic(skills)

	if err != nil {
		return nil, err
	}

	externalServiceURLs, err = updateExternalServiceURLsLogic(externalServiceURLs)

	if err != nil {
		return nil, err
	}

	return &User{
		id,
		name,
		birthday,
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

func (u *User) UpdateName(name string) error {
	name, err := updateNameLogic(name)

	if err != nil {
		return err
	}

	u.name = name
	return nil
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

func (u *User) UpdateBirthday(birthday time.Time) error {
	birthday, age, err := updateBirthdayLogic(birthday)

	if err != nil {
		return err
	}

	u.birthday = birthday
	u.age = age
	return nil
}

func (u *User) IconName() string {
	return u.iconName
}

func (u *User) UpdateIconName(iconName string) {
	u.iconName = iconName
}

func (u *User) HeaderIconName() string {
	return u.headerIconName
}

func (u *User) UpdateHeaderIconName(headerIconName string) {
	u.headerIconName = headerIconName
}

func (u *User) BioPath() string {
	return u.bioPath
}

func (u *User) OrganizationName() string {
	return u.organizationName
}

func (u *User) UpdateOrganizationName(organizationName string) error {
	organizationName, err := updateOrganizationNameLogic(organizationName)

	if err != nil {
		return err
	}

	u.organizationName = organizationName
	return nil
}

func (u *User) OccupationName() string {
	return u.occupationName
}

func (u *User) UpdateOccupationName(occupationName string) error {
	occupationName, err := updateOccupationNameLogic(occupationName)

	if err != nil {
		return err
	}

	u.occupationName = occupationName
	return nil
}

func (u *User) Place() string {
	return u.place
}

func (u *User) UpdatePlace(place string) error {
	place, err := updatePlaceLogic(place)

	if err != nil {
		return err
	}

	u.place = place
	return nil
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
