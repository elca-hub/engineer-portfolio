package user

import (
	"devport/domain/model"
	"devport/domain/repository/nosql"
	"devport/domain/repository/sql"
)

type (
	LoginUserUseCase interface {
		Execute(LoginUserInput) (LoginUserOutput, error)
	}

	LoginUserInput struct {
		Email string `validate:"required,email"`
	}

	LoginUserPresenter interface {
		Output(isExists bool, token string) LoginUserOutput
	}

	LoginUserOutput struct {
		IsExists bool
		Token    string
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
	email, err := model.NewEmail(input.Email)

	if err != nil {
		return i.presenter.Output(false, ""), err
	}

	if _, err := i.sqlRepository.FindByEmail(email); err != nil {
		if err.Error() == "record not found" {
			return i.presenter.Output(false, ""), nil
		} else {
			return i.presenter.Output(false, ""), err
		}
	}

	session, err := i.noSqlRepository.StartSession(email)

	if err != nil {
		return i.presenter.Output(false, ""), err
	}

	return i.presenter.Output(true, session), nil
}
