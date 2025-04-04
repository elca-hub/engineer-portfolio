package user_presenter

import (
	"devport/usecase/user"
)

type isExistsUserPresenter struct{}

func NewIsExistsUserPresenter() user.IsExistsUserPresenter {
	return &isExistsUserPresenter{}
}

func (p isExistsUserPresenter) Output(isExists bool, userId string) user.IsExistsUserOutput {
	return user.IsExistsUserOutput{
		IsExists: isExists,
		UserId:   userId,
	}
}
