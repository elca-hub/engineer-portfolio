package model

import (
	"errors"
	"regexp"
)

type ExternalServiceUrl struct {
	serviceType int
	url         string
}

const (
	ServiceTypeQiita = iota
	ServiceTypeZenn
	ServiceTypeNote
	ServiceTypeGithub
	ServiceTypeX
	ServiceTypeOther
)

func NewExternalServiceUrl(serviceType int, url string) (*ExternalServiceUrl, error) {
	if len(url) == 0 {
		return nil, errors.New("URLを入力してください")
	}

	if !regexp.MustCompile("^(http|https)://").MatchString(url) {
		return nil, errors.New("URLの形式が異なっています")
	}

	if serviceType < 0 || serviceType > ServiceTypeOther {
		return nil, errors.New("サービスタイプの値が不正です")
	}

	return &ExternalServiceUrl{serviceType: serviceType, url: url}, nil
}

func (e *ExternalServiceUrl) ServiceTypeToString() (string, error) {
	switch e.serviceType {
	case ServiceTypeQiita:
		return "Qiita", nil
	case ServiceTypeNote:
		return "note", nil
	case ServiceTypeZenn:
		return "zenn", nil
	case ServiceTypeGithub:
		return "GitHub", nil
	case ServiceTypeX:
		return "X", nil
	case ServiceTypeOther:
		return "other", nil
	default:
		return "", errors.New("サービスタイプの変換に失敗しました")
	}
}

func (e *ExternalServiceUrl) ServiceTypeToInt() int {
	return e.serviceType
}

func (e *ExternalServiceUrl) Url() string {
	return e.url
}
