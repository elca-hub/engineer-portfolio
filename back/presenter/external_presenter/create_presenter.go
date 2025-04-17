package external_presenter

import (
	"devport/domain/dto"
	"devport/domain/model"
	"devport/usecase/external_service_url"
)

type ExternalServiceUrlPresenter struct{}

func NewCreateExternalServiceUrlPresenter() external_service_url.CreateExternalServiceUrlPresenter {
	return ExternalServiceUrlPresenter{}
}

func (p ExternalServiceUrlPresenter) Output(esu *model.ExternalServiceUrl) external_service_url.CreateExternalServiceUrlOutput {
	esuDto := dto.NewExternalServiceUrlDTO(esu)
	return external_service_url.CreateExternalServiceUrlOutput{
		ExternalServiceUrl: esuDto,
	}
}
