package model

import (
	"errors"
	"fmt"
	"time"
)

const (
	MaxWorkTitleLength   = 255
	MaxWorkContentLength = 10000
)

type Work struct {
	id                string
	userId            string
	title             string
	content           string
	contentObjectName *FileIconName
	sortIndex         int
	githubURL         string
	thumbnailURL      *FileIconName
	createdAt         time.Time
	updatedAt         time.Time
}

func updateTitleLogic(title string) (string, error) {
	if len(title) > MaxWorkTitleLength {
		return "", errors.New("タイトルは255文字以内で入力してください")
	}
	return title, nil
}

func updateContentLogic(content string) (string, error) {
	if len(content) > MaxWorkContentLength {
		return "", errors.New("内容は10000文字以内で入力してください")
	}
	return content, nil
}

func NewWork(id string, userId string, title string, content string, thumbnailURL *FileIconName, sortIndex int, githubURL string, createdAt time.Time, updatedAt time.Time) (*Work, error) {
	title, err := updateTitleLogic(title)
	if err != nil {
		return nil, err
	}

	content, err = updateContentLogic(content)
	if err != nil {
		return nil, err
	}

	contentObjectName, err := NewFileName(fmt.Sprintf("%s/content.md", id), WORK_CONTENT_PATH)
	if err != nil {
		return nil, err
	}

	return &Work{id: id, userId: userId, title: title, content: content, contentObjectName: contentObjectName, thumbnailURL: thumbnailURL, sortIndex: sortIndex, githubURL: githubURL, createdAt: createdAt, updatedAt: updatedAt}, nil
}

func (e *Work) ID() string {
	return e.id
}

func (e *Work) UserID() string {
	return e.userId
}

func (e *Work) Title() string {
	return e.title
}

func (e *Work) Content() string {
	return e.content
}

func (e *Work) ContentFileName() *FileIconName {
	return e.contentObjectName
}

func (e *Work) SortIndex() int {
	return e.sortIndex
}

func (e *Work) GithubURL() string {
	return e.githubURL
}

func (e *Work) ThumbnailURL() *FileIconName {
	return e.thumbnailURL
}

func (e *Work) Pinned() bool {
	return e.sortIndex != -1
}

func (e *Work) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Work) UpdatedAt() time.Time {
	return e.updatedAt
}
