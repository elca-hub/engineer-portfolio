package external_service_url

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/db"
	"errors"
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
		externalServiceUrlsRepository db.ExternalServiceUrlsRepository
		userRepository                db.UserRepository
		presenter                     UpdateExternalServiceUrlPresenter
		ctxTimeout                    time.Duration
	}
)

func NewUpdateExternalServiceUrlInteractor(
	externalServiceUrlsRepository db.ExternalServiceUrlsRepository,
	userRepository db.UserRepository,
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

	if isExists, err := i.userRepository.ExistsById(ctx, input.UserId); err != nil || !isExists {
		if err != nil {
			return UpdateExternalServiceUrlOutput{}, err
		}
		if !isExists {
			return UpdateExternalServiceUrlOutput{}, errors.New("ユーザーが存在しません")
		}
	}

	var externalServiceUrl *model.ExternalServiceUrl

	err := i.externalServiceUrlsRepository.WithTransaction(ctx, func(ctx context.Context) error {
		var errTmp error
		externalServiceUrl, errTmp = i.externalServiceUrlsRepository.FindById(ctx, input.Id)

		if errTmp != nil {
			return errTmp
		}

		if err := externalServiceUrl.UpdateUrl(input.Url); err != nil {
			return err
		}

		return i.externalServiceUrlsRepository.Update(ctx, input.UserId, externalServiceUrl)
	})

	if err != nil {
		return UpdateExternalServiceUrlOutput{}, err
	}

	return i.presenter.Output(externalServiceUrl), nil
}
