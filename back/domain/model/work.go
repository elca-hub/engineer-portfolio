package model

import (
	"errors"
	"fmt"
	"strings"
)

const (
	MaxWorkTitleLength        = 100
	MaxExternalServiceUrlsLen = 5
	MaxTagsLen                = 5
)

type Work struct {
	id                  string
	title               string
	content             string
	githubRepositoryUrl string
	externalServiceUrls []*WorkUrl
	tags                []*WorkTag
}

func updateTitleLogic(title string) (string, error) {
	if len(title) == 0 {
		return "", errors.New("タイトルは必須です")
	}
	if len(title) > MaxWorkTitleLength {
		return "", fmt.Errorf("タイトルは%d字を超過しています", MaxWorkTitleLength)
	}
	return title, nil
}

func updateContentLogic(content string) (string, error) {
	if len(content) == 0 {
		return "", errors.New("コンテンツは必須です")
	}
	return content, nil
}

func updateGithubRepositoryUrlLogic(githubRepositoryUrl string) (string, error) {
	if !strings.HasPrefix(githubRepositoryUrl, "https://github.com") {
		return "", errors.New("GithubリポジトリURLはhttps://github.comで始まる必要があります")
	}
	return githubRepositoryUrl, nil
}

func updateExternalServiceUrlsLogic(externalServiceUrls []*WorkUrl) ([]*WorkUrl, error) {
	if len(externalServiceUrls) > MaxExternalServiceUrlsLen {
		return nil, fmt.Errorf("外部サービスURLは%d個まで登録できます", MaxExternalServiceUrlsLen)
	}

	return externalServiceUrls, nil
}

func updateTagsLogic(tags []*WorkTag) ([]*WorkTag, error) {
	if len(tags) > MaxTagsLen {
		return nil, fmt.Errorf("タグは%d個まで登録できます", MaxTagsLen)
	}
	return tags, nil
}

func NewWork(
	id string,
	title string,
	content string,
	githubRepositoryUrl string,
	externalServiceUrls []*WorkUrl,
	tags []*WorkTag,
) (*Work, error) {
	if _, err := updateTitleLogic(title); err != nil {
		return nil, err
	}

	if _, err := updateContentLogic(content); err != nil {
		return nil, err
	}

	if _, err := updateGithubRepositoryUrlLogic(githubRepositoryUrl); err != nil {
		return nil, err
	}

	if _, err := updateExternalServiceUrlsLogic(externalServiceUrls); err != nil {
		return nil, err
	}

	if _, err := updateTagsLogic(tags); err != nil {
		return nil, err
	}

	return &Work{
		id:                  id,
		title:               title,
		content:             content,
		githubRepositoryUrl: githubRepositoryUrl,
		externalServiceUrls: externalServiceUrls,
		tags:                tags,
	}, nil
}
