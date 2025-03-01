package user

import (
	"devport/domain/repository/nosql"
	"devport/domain/repository/sql"
)

type (
	LogoutUserUseCase interface {
		Execute(input LogoutUserInput) (LogoutUserOutput, error)
	}

	LogoutUserInput struct {
		Token string `validate:"required"`
	}

	LogoutUserPresenter interface {
		Output() LogoutUserOutput
	}

	LogoutUserOutput struct{}

	logoutUserInterator struct {
		sqlRepository   sql.UserRepository
		noSqlRepository nosql.UserRepository
		presenter       LogoutUserPresenter
	}
)

func NewLogoutUserInterator(
	noSqlRepository nosql.UserRepository,
	presenter LogoutUserPresenter,
) LogoutUserUseCase {
	return logoutUserInterator{
		noSqlRepository: noSqlRepository,
		presenter:       presenter,
	}
}

func (i logoutUserInterator) Execute(input LogoutUserInput) (LogoutUserOutput, error) {
	err := i.noSqlRepository.DeleteSession(input.Token)

	if err != nil {
		return i.presenter.Output(), err
	}

	return i.presenter.Output(), nil
}
