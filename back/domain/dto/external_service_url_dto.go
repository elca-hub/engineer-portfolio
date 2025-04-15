package dto

import "devport/domain/model"

type ExternalServiceUrlDTO struct {
	URL string `json:"url"`
	ID  uint   `json:"id"`
}

func NewExternalServiceUrlDTO(externalServiceUrl *model.ExternalServiceUrl) *ExternalServiceUrlDTO {
	return &ExternalServiceUrlDTO{
		URL: externalServiceUrl.Url(),
		ID:  externalServiceUrl.ID(),
	}
}
