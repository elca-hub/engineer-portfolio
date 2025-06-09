package user

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/db"
	"devport/domain/repo/file_storage"
	"time"
)

type (
	GetUserInfoUseCase interface {
		Execute(context.Context, GetUserInfoInput) (GetUserInfoOutput, error)
	}

	GetUserInfoInput struct {
		UserId string `validate:"required"`
	}

	GetUserInfoPresenter interface {
		Output(user *model.User) GetUserInfoOutput
	}

	GetUserInfoOutput struct {
		User *dto.UserDTO `json:"user"`
	}

	getUserInfoInterator struct {
		userRepository     db.UserRepository
		bioSentenceStorage file_storage.BioSentenceStorageRepository
		presenter          GetUserInfoPresenter
		ctxTimeout         time.Duration
	}
)

func NewGetUserInfoInterator(
	sqlRepository db.UserRepository,
	bioSentenceStorage file_storage.BioSentenceStorageRepository,
	presenter GetUserInfoPresenter,
	t time.Duration,
) GetUserInfoUseCase {
	return getUserInfoInterator{
		userRepository:     sqlRepository,
		bioSentenceStorage: bioSentenceStorage,
		presenter:          presenter,
		ctxTimeout:         t,
	}
}

func (i getUserInfoInterator) Execute(tx context.Context, input GetUserInfoInput) (GetUserInfoOutput, error) {
	bioString, err := i.bioSentenceStorage.FindByUserId(tx, input.UserId)

	if err != nil {
		return GetUserInfoOutput{}, err
	}

	bio, err := model.NewBio(input.UserId, bioString)

	if err != nil {
		return GetUserInfoOutput{}, err
	}

	user, err := i.userRepository.FindById(tx, input.UserId, bio)

	if err != nil {
		return GetUserInfoOutput{}, err
	}

	return i.presenter.Output(user), nil
}
