package user

import (
	"context"
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
