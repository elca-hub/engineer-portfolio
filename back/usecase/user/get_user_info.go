package user

import (
	"devport/domain/model"
	"devport/domain/repository/nosql"
	"devport/domain/repository/sql"
)

type (
	GetUserInfoUseCase interface {
		Execute(GetUserInfoInput) (GetUserInfoOutput, error)
	}

	GetUserInfoInput struct {
		Token string `validate:"required"`
		Email string `validate:"required"`
	}

	GetUserInfoPresenter interface {
		Output(user model.User, token string) GetUserInfoOutput
	}

	GetUserInfoOutput struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Token string `json:"-"`
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
	email, err := model.NewEmail(input.Email)

	if err != nil {
		return i.presenter.Output(model.User{}, ""), err
	}

	userModel, err := i.sqlRepository.FindByEmail(email)

	if err != nil {
		return i.presenter.Output(model.User{}, ""), err
	}

	session, err := i.noSqlRepository.StartSession(email)
	if err != nil {
		return i.presenter.Output(model.User{}, ""), err
	}

	return i.presenter.Output(*userModel, session), nil
}
