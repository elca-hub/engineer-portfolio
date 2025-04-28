package user_presenter

import (
	"devport/domain/dto"
	usermodel "devport/domain/model"
	"devport/usecase/user"
	usecase "devport/usecase/user"
)

type GetUserInfoPresenter struct{}

func NewGetUserInfoPresenter() *GetUserInfoPresenter {
	return &GetUserInfoPresenter{}
}

func (p *GetUserInfoPresenter) Output(user *usermodel.User) user.GetUserInfoOutput {
	userDto := dto.NewUserDTO(user)

	return usecase.GetUserInfoOutput{
		User: userDto,
	}
}
