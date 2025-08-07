package model

import (
	"errors"
	"fmt"
)

const (
	MaxWorkTagNameLength = 50
)

type WorkTag struct {
	id   string
	name string
}

func updateWorkTagNameLogic(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("タグ名は必須です")
	}

	if len(name) > MaxWorkTagNameLength {
		return "", fmt.Errorf("タグ名は%d字を超過しています", MaxWorkTagNameLength)
	}

	return name, nil
}

func NewWorkTag(id string, name string) (*WorkTag, error) {
	if _, err := updateWorkTagNameLogic(name); err != nil {
		return nil, err
	}

	return &WorkTag{id: id, name: name}, nil
}

func (t *WorkTag) Name() string {
	return t.name
}

func (t *WorkTag) ID() string {
	return t.id
}
