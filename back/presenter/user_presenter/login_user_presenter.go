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

func (p LoginUserPresenter) Output(account model.User, token string) user.LoginUserOutput {
	return user.LoginUserOutput{
		Email: account.Email().Email(),
		Token: token,
	}
}
