package user

import (
	"context"
	"devport/domain/model"
	"devport/domain/repo/sql"
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
		IsExists bool `json:"is_exists"`
	}

	IsExistsUserPresenter interface {
		Output(isExists bool) IsExistsUserOutput
	}

	isExistsUserInteractor struct {
		sqlRepository sql.UserRepository
		presenter     IsExistsUserPresenter
		ctxTimeout    time.Duration
	}
)

func NewIsExistsUserInteractor(
	sqlRepository sql.UserRepository,
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
		return i.presenter.Output(false), err
	}

	isExistsMail, err := i.sqlRepository.Exists(ctx, e)

	if err != nil {
		return i.presenter.Output(false), err
	}

	return i.presenter.Output(isExistsMail), nil
}
