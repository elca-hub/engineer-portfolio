package model

import (
	"errors"
	"regexp"
)

type ExternalServiceUrl struct {
	url string
	id  uint
}

func NewExternalServiceUrl(url string, id uint) (*ExternalServiceUrl, error) {
	if len(url) == 0 {
		return nil, errors.New("URLを入力してください")
	}

	if !regexp.MustCompile("^(http|https)://").MatchString(url) {
		return nil, errors.New("URLの形式が異なっています")
	}

	return &ExternalServiceUrl{url: url, id: id}, nil
}

func (e *ExternalServiceUrl) Url() string {
	return e.url
}

func (e *ExternalServiceUrl) ID() uint {
	return e.id
}
