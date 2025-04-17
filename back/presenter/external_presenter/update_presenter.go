package external_presenter

import (
	"devport/domain/dto"
	"devport/domain/model"
	"devport/usecase/external_service_url"
)

type UpdateExternalServiceUrlPresenter struct{}

func NewUpdateExternalServiceUrlPresenter() external_service_url.UpdateExternalServiceUrlPresenter {
	return UpdateExternalServiceUrlPresenter{}
}

func (p UpdateExternalServiceUrlPresenter) Output(esu *model.ExternalServiceUrl) external_service_url.UpdateExternalServiceUrlOutput {
	esuDto := dto.NewExternalServiceUrlDTO(esu)
	return external_service_url.UpdateExternalServiceUrlOutput{
		ExternalServiceUrl: esuDto,
	}
}
