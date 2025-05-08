package user_presenter

import (
	usecase "devport/usecase/user"
)

type UploadBioImagePresenter struct{}

func NewUploadBioImagePresenter() *UploadBioImagePresenter {
	return &UploadBioImagePresenter{}
}

func (p *UploadBioImagePresenter) Output(imageId string) usecase.UploadBioImageOutput {
	return usecase.UploadBioImageOutput{
		ImageId: imageId,
	}
}
