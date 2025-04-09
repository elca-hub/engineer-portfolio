package user

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/sql"
	"time"
)

type (
	FetchUserInfoUseCase interface {
		Execute(context.Context, FetchUserInfoInput) (FetchUserInfoOutput, error)
	}

	FetchUserInfoInput struct {
		UserId string `json:"user_id"`
	}

	FetchUserInfoPresenter interface {
		Output(user *model.User) FetchUserInfoOutput
	}

	FetchUserInfoOutput struct {
		User *dto.UserDTO `json:"user"`
	}

	fetchUserInfoInterator struct {
		sqlRepository sql.UserRepository
		presenter     FetchUserInfoPresenter
		ctxTimeout    time.Duration
	}
)

func NewFetchUserInfoInterator(
	sqlRepository sql.UserRepository,
	presenter FetchUserInfoPresenter,
	t time.Duration,
) FetchUserInfoUseCase {
	return fetchUserInfoInterator{
		sqlRepository: sqlRepository,
		presenter:     presenter,
		ctxTimeout:    t,
	}
}

func (i fetchUserInfoInterator) Execute(tx context.Context, input FetchUserInfoInput) (FetchUserInfoOutput, error) {
	ctx, cancel := context.WithTimeout(tx, i.ctxTimeout)
	defer cancel()

	user, err := i.sqlRepository.FindById(ctx, input.UserId)

	if err != nil {
		return FetchUserInfoOutput{}, err
	}

	return i.presenter.Output(user), nil
}
