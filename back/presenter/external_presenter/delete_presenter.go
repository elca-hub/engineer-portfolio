package external_presenter

import (
	"devport/domain/dto"
	"devport/domain/model"
	"devport/usecase/external_service_url"
)

type DeleteExternalServiceUrlPresenter struct{}

func NewDeleteExternalServiceUrlPresenter() external_service_url.DeleteExternalServiceUrlPresenter {
	return DeleteExternalServiceUrlPresenter{}
}

func (p DeleteExternalServiceUrlPresenter) Output(esu *model.ExternalServiceUrl) external_service_url.DeleteExternalServiceUrlOutput {
	return external_service_url.DeleteExternalServiceUrlOutput{
		ExternalServiceUrl: dto.NewExternalServiceUrlDTO(esu),
	}
}
