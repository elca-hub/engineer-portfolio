package user

import (
	"devport/domain/repo/nosql"
	"devport/domain/repo/sql"
)

type (
	LogoutUserUseCase interface {
		Execute(input LogoutUserInput) (LogoutUserOutput, error)
	}

	LogoutUserInput struct {
		Token string `validate:"required" json:"token"`
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
