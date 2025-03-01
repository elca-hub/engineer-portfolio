package user_presenter

import (
	"devport/usecase/user"
)

type LogoutUserPresenter struct{}

func NewLogoutUserPresenter() *LogoutUserPresenter {
	return &LogoutUserPresenter{}
}

func (p *LogoutUserPresenter) Output() user.LogoutUserOutput {
	return user.LogoutUserOutput{}
}
