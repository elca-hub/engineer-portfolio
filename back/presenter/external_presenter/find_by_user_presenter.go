package external_presenter

import (
	"devport/domain/dto"
	"devport/domain/model"
	"devport/usecase/external_service_url"
)

type FindByUserExternalServiceUrlPresenter struct{}

func NewFindByUserExternalServiceUrlPresenter() external_service_url.FindByUserExternalServiceUrlPresenter {
	return FindByUserExternalServiceUrlPresenter{}
}

func (p FindByUserExternalServiceUrlPresenter) Output(esus []*model.ExternalServiceUrl) external_service_url.FindByUserExternalServiceUrlOutput {
	esuDtos := make([]*dto.ExternalServiceUrlDTO, len(esus))
	for i, esu := range esus {
		esuDtos[i] = dto.NewExternalServiceUrlDTO(esu)
	}
	return external_service_url.FindByUserExternalServiceUrlOutput{
		ExternalServiceUrls: esuDtos,
	}
}
