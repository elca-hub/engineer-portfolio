package dto

import "devport/domain/model"

type ExternalServiceUrlDTO struct {
	ServiceType string `json:"service_type"`
	ServiceID   int    `json:"service_id"`
	URL         string `json:"url"`
}

func NewExternalServiceUrlDTO(externalServiceUrl *model.ExternalServiceUrl) *ExternalServiceUrlDTO {
	fetch, err := externalServiceUrl.ServiceTypeToString()

	if err != nil {
		return &ExternalServiceUrlDTO{}
	}

	return &ExternalServiceUrlDTO{
		ServiceType: fetch,
		ServiceID:   externalServiceUrl.ServiceTypeToInt(),
		URL:         externalServiceUrl.Url(),
	}
}
