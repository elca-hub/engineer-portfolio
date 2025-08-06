package certification_presenter

import (
	"devport/usecase/certification"
)

type DeleteCertificationPresenter struct{}

func NewDeleteCertificationPresenter() certification.DeleteCertificationPresenter {
	return DeleteCertificationPresenter{}
}

func (p DeleteCertificationPresenter) Output() certification.DeleteCertificationOutput {
	return certification.DeleteCertificationOutput{}
}
