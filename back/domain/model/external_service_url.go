package model

import (
	"errors"
	"regexp"
)

type ExternalServiceUrl struct {
	name string
	url  string
}

const (
	MaxExternalServiceNameLength = 100
)

func NewExternalServiceUrl(name, url string) (*ExternalServiceUrl, error) {
	if len(name) > MaxExternalServiceNameLength {
		return nil, errors.New("外部サービス名は100文字以内で入力してください")
	}

	if len(name) == 0 {
		return nil, errors.New("外部サービス名を入力してください")
	}

	if len(url) == 0 {
		return nil, errors.New("URLを入力してください")
	}

	if !regexp.MustCompile("^(http|https)://").MatchString(url) {
		return nil, errors.New("URLの形式が異なっています")
	}

	return &ExternalServiceUrl{name: name, url: url}, nil
}

func (e *ExternalServiceUrl) Name() string {
	return e.name
}

func (e *ExternalServiceUrl) Url() string {
	return e.url
}
