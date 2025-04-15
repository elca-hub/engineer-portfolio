package user

import (
	"context"
	"devport/domain/model"
	"devport/domain/repo/sql"
	"time"
)

type (
	CreateExternalServiceUrlUseCase interface {
		Execute(context.Context, CreateExternalServiceUrlInput) (CreateExternalServiceUrlOutput, error)
	}

	CreateExternalServiceUrlInput struct {
		ServiceId uint   `json:"service_id" validate:"required"`
		Url       string `json:"url" validate:"required"`
		UserId    string `json:"user_id" validate:"required"`
	}

	CreateExternalServiceUrlOutput struct{}

	CreateExternalServiceUrlPresenter interface {
		Output(email string) CreateExternalServiceUrlOutput
	}
	createExternalServiceUrlInteractor struct {
		externalServiceUrlsRepository sql.ExternalServiceUrlsRepository
		userRepository                sql.UserRepository
		presenter                     CreateExternalServiceUrlPresenter
		ctxTimeout                    time.Duration
	}
)

func NewCreateExternalServiceUrlInteractor(
	externalServiceUrlsRepository sql.ExternalServiceUrlsRepository,
	userRepository sql.UserRepository,
	presenter CreateExternalServiceUrlPresenter,
	t time.Duration,
) CreateExternalServiceUrlUseCase {
	return createExternalServiceUrlInteractor{
		externalServiceUrlsRepository: externalServiceUrlsRepository,
		userRepository:                userRepository,
		presenter:                     presenter,
		ctxTimeout:                    t,
	}
}

func (i createExternalServiceUrlInteractor) Execute(ctx context.Context, input CreateExternalServiceUrlInput) (CreateExternalServiceUrlOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	err := i.externalServiceUrlsRepository.WithTransaction(ctx, func(ctx context.Context) error {
		externalServiceUrl, err := model.NewExternalServiceUrl(input.Url, input.ServiceId)

		if err != nil {
			return err
		}

		user, err := i.userRepository.FindById(ctx, input.UserId)

		if err != nil {
			return err
		}

		return i.externalServiceUrlsRepository.Create(ctx, user, externalServiceUrl)
	})

	if err != nil {
		return CreateExternalServiceUrlOutput{}, err
	}

	return CreateExternalServiceUrlOutput{}, nil
}
