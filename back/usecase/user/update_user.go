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
	UpdateUserUseCase interface {
		Execute(context.Context, UpdateUserInput) (UpdateUserOutput, error)
	}

	UpdateUserInput struct {
		User *dto.UserDTO `validate:"required"`
	}

	UpdateUserPresenter interface {
		Output(user *model.User) UpdateUserOutput
	}

	UpdateUserOutput struct {
		User *dto.UserDTO `json:"user"`
	}

	updateUserInterator struct {
		userRepository                db.UserRepository
		externalServiceUrlsRepository db.ExternalServiceUrlsRepository
		bioSentenceStorage            file_storage.BioSentenceStorageRepository
		presenter                     UpdateUserPresenter
		ctxTimeout                    time.Duration
	}
)

func NewUpdateUserInterator(
	userRepository db.UserRepository,
	externalServiceUrlsRepository db.ExternalServiceUrlsRepository,
	bioSentenceStorage file_storage.BioSentenceStorageRepository,
	presenter UpdateUserPresenter,
	t time.Duration,
) UpdateUserUseCase {
	return updateUserInterator{
		userRepository:                userRepository,
		externalServiceUrlsRepository: externalServiceUrlsRepository,
		bioSentenceStorage:            bioSentenceStorage,
		presenter:                     presenter,
		ctxTimeout:                    t,
	}
}

func (i updateUserInterator) Execute(tx context.Context, input UpdateUserInput) (UpdateUserOutput, error) {
	ctx, cancel := context.WithTimeout(tx, i.ctxTimeout)
	defer cancel()

	user, err := i.userRepository.FindById(ctx, input.User.UserId, nil)
	if err != nil {
		return UpdateUserOutput{}, err
	}

	if err := user.UpdateName(input.User.Name); err != nil {
		return UpdateUserOutput{}, err
	}

	jst, _ := time.LoadLocation("Asia/Tokyo")
	birthday, err := time.ParseInLocation("2006-01-02", input.User.Birthday, jst)
	if err != nil {
		return UpdateUserOutput{}, err
	}
	if err := user.UpdateBirthday(birthday); err != nil {
		return UpdateUserOutput{}, err
	}

	if err := user.UpdateOrganizationName(input.User.OrganizationName); err != nil {
		return UpdateUserOutput{}, err
	}

	if err := user.UpdateOccupationName(input.User.OccupationName); err != nil {
		return UpdateUserOutput{}, err
	}

	if err := user.UpdatePlace(input.User.Place); err != nil {
		return UpdateUserOutput{}, err
	}

	if err := i.userRepository.Update(tx, user); err != nil {
		return UpdateUserOutput{}, err
	}

	return i.presenter.Output(user), nil
}
