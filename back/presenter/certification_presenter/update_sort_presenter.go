package certification_presenter

import (
	"devport/usecase/certification"
)

type UpdateCertificationsSortPresenter struct{}

func NewUpdateCertificationsSortPresenter() certification.UpdateCertificationSortPresenter {
	return UpdateCertificationsSortPresenter{}
}

func (p UpdateCertificationsSortPresenter) Output() certification.UpdateCertificationSortOutput {
	return certification.UpdateCertificationSortOutput{}
}
