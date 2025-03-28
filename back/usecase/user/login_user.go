package user

import (
	"context"
	"devport/domain/model"
	"devport/domain/repo/nosql"
	"devport/domain/repo/sql"
	"time"
)

type (
	LoginUserUseCase interface {
		Execute(context.Context, LoginUserInput) (LoginUserOutput, error)
	}

	LoginUserInput struct {
		Email string `validate:"required,email" json:"email"`
	}

	LoginUserPresenter interface {
		Output(token string) LoginUserOutput
	}

	LoginUserOutput struct {
		Token string `json:"token"`
	}

	loginUserInterator struct {
		sqlRepository   sql.UserRepository
		noSqlRepository nosql.UserRepository
		presenter       LoginUserPresenter
		ctxTimeout      time.Duration
	}
)

func NewLoginUserInterator(
	sqlRepository sql.UserRepository,
	noSqlRepository nosql.UserRepository,
	presenter LoginUserPresenter,
	t time.Duration,
) LoginUserUseCase {
	return loginUserInterator{
		sqlRepository:   sqlRepository,
		noSqlRepository: noSqlRepository,
		presenter:       presenter,
		ctxTimeout:      t,
	}
}

func (i loginUserInterator) Execute(tx context.Context, input LoginUserInput) (LoginUserOutput, error) {
	ctx, cancel := context.WithTimeout(tx, i.ctxTimeout)
	defer cancel()

	var (
		session string
	)

	email, err := model.NewEmail(input.Email)
	if err != nil {
		return i.presenter.Output(""), err
	}

	if _, err := i.sqlRepository.FindByEmail(ctx, email); err != nil {
		return i.presenter.Output(""), err
	}

	session, err = i.noSqlRepository.StartSession(email)

	if err != nil {
		return i.presenter.Output(""), err
	}

	return i.presenter.Output(session), nil
}
