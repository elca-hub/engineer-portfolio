package certification_presenter

import (
	"devport/domain/dto"
	"devport/domain/model"
	"devport/usecase/certification"
)

type UpdateCertificationPresenter struct{}

func NewUpdateCertificationPresenter() certification.UpdateCertificationPresenter {
	return UpdateCertificationPresenter{}
}

func (p UpdateCertificationPresenter) Output(c *model.Certification) certification.UpdateCertificationOutput {
	certificationDto := dto.NewCertificationDTO(c)
	return certification.UpdateCertificationOutput{
		Certification: certificationDto,
	}
}