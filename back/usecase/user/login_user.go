package user

import (
	"context"
	"devport/domain/model"
	"devport/domain/repo/nosql"
	"devport/domain/repo/sql"
	"errors"
	"gorm.io/gorm"
	"time"
)

type (
	LoginUserUseCase interface {
		Execute(context.Context, LoginUserInput) (LoginUserOutput, error)
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

	err := i.sqlRepository.WithTransaction(ctx, func(tx context.Context) error {
		email, err := model.NewEmail(input.Email)
		if err != nil {
			return err
		}

		if _, err := i.sqlRepository.FindByEmail(tx, email); err != nil {
			return err
		}

		session, err = i.noSqlRepository.StartSession(email)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return i.presenter.Output(false, ""), nil
		}

		return i.presenter.Output(false, ""), err
	}

	return i.presenter.Output(true, session), nil
}
