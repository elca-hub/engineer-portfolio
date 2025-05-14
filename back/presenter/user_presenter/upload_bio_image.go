package user_presenter

import (
	usecase "devport/usecase/user"
)

type UploadBioImagePresenter struct{}

func NewUploadBioImagePresenter() *UploadBioImagePresenter {
	return &UploadBioImagePresenter{}
}

func (p *UploadBioImagePresenter) Output(imageUrl string) usecase.UploadBioImageOutput {
	return usecase.UploadBioImageOutput{
		ImageUrl: imageUrl,
	}
}
