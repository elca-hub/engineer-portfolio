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
	FetchWorkUseCase interface {
		Execute(context.Context, FetchWorkInput) (FetchWorkOutput, error)
	}

	FetchWorkInput struct {
		UserId string `validate:"required"`

		Type string `validate:"required"`
	}

	FetchWorkOutput struct {
		Works []*dto.WorkDTO `json:"works"`
	}

	FetchWorkPresenter interface {
		Output(works []*model.Work) FetchWorkOutput
	}

	fetchWorkInteractor struct {
		userRepository db.UserRepository
		workRepository db.WorkRepository
		presenter      FetchWorkPresenter
		ctxTimeout     time.Duration
	}
)

func NewFetchWorkInteractor(
	userRepository db.UserRepository,
	workRepository db.WorkRepository,
	presenter FetchWorkPresenter,
	t time.Duration,
) FetchWorkUseCase {
	return fetchWorkInteractor{
		userRepository: userRepository,
		workRepository: workRepository,
		presenter:      presenter,
		ctxTimeout:     t,
	}
}

func (i fetchWorkInteractor) Execute(ctx context.Context, input FetchWorkInput) (FetchWorkOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	if isExists, err := i.userRepository.ExistsById(ctx, input.UserId); err != nil || !isExists {
		if err != nil {
			return FetchWorkOutput{}, err
		}
		return FetchWorkOutput{}, errors.New("ユーザが存在しません")
	}

	wg, err := i.workRepository.FindAll(ctx, input.UserId)

	if err != nil {
		return FetchWorkOutput{}, err
	}

	switch input.Type {
	case "all":
		return i.presenter.Output(wg.Works()), nil
	case "public":
		return i.presenter.Output(wg.PublicWorks()), nil
	case "draft":
		return i.presenter.Output(wg.DraftWorks()), nil
	}

	return FetchWorkOutput{}, errors.New("不正なタイプです")
}
