package user

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/sql"
	"time"
)

type (
	GetUserInfoUseCase interface {
		Execute(context.Context, GetUserInfoInput) (GetUserInfoOutput, error)
	}

	GetUserInfoInput struct {
		Token string `validate:"required"`
		Email string `validate:"required"`
	}

	GetUserInfoPresenter interface {
		Output(user model.User) GetUserInfoOutput
	}

	GetUserInfoOutput struct {
		User *dto.UserDTO `json:"user"`
	}

	getUserInfoInterator struct {
		sqlRepository sql.UserRepository
		presenter     GetUserInfoPresenter
		ctxTimeout    time.Duration
	}
)

func NewGetUserInfoInterator(
	sqlRepository sql.UserRepository,
	presenter GetUserInfoPresenter,
	t time.Duration,
) GetUserInfoUseCase {
	return getUserInfoInterator{
		sqlRepository: sqlRepository,
		presenter:     presenter,
		ctxTimeout:    t,
	}
}

func (i getUserInfoInterator) Execute(tx context.Context, input GetUserInfoInput) (GetUserInfoOutput, error) {
	ctx, cancel := context.WithTimeout(tx, i.ctxTimeout)
	defer cancel()

	var (
		userModel *model.User
	)

	err := i.sqlRepository.WithTransaction(ctx, func(tx context.Context) error {
		email, err := model.NewEmail(input.Email)

		if err != nil {
			return err
		}

		userModel, err = i.sqlRepository.FindByEmail(tx, email)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return GetUserInfoOutput{}, err
	}

	return i.presenter.Output(*userModel), nil
}
