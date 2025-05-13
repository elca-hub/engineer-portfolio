package user

import (
	"context"
	"devport/domain/model"
	"devport/domain/repo/nosql"
	"devport/domain/repo/sql"
	"errors"
	"time"
)

type (
	VerifyCookieTokenUseCase interface {
		Execute(context.Context, VerifyCookieTokenInput) (VerifyCookieTokenOutput, error)
	}

	VerifyCookieTokenInput struct {
		Token string `validate:"required"`
	}

	VerifyCookieTokenPresenter interface {
		Output(email *model.User) VerifyCookieTokenOutput
	}

	VerifyCookieTokenOutput struct {
		User *model.User
	}

	verifyCookieTokenInterator struct {
		sqlRepository   sql.UserRepository
		noSqlRepository nosql.UserRepository
		presenter       VerifyCookieTokenPresenter
		ctxTimeout      time.Duration
	}
)

func NewVerifyCookieTokenInterator(
	sqlRepository sql.UserRepository,
	noSqlRepository nosql.UserRepository,
	presenter VerifyCookieTokenPresenter,
	t time.Duration,
) VerifyCookieTokenUseCase {
	return verifyCookieTokenInterator{
		sqlRepository:   sqlRepository,
		noSqlRepository: noSqlRepository,
		presenter:       presenter,
		ctxTimeout:      t,
	}
}

func (i verifyCookieTokenInterator) Execute(cx context.Context, input VerifyCookieTokenInput) (VerifyCookieTokenOutput, error) {
	var (
		user *model.User
	)

	ctx, cancel := context.WithTimeout(cx, i.ctxTimeout)

	defer cancel()

	err := i.sqlRepository.WithTransaction(ctx, func(tx context.Context) error {

		isExist, err := i.noSqlRepository.IsExistSession(input.Token)

		if err != nil {
			return err
		}

		if !isExist {
			return errors.New("セッションが無効化されました。再度ログインしてください")
		}

		email, err := i.noSqlRepository.GetSession(input.Token)

		if err != nil {
			return err
		}

		user, err = i.sqlRepository.FindByEmail(tx, email)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return VerifyCookieTokenOutput{}, err
	}

	return i.presenter.Output(user), nil
}
