package external_service_url

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/sql"
	"time"
)

type (
	FindByUserExternalServiceUrlUseCase interface {
		Execute(context.Context, FindByUserExternalServiceUrlInput) (FindByUserExternalServiceUrlOutput, error)
	}

	FindByUserExternalServiceUrlInput struct {
		UserId string `validate:"required"`
	}

	FindByUserExternalServiceUrlPresenter interface {
		Output(esus []*model.ExternalServiceUrl) FindByUserExternalServiceUrlOutput
	}

	FindByUserExternalServiceUrlOutput struct {
		ExternalServiceUrls []*dto.ExternalServiceUrlDTO `json:"external_service_urls"`
	}

	findByUserExternalServiceUrlInterator struct {
		userRepo   sql.UserRepository
		esuRepo    sql.ExternalServiceUrlsRepository
		presenter  FindByUserExternalServiceUrlPresenter
		ctxTimeout time.Duration
	}
)

func NewFindByUserExternalServiceUrlInterator(
	sqlRepository sql.UserRepository,
	esuRepository sql.ExternalServiceUrlsRepository,
	presenter FindByUserExternalServiceUrlPresenter,
	t time.Duration,
) FindByUserExternalServiceUrlUseCase {
	return findByUserExternalServiceUrlInterator{
		userRepo:   sqlRepository,
		esuRepo:    esuRepository,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (i findByUserExternalServiceUrlInterator) Execute(tx context.Context, input FindByUserExternalServiceUrlInput) (FindByUserExternalServiceUrlOutput, error) {
	ctx, cancel := context.WithTimeout(tx, i.ctxTimeout)
	defer cancel()

	userModel, err := i.userRepo.FindById(ctx, input.UserId)

	if err != nil {
		return FindByUserExternalServiceUrlOutput{}, err
	}

	externalServiceUrls, err := i.esuRepo.FindByUserId(ctx, userModel)

	if err != nil {
		return FindByUserExternalServiceUrlOutput{}, err
	}

	return i.presenter.Output(externalServiceUrls), nil
}
