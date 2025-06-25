package external_service_url

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/db"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type (
	CreateExternalServiceUrlUseCase interface {
		Execute(context.Context, CreateExternalServiceUrlInput) (CreateExternalServiceUrlOutput, error)
	}

	CreateExternalServiceUrlInput struct {
		ServiceType int    `json:"service_type" validate:"min=0,max=5"`
		Url         string `json:"url" validate:"required"`
		UserId      string `validate:"required"`
	}

	CreateExternalServiceUrlOutput struct {
		ExternalServiceUrl *dto.ExternalServiceUrlDTO `json:"external_service_url"`
	}

	CreateExternalServiceUrlPresenter interface {
		Output(esu *model.ExternalServiceUrl) CreateExternalServiceUrlOutput
	}

	createExternalServiceUrlInteractor struct {
		externalServiceUrlsRepository db.ExternalServiceUrlsRepository
		userRepository                db.UserRepository
		presenter                     CreateExternalServiceUrlPresenter
		ctxTimeout                    time.Duration
	}
)

func NewCreateExternalServiceUrlInteractor(
	externalServiceUrlsRepository db.ExternalServiceUrlsRepository,
	userRepository db.UserRepository,
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

	var externalServiceUrl *model.ExternalServiceUrl

	err := i.externalServiceUrlsRepository.WithTransaction(ctx, func(ctx context.Context) error {
		id, err := uuid.NewRandom()

		if err != nil {
			return err
		}

		externalServiceUrl, err = model.NewExternalServiceUrl(id.String(), input.ServiceType, input.Url)

		if err != nil {
			return err
		}

		user, err := i.userRepository.FindById(ctx, input.UserId, nil)

		if err != nil {
			return err
		}

		exists, err := i.externalServiceUrlsRepository.IsExistsByServiceType(ctx, user, input.ServiceType)

		if err != nil {
			return err
		}

		if exists {
			return fmt.Errorf("既に登録されています: %d", input.ServiceType)
		}

		return i.externalServiceUrlsRepository.Create(ctx, user, externalServiceUrl)
	})

	if err != nil {
		return CreateExternalServiceUrlOutput{}, err
	}

	return i.presenter.Output(externalServiceUrl), nil
}
