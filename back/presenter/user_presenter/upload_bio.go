package user_presenter

import (
	usecase "devport/usecase/user"
)

type UploadBioPresenter struct{}

func NewUploadBioPresenter() *UploadBioPresenter {
	return &UploadBioPresenter{}
}

func (p *UploadBioPresenter) Output(bioName string) usecase.UploadBioOutput {
	return usecase.UploadBioOutput{
		BioName: bioName,
	}
}
