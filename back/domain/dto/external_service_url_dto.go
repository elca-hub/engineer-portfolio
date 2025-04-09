package dto

import "devport/domain/model"

type ExternalServiceUrlDTO struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func NewExternalServiceUrlDTO(externalServiceUrl *model.ExternalServiceUrl) *ExternalServiceUrlDTO {
	return &ExternalServiceUrlDTO{
		Name: externalServiceUrl.Name(),
		Url:  externalServiceUrl.Url(),
	}
}
