package user

import (
	"context"
	"devport/domain/model"
	"devport/domain/repository/nosql"
	"devport/domain/repository/sql"
	"errors"
	"gorm.io/gorm"
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
		Output(email *model.Email) VerifyCookieTokenOutput
	}

	VerifyCookieTokenOutput struct {
		Email string
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

func (i verifyCookieTokenInterator) Execute(ctx context.Context, input VerifyCookieTokenInput) (VerifyCookieTokenOutput, error) {
	var (
		email *model.Email
	)
	err := i.sqlRepository.WithTransaction(ctx, func(tx *gorm.DB) error {
		email, err := i.noSqlRepository.GetSession(input.Token)

		if err != nil {
			return err
		}

		isExist, err := i.sqlRepository.Exists(tx, email)

		if err != nil {
			return err
		}

		if !isExist {
			return errors.New("ユーザが存在しません")
		}

		return nil
	})

	if err != nil {
		return VerifyCookieTokenOutput{}, err
	}

	return i.presenter.Output(email), nil
}
