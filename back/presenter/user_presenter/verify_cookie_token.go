package user_presenter

import (
	"devport/domain/model"
	"devport/usecase/user"
)

type VerifyCookieTokenResponse struct {
	Email string `json:"email"`
}

type VerifyCookieTokenPresenter struct{}

func NewVerifyCookieTokenPresenter() user.VerifyCookieTokenPresenter {
	return VerifyCookieTokenPresenter{}
}

func (p VerifyCookieTokenPresenter) Output(userModel *model.User) user.VerifyCookieTokenOutput {
	return user.VerifyCookieTokenOutput{
		User: userModel,
	}
}
