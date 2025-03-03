package user

import (
	"devport/domain/model"
	"devport/domain/repository/nosql"
	"devport/domain/repository/sql"
	"errors"
)

type (
	VerifyCookieTokenUseCase interface {
		Execute(input VerifyCookieTokenInput) (VerifyCookieTokenOutput, error)
	}

	VerifyCookieTokenInput struct {
		Token string `validate:"required"`
	}

	VerifyCookieTokenPresenter interface {
		Output(email model.Email, Token string) VerifyCookieTokenOutput
	}

	VerifyCookieTokenOutput struct {
		Email string
		Token string
	}

	verifyCookieTokenInterator struct {
		sqlRepository   sql.UserRepository
		noSqlRepository nosql.UserRepository
		presenter       VerifyCookieTokenPresenter
	}
)

func NewVerifyCookieTokenInterator(
	sqlRepository sql.UserRepository,
	noSqlRepository nosql.UserRepository,
	presenter VerifyCookieTokenPresenter,
) VerifyCookieTokenUseCase {
	return verifyCookieTokenInterator{
		sqlRepository:   sqlRepository,
		noSqlRepository: noSqlRepository,
		presenter:       presenter,
	}
}

func (i verifyCookieTokenInterator) Execute(input VerifyCookieTokenInput) (VerifyCookieTokenOutput, error) {
	email, err := i.noSqlRepository.GetSession(input.Token)

	if err != nil {
		return i.presenter.Output(model.Email{}, ""), err
	}

	isExist, err := i.sqlRepository.Exists(email)

	if err != nil {
		return i.presenter.Output(model.Email{}, ""), err
	}

	if !isExist {
		return i.presenter.Output(model.Email{}, ""), errors.New("user not found")
	}

	token, err := i.noSqlRepository.StartSession(email)

	if err != nil {
		return i.presenter.Output(*email, ""), err
	}

	return i.presenter.Output(*email, token), nil
}
