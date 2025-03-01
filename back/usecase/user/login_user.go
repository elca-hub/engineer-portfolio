package user

import (
	"devport/domain/model"
	"devport/domain/repository/nosql"
	"devport/domain/repository/sql"
	"devport/infra/security"
	"errors"
)

type (
	LoginUserUseCase interface {
		Execute(LoginUserInput) (LoginUserOutput, error)
	}

	LoginUserInput struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=8,max=32"`
	}

	LoginUserPresenter interface {
		Output(user model.User, token string) LoginUserOutput
	}

	LoginUserOutput struct {
		Email string
		Token string
	}

	loginUserInterator struct {
		sqlRepository   sql.UserRepository
		noSqlRepository nosql.UserRepository
		presenter       LoginUserPresenter
	}
)

func NewLoginUserInterator(
	sqlRepository sql.UserRepository,
	noSqlRepository nosql.UserRepository,
	presenter LoginUserPresenter,
) LoginUserUseCase {
	return loginUserInterator{
		sqlRepository:   sqlRepository,
		noSqlRepository: noSqlRepository,
		presenter:       presenter,
	}
}

func (i loginUserInterator) Execute(input LoginUserInput) (LoginUserOutput, error) {
	inputEmail, err := model.NewEmail(input.Email)

	if err != nil {
		return i.presenter.Output(model.User{}, ""), err
	}

	userModel, err := i.sqlRepository.FindByEmail(inputEmail)

	if err != nil {
		return i.presenter.Output(model.User{}, ""), err
	}

	rawPassword, err := model.NewRawPassword(input.Password)

	if err != nil {
		return i.presenter.Output(*userModel, ""), err
	}

	if !security.CheckPasswordHash(rawPassword, userModel.Password()) {
		return i.presenter.Output(*userModel, ""), errors.New("invalid password")
	}

	session, err := i.noSqlRepository.StartSession(inputEmail)

	if err != nil {
		return i.presenter.Output(*userModel, ""), err
	}

	return i.presenter.Output(*userModel, session), nil
}
