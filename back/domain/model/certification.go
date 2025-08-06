package model

import (
	"errors"
	"fmt"
	"time"
)

const (
	MaxCertificationNameLength    = 50
	MaxCertificationCommentLength = 500
)

type Certification struct {
	id        string
	name      string
	year      time.Time
	comment   string
	sortIndex int
}

func updateCertificationNameLogic(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("資格名は必須です")
	}
	if len(name) > MaxCertificationNameLength {
		return "", fmt.Errorf("資格名は%d字を超過しています", MaxCertificationNameLength)
	}
	return name, nil
}

func updateCertificationYearLogic(year time.Time) (time.Time, error) {
	if year.IsZero() {
		return time.Time{}, errors.New("資格の年は必須です")
	}
	return year, nil
}

func updateCertificationSortIndexLogic(sortIndex int) (int, error) {
	if sortIndex < 0 {
		return 0, errors.New("ソートインデックスは0以上である必要があります")
	}
	return sortIndex, nil
}

func updateCertificationCommentLogic(comment string) (string, error) {
	if len(comment) > MaxCertificationCommentLength {
		return "", fmt.Errorf("資格のコメントは%d字を超過しています", MaxCertificationCommentLength)
	}
	return comment, nil
}

func NewCertification(name string, year time.Time, comment string, sortIndex int) (*Certification, error) {
	if _, err := updateCertificationNameLogic(name); err != nil {
		return nil, err
	}

	validatedYear, err := updateCertificationYearLogic(year)
	if err != nil {
		return nil, err
	}
	year = validatedYear.Truncate(24 * time.Hour) // 年の精度を日単位にする

	if _, err := updateCertificationSortIndexLogic(sortIndex); err != nil {
		return nil, err
	}

	if _, err := updateCertificationCommentLogic(comment); err != nil {
		return nil, err
	}

	return &Certification{
		id:        "", // IDは作成時には空文字、リポジトリでセット
		name:      name,
		year:      year,
		comment:   comment,
		sortIndex: sortIndex,
	}, nil
}

func NewCertificationWithID(id string, name string, year time.Time, comment string, sortIndex int) (*Certification, error) {
	certification, err := NewCertification(name, year, comment, sortIndex)
	if err != nil {
		return nil, err
	}
	certification.id = id
	return certification, nil
}

func (s *Certification) ID() string {
	return s.id
}

func (s *Certification) Name() string {
	return s.name
}

func (s *Certification) Year() time.Time {
	return s.year
}

func (s *Certification) Comment() string {
	return s.comment
}

func (s *Certification) SortIndex() int {
	return s.sortIndex
}

func (s *Certification) UpdateName(name string) error {
	updatedName, err := updateCertificationNameLogic(name)
	if err != nil {
		return err
	}
	s.name = updatedName
	return nil
}

func (s *Certification) UpdateYear(year time.Time) error {
	updatedYear, err := updateCertificationYearLogic(year)
	if err != nil {
		return err
	}
	s.year = updatedYear.Truncate(24 * time.Hour) // 年の精度を日単位にする
	return nil
}

func (s *Certification) UpdateSortIndex(sortIndex int) error {
	updatedSortIndex, err := updateCertificationSortIndexLogic(sortIndex)
	if err != nil {
		return err
	}
	s.sortIndex = updatedSortIndex
	return nil
}

func (s *Certification) UpdateComment(comment string) error {
	updatedComment, err := updateCertificationCommentLogic(comment)
	if err != nil {
		return err
	}
	s.comment = updatedComment
	return nil
}
