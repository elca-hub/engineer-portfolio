package user_presenter

import (
	"devport/domain/model"
	"devport/usecase/user"
)

type LoginUserResponse struct {
	Email string `json:"email"`
}

type LoginUserPresenter struct{}

func NewLoginUserPresenter() user.LoginUserPresenter {
	return LoginUserPresenter{}
}

func (p LoginUserPresenter) Output(email model.Email, token string) user.LoginUserOutput {
	return user.LoginUserOutput{
		Email: email.Email(),
		Token: token,
	}
}
