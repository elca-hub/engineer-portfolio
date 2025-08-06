package model

import (
	"errors"
	"fmt"
	"time"
)

const (
	MaxSkillNameLength    = 50
	MaxSkillCommentLength = 500
)

type Skill struct {
	id        string
	name      string
	year      time.Time
	comment   string
	sortIndex int
}

func updateSKillNameLogic(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("スキル名は必須です")
	}
	if len(name) > MaxSkillNameLength {
		return "", fmt.Errorf("スキル名は%d字を超過しています", MaxSkillNameLength)
	}
	return name, nil
}

func updateSKillYearLogic(year time.Time) (time.Time, error) {
	if year.IsZero() {
		return time.Time{}, errors.New("スキルの年は必須です")
	}
	return year, nil
}

func updateSkillSortIndexLogic(sortIndex int) (int, error) {
	if sortIndex < 0 {
		return 0, errors.New("ソートインデックスは0以上である必要があります")
	}
	return sortIndex, nil
}

func updateSkillCommentLogic(comment string) (string, error) {
	if len(comment) > MaxSkillCommentLength {
		return "", fmt.Errorf("スキルのコメントは%d字を超過しています", MaxSkillCommentLength)
	}
	return comment, nil
}

func NewSkill(name string, year time.Time, comment string, sortIndex int) (*Skill, error) {
	if _, err := updateSKillNameLogic(name); err != nil {
		return nil, err
	}

	validatedYear, err := updateSKillYearLogic(year)
	if err != nil {
		return nil, err
	}
	year = validatedYear.Truncate(24 * time.Hour) // 年の精度を日単位にする

	if _, err := updateSkillSortIndexLogic(sortIndex); err != nil {
		return nil, err
	}

	if _, err := updateSkillCommentLogic(comment); err != nil {
		return nil, err
	}

	return &Skill{
		id:        "",
		name:      name,
		year:      year,
		comment:   comment,
		sortIndex: sortIndex,
	}, nil
}

func NewSkillWithID(id string, name string, year time.Time, comment string, sortIndex int) (*Skill, error) {
	skill, err := NewSkill(name, year, comment, sortIndex)
	if err != nil {
		return nil, err
	}
	skill.id = id
	return skill, nil
}

func (s *Skill) Name() string {
	return s.name
}

func (s *Skill) Year() time.Time {
	return s.year
}

func (s *Skill) Comment() string {
	return s.comment
}

func (s *Skill) ID() string {
	return s.id
}

func (s *Skill) SortIndex() int {
	return s.sortIndex
}

func (s *Skill) UpdateName(name string) error {
	updatedName, err := updateSKillNameLogic(name)
	if err != nil {
		return err
	}
	s.name = updatedName
	return nil
}

func (s *Skill) UpdateYear(year time.Time) error {
	updatedYear, err := updateSKillYearLogic(year)
	if err != nil {
		return err
	}
	s.year = updatedYear.Truncate(24 * time.Hour) // 年の精度を日単位にする
	return nil
}

func (s *Skill) UpdateSortIndex(sortIndex int) error {
	updatedSortIndex, err := updateSkillSortIndexLogic(sortIndex)
	if err != nil {
		return err
	}
	s.sortIndex = updatedSortIndex
	return nil
}

func (s *Skill) UpdateComment(comment string) error {
	updatedComment, err := updateSkillCommentLogic(comment)
	if err != nil {
		return err
	}
	s.comment = updatedComment
	return nil
}
