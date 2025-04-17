package dto

import "devport/domain/model"

type ExternalServiceUrlDTO struct {
	URL         string `json:"url"`
	ID          string `json:"id"`
	ServiceType int    `json:"service_type"`
}

func NewExternalServiceUrlDTO(externalServiceUrl *model.ExternalServiceUrl) *ExternalServiceUrlDTO {
	return &ExternalServiceUrlDTO{
		URL:         externalServiceUrl.Url(),
		ID:          externalServiceUrl.ID(),
		ServiceType: externalServiceUrl.ServiceType(),
	}
}
