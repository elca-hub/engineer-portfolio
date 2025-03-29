package model

import (
	"errors"
	"time"
)

const (
	SkillStatusUnderOneYear    string = "1年未満"
	SkillStatusOneToThreeYear  string = "1年以上3年未満"
	SkillStatusThreeToFiveYear string = "3年以上5年未満"
	SkillStatusFiveToTenYear   string = "5年以上10年未満"
	SkillStatusOverTenYear     string = "10年以上"
	SkillStatusFetchLicense    string = "保有・合格"
	MaxSkillNameLength         int    = 50
)

type Skill struct {
	name      string
	status    string
	when      time.Time
	sortIndex int
}

func NewSkill(name string, status string, when time.Time, sortIndex int) (*Skill, error) {
	if len(name) > MaxSkillNameLength {
		return nil, errors.New("スキル名は50文字以内で入力してください")
	}

	if len(name) == 0 {
		return nil, errors.New("スキル名を入力してください")
	}

	switch status {
	case SkillStatusUnderOneYear:
	case SkillStatusOneToThreeYear:
	case SkillStatusThreeToFiveYear:
	case SkillStatusFiveToTenYear:
	case SkillStatusOverTenYear:
	case SkillStatusFetchLicense:
	default:
		return nil, errors.New("スキルのステータスが不正です")
	}

	return &Skill{name: name, status: status, when: when, sortIndex: sortIndex}, nil
}

func (s *Skill) Name() string {
	return s.name
}

func (s *Skill) Status() string {
	return s.status
}

func (s *Skill) When() time.Time {
	return s.when
}

func (s *Skill) SortIndex() int {
	return s.sortIndex
}
