package certification_presenter

import (
	"devport/domain/dto"
	"devport/domain/model"
	"devport/usecase/certification"
)

type CertificationPresenter struct{}

func NewCreateCertificationPresenter() certification.CreateCertificationPresenter {
	return CertificationPresenter{}
}

func (p CertificationPresenter) Output(c *model.Certification) certification.CreateCertificationOutput {
	certificationDto := dto.NewCertificationDTO(c)
	return certification.CreateCertificationOutput{
		Certification: certificationDto,
	}
}