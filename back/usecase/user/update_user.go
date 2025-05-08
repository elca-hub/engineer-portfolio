package user

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/sql"
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
		userRepository                sql.UserRepository
		externalServiceUrlsRepository sql.ExternalServiceUrlsRepository
		presenter                     UpdateUserPresenter
		ctxTimeout                    time.Duration
	}
)

func NewUpdateUserInterator(
	userRepository sql.UserRepository,
	externalServiceUrlsRepository sql.ExternalServiceUrlsRepository,
	presenter UpdateUserPresenter,
	t time.Duration,
) UpdateUserUseCase {
	return updateUserInterator{
		userRepository:                userRepository,
		externalServiceUrlsRepository: externalServiceUrlsRepository,
		presenter:                     presenter,
		ctxTimeout:                    t,
	}
}

func (i updateUserInterator) Execute(tx context.Context, input UpdateUserInput) (UpdateUserOutput, error) {
	ctx, cancel := context.WithTimeout(tx, i.ctxTimeout)
	defer cancel()

	user, err := i.userRepository.FindById(ctx, input.User.UserId)
	if err != nil {
		return UpdateUserOutput{}, err
	}

	if input.User.Name != "" {
		if err := user.UpdateName(input.User.Name); err != nil {
			return UpdateUserOutput{}, err
		}
	}

	if input.User.Birthday != "" {
		jst, _ := time.LoadLocation("Asia/Tokyo")
		birthday, err := time.ParseInLocation("2006-01-02", input.User.Birthday, jst)
		if err != nil {
			return UpdateUserOutput{}, err
		}
		if err := user.UpdateBirthday(birthday); err != nil {
			return UpdateUserOutput{}, err
		}
	}

	if input.User.IconName != "" {
		user.UpdateIconName(input.User.IconName)
	}

	if input.User.HeaderIconName != "" {
		user.UpdateHeaderIconName(input.User.HeaderIconName)
	}

	if input.User.OrganizationName != "" {
		if err := user.UpdateOrganizationName(input.User.OrganizationName); err != nil {
			return UpdateUserOutput{}, err
		}
	}

	if input.User.OccupationName != "" {
		if err := user.UpdateOccupationName(input.User.OccupationName); err != nil {
			return UpdateUserOutput{}, err
		}
	}

	if input.User.Place != "" {
		if err := user.UpdatePlace(input.User.Place); err != nil {
			return UpdateUserOutput{}, err
		}
	}

	if err := i.userRepository.Update(tx, user); err != nil {
		return UpdateUserOutput{}, err
	}

	return i.presenter.Output(user), nil
}
