package user_presenter

import (
	"devport/usecase/user"
)

type LoginUserResponse struct {
	Email string `json:"email"`
}

type LoginUserPresenter struct{}

func NewLoginUserPresenter() user.LoginUserPresenter {
	return LoginUserPresenter{}
}

func (p LoginUserPresenter) Output(token string) user.LoginUserOutput {
	return user.LoginUserOutput{
		Token: token,
	}
}
