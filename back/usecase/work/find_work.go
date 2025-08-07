package work

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/db"
	"errors"
	"time"
)

type (
	FindWorkUseCase interface {
		Execute(context.Context, FindWorkInput) (FindWorkOutput, error)
	}

	FindWorkInput struct {
		UserId string `validate:"required"`
		WorkId string `validate:"required"`
	}

	FindWorkOutput struct {
		Work *dto.WorkDTO `json:"work"`
	}

	FindWorkPresenter interface {
		Output(work *model.Work) FindWorkOutput
	}

	findWorkInteractor struct {
		userRepository db.UserRepository
		workRepository db.WorkRepository
		presenter      FindWorkPresenter
		ctxTimeout     time.Duration
	}
)

func NewFindWorkInteractor(
	userRepository db.UserRepository,
	workRepository db.WorkRepository,
	presenter FindWorkPresenter,
	t time.Duration,
) FindWorkUseCase {
	return findWorkInteractor{
		userRepository: userRepository,
		workRepository: workRepository,
		presenter:      presenter,
		ctxTimeout:     t,
	}
}

func (i findWorkInteractor) Execute(ctx context.Context, input FindWorkInput) (FindWorkOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	if isExists, err := i.userRepository.ExistsById(ctx, input.UserId); err != nil || !isExists {
		if err != nil {
			return FindWorkOutput{}, err
		}
		return FindWorkOutput{}, errors.New("ユーザが存在しません")
	}

	work, err := i.workRepository.FindById(ctx, input.UserId, input.WorkId)
	if err != nil {
		return FindWorkOutput{}, err
	}

	return i.presenter.Output(work), nil
}
