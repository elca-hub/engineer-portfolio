package model

import (
	"errors"
	"fmt"
	"net/url"
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

	updatedServiceType, err := updateServiceType(serviceType)

	if err != nil {
		return nil, err
	}

	updatedUrl, err := updateUrlLogic(url, updatedServiceType)

	if err != nil {
		return nil, err
	}

	return &ExternalServiceUrl{serviceType: updatedServiceType, url: updatedUrl, id: id}, nil
}

func updateUrlLogic(urlRaw string, serviceType int) (string, error) {
	_, err := updateServiceType(serviceType) // serviceTypeのvalidate

	if err != nil {
		return "", err
	}

	// URLの形式が異なっている場合エラー
	parsedURL, err := url.ParseRequestURI(urlRaw)
	if err != nil {
		return "", fmt.Errorf("無効なURL形式です: %w", err)
	}
	// スキーム（http/httpsなど）とホストが存在しているかもチェックする
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return "", fmt.Errorf("URLにスキームまたはホストが不足しています")
	}

	hostname := parsedURL.Hostname()

	type combination struct {
		serviceType int
		hostname    string
	}

	combinations := []combination{
		{
			serviceType: ExternalServiceGithub,
			hostname:    "github.com",
		},
		{
			serviceType: ExternalServiceX,
			hostname:    "x.com",
		},
		{
			serviceType: ExternalServiceZenn,
			hostname:    "zenn.dev",
		},
		{
			serviceType: ExternalServiceNote,
			hostname:    "note.com",
		},
		{
			serviceType: ExternalServiceQiita,
			hostname:    "qiita.com",
		},
	}

	for _, c := range combinations {
		if c.serviceType == serviceType && c.hostname != hostname {
			return "", fmt.Errorf("servicetypeは%dですが、ホスト名が%sで異なります", serviceType, hostname)
		}
	}

	return urlRaw, nil
}

func updateServiceType(serviceType int) (int, error) {
	switch serviceType {
	case ExternalServiceOther:
	case ExternalServiceGithub:
	case ExternalServiceX:
	case ExternalServiceQiita:
	case ExternalServiceZenn:
	case ExternalServiceNote:
	default:
		return 0, fmt.Errorf("unknown service type:%d", serviceType)
	}
	return serviceType, nil
}

func (e *ExternalServiceUrl) UpdateUrl(url string) error {
	updatedUrl, err := updateUrlLogic(url, e.serviceType)

	if err != nil {
		return err
	}

	e.url = updatedUrl

	return nil
}

func (e *ExternalServiceUrl) UpdateServiceType(serviceType int) error {
	updatedServiceType, err := updateServiceType(serviceType)
	if err != nil {
		return err
	}

	e.serviceType = updatedServiceType
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
