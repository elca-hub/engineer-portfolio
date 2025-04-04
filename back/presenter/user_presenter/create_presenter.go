package user_presenter

import (
	"devport/usecase/user"
)

type CreateUserPresenter struct{}

func NewCreateUserPresenter() user.CreateUserPresenter {
	return CreateUserPresenter{}
}

func (p CreateUserPresenter) Output(email string) user.CreateUserOutput {
	return user.CreateUserOutput{
		Email: email,
	}
}
