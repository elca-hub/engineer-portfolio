package user_presenter

import (
	"devport/usecase/user"
)

type VerificationEmailPresenter struct{}

func NewVerificationEmailPresenter() *VerificationEmailPresenter {
	return &VerificationEmailPresenter{}
}

func (p *VerificationEmailPresenter) Output(token string) user.VerificationEmailOutput {
	return user.VerificationEmailOutput{
		Token: token,
	}
}
