package model

import (
	"errors"
	"regexp"
)

const (
	ExternalServiceOther = iota
	ExternalServiceGithub
	ExternalServiceX
	ExternalServiceQiita
	ExternalServiceZenn
	ExternalServiceNote
)

type ExternalServiceUrl struct {
	url         string
	serviceType int
	id          string
}

func NewExternalServiceUrl(id string, serviceType int, url string) (*ExternalServiceUrl, error) {
	if len(url) == 0 {
		return nil, errors.New("URLを入力してください")
	}

	updatedUrl, err := updateUrlLogic(url)

	if err != nil {
		return nil, err
	}

	return &ExternalServiceUrl{serviceType: serviceType, url: updatedUrl, id: id}, nil
}

func updateUrlLogic(url string) (string, error) {
	if !regexp.MustCompile("^(http|https)://").MatchString(url) {
		return "", errors.New("URLの形式が異なっています")
	}

	return url, nil
}

func (e *ExternalServiceUrl) UpdateUrl(url string) error {
	updatedUrl, err := updateUrlLogic(url)

	if err != nil {
		return err
	}

	e.url = updatedUrl

	return nil
}

func (e *ExternalServiceUrl) Url() string {
	return e.url
}

func (e *ExternalServiceUrl) ServiceType() int {
	return e.serviceType
}

func (e *ExternalServiceUrl) ID() string {
	return e.id
}
