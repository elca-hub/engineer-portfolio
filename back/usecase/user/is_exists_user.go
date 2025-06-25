package user

import (
	"context"
	"devport/domain/model"
	"devport/domain/repo/db"
	"time"
)

type (
	IsExistsUserUseCase interface {
		Execute(context.Context, IsExistsUserInput) (IsExistsUserOutput, error)
	}

	IsExistsUserInput struct {
		Email string `json:"email" validate:"required,email"`
	}

	IsExistsUserOutput struct {
		IsExists bool   `json:"is_exists"`
		UserId   string `json:"user_id"`
	}

	IsExistsUserPresenter interface {
		Output(isExists bool, userId string) IsExistsUserOutput
	}

	isExistsUserInteractor struct {
		sqlRepository db.UserRepository
		presenter     IsExistsUserPresenter
		ctxTimeout    time.Duration
	}
)

func NewIsExistsUserInteractor(
	sqlRepository db.UserRepository,
	presenter IsExistsUserPresenter,
	t time.Duration,
) IsExistsUserUseCase {
	return isExistsUserInteractor{
		sqlRepository: sqlRepository,
		presenter:     presenter,
		ctxTimeout:    t,
	}
}

func (i isExistsUserInteractor) Execute(ctx context.Context, input IsExistsUserInput) (IsExistsUserOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	e, err := model.NewEmail(input.Email)

	if err != nil {
		return i.presenter.Output(false, ""), err
	}

	isExistsMail, err := i.sqlRepository.Exists(ctx, e)

	if err != nil {
		return i.presenter.Output(false, ""), err
	}

	if !isExistsMail {
		return i.presenter.Output(false, ""), nil
	}

	user, err := i.sqlRepository.FindByEmail(ctx, e, nil)

	if err != nil {
		return i.presenter.Output(false, ""), err
	}

	return i.presenter.Output(isExistsMail, user.ID()), nil
}
