package user_presenter

import (
	"devport/domain/dto"
	"devport/domain/model"
	usecase "devport/usecase/user"
)

type UpdateUserPresenter struct{}

func NewUpdateUserPresenter() *UpdateUserPresenter {
	return &UpdateUserPresenter{}
}

func (p *UpdateUserPresenter) Output(user *model.User) usecase.UpdateUserOutput {
	userDto := dto.NewUserDTO(user)

	return usecase.UpdateUserOutput{
		User: userDto,
	}
}
