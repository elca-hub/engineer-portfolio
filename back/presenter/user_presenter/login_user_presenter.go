package user_presenter

import (
	"devport/usecase/user"
)

type LoginUserPresenter struct{}

func NewLoginUserPresenter() user.LoginUserPresenter {
	return LoginUserPresenter{}
}

func (p LoginUserPresenter) Output(token string, userId string) user.LoginUserOutput {
	return user.LoginUserOutput{
		Token:  token,
		UserId: userId,
	}
}
