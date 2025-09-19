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

type PublishStatus string

const (
	PublishStatusDraft   PublishStatus = "draft"
	PublishStatusPrivate PublishStatus = "private"
	PublishStatusLimited PublishStatus = "limited"
	PublishStatusPublic  PublishStatus = "public"
)

type Work struct {
	id                  string
	title               string
	content             string
	githubRepositoryUrl string
	externalServiceUrls []*WorkUrl
	tags                []string
	isDraft             bool
	sortIndex           int
	thumbnailImageUrl   *string // サムネイル画像がnilの場合はデフォルトの画像を使用
	publishStatus       PublishStatus
}

func updateTitleLogic(title string) (string, error) {
	runeSlice := []rune(title)
	if len(runeSlice) > MaxWorkTitleLength {
		return "", fmt.Errorf("タイトルは%d字を超過しています", MaxWorkTitleLength)
	}
	return title, nil
}

func updateContentLogic(content string) (string, error) {
	return content, nil
}

func updateGithubRepositoryUrlLogic(githubRepositoryUrl string) (string, error) {
	// 空文字列の場合は許可
	if githubRepositoryUrl == "" {
		return githubRepositoryUrl, nil
	}

	// GitHubのURLかどうかをチェック
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

func updateTagsLogic(tags []string) ([]string, error) {
	if len(tags) > MaxTagsLen {
		return nil, fmt.Errorf("タグは%d個まで登録できます", MaxTagsLen)
	}
	return tags, nil
}

/**
* workを作成する際に使用
 */
func NewWorkInit(
	id string,
	sortIndex int,
) *Work {
	return &Work{
		id:                  id,
		title:               "",
		content:             "",
		githubRepositoryUrl: "",
		externalServiceUrls: []*WorkUrl{},
		tags:                []string{},
		isDraft:             true,
		sortIndex:           sortIndex,
		publishStatus:       PublishStatusDraft,
	}
}

func NewWork(
	id string,
	title string,
	content string,
	githubRepositoryUrl string,
	externalServiceUrls []*WorkUrl,
	tags []string,
	isDraft bool,
	sortIndex int,
	thumbnailImageUrl *string,
	publishStatus PublishStatus,
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
		isDraft:             isDraft,
		sortIndex:           sortIndex,
		thumbnailImageUrl:   thumbnailImageUrl,
		publishStatus:       publishStatus,
	}, nil
}

func (w *Work) UpdateTitle(title string) error {
	if _, err := updateTitleLogic(title); err != nil {
		return err
	}
	w.title = title
	return nil
}

func (w *Work) UpdateContent(content string) error {
	if _, err := updateContentLogic(content); err != nil {
		return err
	}
	w.content = content
	return nil
}

func (w *Work) UpdateGithubRepositoryUrl(githubRepositoryUrl string) error {
	if _, err := updateGithubRepositoryUrlLogic(githubRepositoryUrl); err != nil {
		return err
	}
	w.githubRepositoryUrl = githubRepositoryUrl
	return nil
}

func (w *Work) UpdateExternalServiceUrls(externalServiceUrls []*WorkUrl) error {
	if _, err := updateExternalServiceUrlsLogic(externalServiceUrls); err != nil {
		return err
	}
	w.externalServiceUrls = externalServiceUrls
	return nil
}

func (w *Work) UpdateTags(tags []string) error {
	if _, err := updateTagsLogic(tags); err != nil {
		return err
	}
	w.tags = tags
	return nil
}

func (w *Work) UpdateThumbnailImageUrl(thumbnailImageUrl *string) error {
	w.thumbnailImageUrl = thumbnailImageUrl
	return nil
}

func (w *Work) IsDraft() bool {
	return w.isDraft
}

func (w *Work) ID() string {
	return w.id
}

func (w *Work) Tags() []string {
	return w.tags
}

func (w *Work) Title() string {
	return w.title
}

func (w *Work) Content() string {
	return w.content
}

func (w *Work) GithubRepositoryUrl() string {
	return w.githubRepositoryUrl
}

func (w *Work) ExternalServiceUrls() []*WorkUrl {
	return w.externalServiceUrls
}

func (w *Work) SortIndex() int {
	return w.sortIndex
}

func (w *Work) ThumbnailImageUrl() *string {
	return w.thumbnailImageUrl
}

func (w *Work) PublishStatus() PublishStatus {
	return w.publishStatus
}
