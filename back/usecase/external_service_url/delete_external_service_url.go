package external_service_url

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/db"
	"time"
)

type (
	DeleteExternalServiceUrlUseCase interface {
		Execute(context.Context, DeleteExternalServiceUrlInput) (DeleteExternalServiceUrlOutput, error)
	}

	DeleteExternalServiceUrlInput struct {
		UserId               string `validate:"required"`
		ExternalServiceUrlId string `validate:"required"`
	}

	DeleteExternalServiceUrlOutput struct {
		ExternalServiceUrl *dto.ExternalServiceUrlDTO `json:"external_service_url"`
	}

	DeleteExternalServiceUrlPresenter interface {
		Output(esu *model.ExternalServiceUrl) DeleteExternalServiceUrlOutput
	}

	deleteExternalServiceUrlInteractor struct {
		externalServiceUrlsRepository db.ExternalServiceUrlsRepository
		userRepository                db.UserRepository
		presenter                     DeleteExternalServiceUrlPresenter
		ctxTimeout                    time.Duration
	}
)

func NewDeleteExternalServiceUrlInteractor(
	externalServiceUrlsRepository db.ExternalServiceUrlsRepository,
	userRepository db.UserRepository,
	presenter DeleteExternalServiceUrlPresenter,
	t time.Duration,
) DeleteExternalServiceUrlUseCase {
	return deleteExternalServiceUrlInteractor{
		externalServiceUrlsRepository: externalServiceUrlsRepository,
		userRepository:                userRepository,
		presenter:                     presenter,
		ctxTimeout:                    t,
	}
}

func (i deleteExternalServiceUrlInteractor) Execute(ctx context.Context, input DeleteExternalServiceUrlInput) (DeleteExternalServiceUrlOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	var esu *model.ExternalServiceUrl

	err := i.externalServiceUrlsRepository.WithTransaction(ctx, func(ctx context.Context) error {
		user, err := i.userRepository.FindById(ctx, input.UserId, nil)

		if err != nil {
			return err
		}

		esu, err = i.externalServiceUrlsRepository.FindById(ctx, input.ExternalServiceUrlId)
		if err != nil {
			return err
		}

		return i.externalServiceUrlsRepository.Delete(ctx, user, esu)
	})

	if err != nil {
		return DeleteExternalServiceUrlOutput{}, err
	}

	return i.presenter.Output(esu), nil
}
