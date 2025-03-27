package user_presenter

import (
	"devport/usecase/user"
)

type isExistsUserPresenter struct{}

func NewIsExistsUserPresenter() user.IsExistsUserPresenter {
	return &isExistsUserPresenter{}
}

func (p isExistsUserPresenter) Output(isExists bool) user.IsExistsUserOutput {
	return user.IsExistsUserOutput{
		IsExists: isExists,
	}
}
