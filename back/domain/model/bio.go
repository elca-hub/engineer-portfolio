package model

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

type Bio struct {
	id         string
	content    string
	objectName string
	imageIds   []string
}

const (
	MaxBioLen       = 1000
	MaxBioImagesLen = 5
)

func NewBio(userId string, id string, content string) (*Bio, error) {
	if len(content) > MaxBioLen {
		return nil, errors.New("自己紹介文の内容が長すぎます")
	}

	if id == "" {
		id = uuid.New().String()
	}

	imageIds := []string{}

	// markdown形式のcontentから画像のidを取得
	re := regexp.MustCompile(`!\[.*?\]\((.*?)\)`)
	matches := re.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		imageIds = append(imageIds, match[1])
	}

	if len(imageIds) > MaxBioImagesLen {
		return nil, errors.New("自己紹介文に使用している画像が多すぎます")
	}

	objectName := fmt.Sprintf("bio/%s/%s.md", userId, id)

	return &Bio{id: id, content: content, objectName: objectName, imageIds: imageIds}, nil
}

func (b *Bio) ID() string {
	return b.id
}

func (b *Bio) Content() string {
	return b.content
}

func (b *Bio) ObjectName() string {
	return b.objectName
}

func (b *Bio) ImageIds() []string {
	return b.imageIds
}

func (b *Bio) IsFullImage() bool {
	return len(b.imageIds) == MaxBioImagesLen
}
