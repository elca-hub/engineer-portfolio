package user_presenter

import (
	usermodel "devport/domain/model"
	"devport/usecase/user"
)

type GetUserInfoPresenter struct{}

func NewGetUserInfoPresenter() *GetUserInfoPresenter {
	return &GetUserInfoPresenter{}
}

func (p *GetUserInfoPresenter) Output(model usermodel.User, token string) user.GetUserInfoOutput {
	emailModel := model.Email()
	var email string
	if emailModel == nil {
		email = ""
	} else {
		email = emailModel.Email()
	}
	
	return user.GetUserInfoOutput{
		Email: email,
		Name:  model.Name(),
		Age:   model.Age(),
		Token: token,
	}
}
