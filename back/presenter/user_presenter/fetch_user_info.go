package user_presenter

import (
	"devport/domain/dto"
	usermodel "devport/domain/model"
	"devport/usecase/user"
	usecase "devport/usecase/user"
)

type FetchUserInfoPresenter struct{}

func NewFetchUserInfoPresenter() *FetchUserInfoPresenter {
	return &FetchUserInfoPresenter{}
}

func (p *FetchUserInfoPresenter) Output(user *usermodel.User) user.FetchUserInfoOutput {
	userDto := dto.NewUserDTO(user)

	return usecase.FetchUserInfoOutput{
		User: userDto,
	}
}
