package user

import (
	"context"
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
		Output(user model.User) FetchUserInfoOutput
	}

	FetchUserInfoOutput struct {
		Email      string `json:"email"`
		Name       string `json:"name"`
		IconPath   string `json:"icon_path"`
		HeaderPath string `json:"header_path"`
		BioPath    string `json:"bio_path"`
		Skills     []struct {
			Name      string `json:"name"`
			Status    string `json:"status"`
			When      string `json:"when"`
			SortIndex int    `json:"sort_index"`
		} `json:"skills"`
		ExternalServiceUrl []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"external_service_url"`
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

	return i.presenter.Output(*user), nil
}
