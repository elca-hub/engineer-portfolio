package user_presenter

import (
	usecase "devport/usecase/user"
)

type UploadBioPresenter struct{}

func NewUploadBioPresenter() *UploadBioPresenter {
	return &UploadBioPresenter{}
}

func (p *UploadBioPresenter) Output(bio string) usecase.UploadBioOutput {
	return usecase.UploadBioOutput{
		Bio: bio,
	}
}
