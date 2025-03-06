package user

import (
	"devport/domain/model"
	"devport/domain/repository/nosql"
	"devport/domain/repository/sql"
	"errors"
)

type (
	VerificationEmailUseCase interface {
		Execute(VerificationEmailInput) (VerificationEmailOutput, error)
	}

	VerificationEmailInput struct {
		AccessCode int64  `validate:"required" json:"access_code"`
		Email      string `validate:"required,email" json:"email"`
	}

	VerificationEmailPresenter interface {
		Output(token string) VerificationEmailOutput
	}

	VerificationEmailOutput struct {
		Token string
	}

	verificationEmailInterator struct {
		sqlRepository   sql.UserRepository
		noSqlRepository nosql.UserRepository
		presenter       VerificationEmailPresenter
	}
)

func NewVerificationEmailInterator(
	sqlRepository sql.UserRepository,
	noSqlRepository nosql.UserRepository,
	presenter VerificationEmailPresenter,
) VerificationEmailUseCase {
	return verificationEmailInterator{
		sqlRepository:   sqlRepository,
		noSqlRepository: noSqlRepository,
		presenter:       presenter,
	}
}

func (i verificationEmailInterator) Execute(input VerificationEmailInput) (VerificationEmailOutput, error) {
	userEmail, err := model.NewEmail(input.Email)

	if err != nil {
		return i.presenter.Output(""), err
	}

	code, err := i.noSqlRepository.GetConfirmationCode(userEmail)

	if err != nil {
		return i.presenter.Output(""), err
	}

	if code != input.AccessCode {
		return i.presenter.Output(""), errors.New("アクセスコードが違います。再度ログインしてください")
	}

	userModel, err := i.sqlRepository.FindByEmail(userEmail)

	if err != nil {
		return i.presenter.Output(""), err
	}

	userModel.UpdateEmailVerification(model.Confirmed)

	if err := i.sqlRepository.Update(userModel); err != nil {
		return i.presenter.Output(""), err
	}

	token, err := i.noSqlRepository.StartSession(userEmail)

	if err != nil {
		return i.presenter.Output(""), err
	}

	return i.presenter.Output(token), nil
}
