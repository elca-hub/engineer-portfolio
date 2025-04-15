package user_presenter

import (
	usecase "devport/usecase/user"
)

type UploadUserImagePresenter struct{}

func NewUploadUserImagePresenter() *UploadUserImagePresenter {
	return &UploadUserImagePresenter{}
}

func (p *UploadUserImagePresenter) Output(iconName string, headerIconName string) usecase.UploadUserImageOutput {
	return usecase.UploadUserImageOutput{
		IconName:       iconName,
		HeaderIconName: headerIconName,
	}
}
