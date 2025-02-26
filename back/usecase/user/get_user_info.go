package user

import (
	"devport/domain/model"
	"devport/domain/repository/nosql"
	"devport/domain/repository/sql"
	"errors"
)

type (
	GetUserInfoUseCase interface {
		Execute(GetUserInfoInput) (GetUserInfoOutput, error)
	}

	GetUserInfoInput struct {
		Token string `validate:"required"`
	}

	GetUserInfoPresenter interface {
		Output(user model.User, token string) GetUserInfoOutput
	}

	GetUserInfoOutput struct {
		Email string
		Token string
	}

	getUserInfoInterator struct {
		sqlRepository   sql.UserRepository
		noSqlRepository nosql.UserRepository
		presenter       GetUserInfoPresenter
	}
)

func NewGetUserInfoInterator(
	sqlRepository sql.UserRepository,
	noSqlRepository nosql.UserRepository,
	presenter GetUserInfoPresenter,
) GetUserInfoUseCase {
	return getUserInfoInterator{
		sqlRepository:   sqlRepository,
		noSqlRepository: noSqlRepository,
		presenter:       presenter,
	}
}

func (i getUserInfoInterator) Execute(input GetUserInfoInput) (GetUserInfoOutput, error) {
	userEmail, err := i.noSqlRepository.GetSession(input.Token)

	if err != nil {
		return i.presenter.Output(model.User{}, ""), errors.New("invalid token_auth")
	}

	userModel, err := i.sqlRepository.FindByEmail(userEmail)

	if err != nil {
		return i.presenter.Output(model.User{}, ""), err
	}

	session, err := i.noSqlRepository.StartSession(userEmail)
	if err != nil {
		return i.presenter.Output(model.User{}, ""), err
	}

	return i.presenter.Output(*userModel, session), nil
}
