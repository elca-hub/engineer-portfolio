package model

import (
	"errors"
)

type Bio struct {
	content string
}

const (
	MaxBioLen = 6000
)

func NewBio(content string) (*Bio, error) {
	if len(content) > MaxBioLen {
		return nil, errors.New("自己紹介文の内容が長すぎます")
	}

	return &Bio{content: content}, nil
}

func (b *Bio) Content() string {
	return b.content
}
