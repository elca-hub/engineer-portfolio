package external_service_url

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/sql"
	"time"
)

type (
	UpdateExternalServiceUrlUseCase interface {
		Execute(context.Context, UpdateExternalServiceUrlInput) (UpdateExternalServiceUrlOutput, error)
	}

	UpdateExternalServiceUrlInput struct {
		Url    string `json:"url" validate:"required"`
		Id     string `validate:"required"`
		UserId string `validate:"required"`
	}

	UpdateExternalServiceUrlOutput struct {
		ExternalServiceUrl *dto.ExternalServiceUrlDTO `json:"external_service_url"`
	}

	UpdateExternalServiceUrlPresenter interface {
		Output(esu *model.ExternalServiceUrl) UpdateExternalServiceUrlOutput
	}

	updateExternalServiceUrlInteractor struct {
		externalServiceUrlsRepository sql.ExternalServiceUrlsRepository
		userRepository                sql.UserRepository
		presenter                     UpdateExternalServiceUrlPresenter
		ctxTimeout                    time.Duration
	}
)

func NewUpdateExternalServiceUrlInteractor(
	externalServiceUrlsRepository sql.ExternalServiceUrlsRepository,
	userRepository sql.UserRepository,
	presenter UpdateExternalServiceUrlPresenter,
	t time.Duration,
) UpdateExternalServiceUrlUseCase {
	return updateExternalServiceUrlInteractor{
		externalServiceUrlsRepository: externalServiceUrlsRepository,
		userRepository:                userRepository,
		presenter:                     presenter,
		ctxTimeout:                    t,
	}
}

func (i updateExternalServiceUrlInteractor) Execute(ctx context.Context, input UpdateExternalServiceUrlInput) (UpdateExternalServiceUrlOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	var externalServiceUrl *model.ExternalServiceUrl

	err := i.externalServiceUrlsRepository.WithTransaction(ctx, func(ctx context.Context) error {
		var errTmp error
		externalServiceUrl, errTmp = i.externalServiceUrlsRepository.FindById(ctx, input.Id)

		if errTmp != nil {
			return errTmp
		}

		user, err := i.userRepository.FindById(ctx, input.UserId)

		if err != nil {
			return err
		}

		if err := externalServiceUrl.UpdateUrl(input.Url); err != nil {
			return err
		}

		return i.externalServiceUrlsRepository.Update(ctx, user, externalServiceUrl)
	})

	if err != nil {
		return UpdateExternalServiceUrlOutput{}, err
	}

	return i.presenter.Output(externalServiceUrl), nil
}
