package certification_presenter

import (
	"devport/domain/dto"
	"devport/domain/model"
	"devport/usecase/certification"
)

type FindByUserCertificationPresenter struct{}

func NewFindByUserCertificationPresenter() certification.FindByUserCertificationPresenter {
	return FindByUserCertificationPresenter{}
}

func (p FindByUserCertificationPresenter) Output(certifications []*model.Certification) certification.FindByUserCertificationOutput {
	certificationDtos := make([]*dto.CertificationDTO, len(certifications))
	for i, c := range certifications {
		certificationDtos[i] = dto.NewCertificationDTO(c)
	}
	return certification.FindByUserCertificationOutput{
		Certifications: certificationDtos,
	}
}