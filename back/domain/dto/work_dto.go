package dto

import (
	"devport/domain/model"
)

type WorkDTO struct {
	ID                  string        `json:"id"`
	Title               string        `json:"title"`
	Content             string        `json:"content"`
	GithubRepositoryUrl string        `json:"github_repository_url"`
	ExternalServiceUrls []*WorkUrlDTO `json:"external_service_urls"`
	Tags                []*WorkTagDTO `json:"tags"`
	IsDraft             bool          `json:"is_draft"`
	ThumbnailImageUrl   *string       `json:"thumbnail_image_url"`
	PublishStatus       string        `json:"publish_status"`
}

func NewWorkDTO(work *model.Work) *WorkDTO {
	esuDto := make([]*WorkUrlDTO, len(work.ExternalServiceUrls()))
	for i, url := range work.ExternalServiceUrls() {
		esuDto[i] = NewWorkUrlDTO(url)
	}

	tagDto := make([]*WorkTagDTO, len(work.Tags()))
	for i, tag := range work.Tags() {
		tagDto[i] = NewWorkTagDTO(tag)
	}

	return &WorkDTO{
		ID:                  work.ID(),
		Title:               work.Title(),
		Content:             work.Content(),
		GithubRepositoryUrl: work.GithubRepositoryUrl(),
		ExternalServiceUrls: esuDto,
		Tags:                tagDto,
		IsDraft:             work.IsDraft(),
		ThumbnailImageUrl:   work.ThumbnailImageUrl(),
		PublishStatus:       string(work.PublishStatus()),
	}
}
