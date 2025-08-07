package work

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/db"
	"devport/domain/repo/file_storage"
	"devport/infra/email"
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

type (
	CreateWorkUseCase interface {
		Execute(context.Context, CreateWorkInput) (CreateWorkOutput, error)
	}

	CreateWorkInput struct {
		UserId string `validate:"required"`
	}

	CreateWorkOutput struct {
		Work *dto.WorkDTO `json:"work"`
	}

	CreateWorkPresenter interface {
		Output(work *model.Work) CreateWorkOutput
	}

	createWorkInteractor struct {
		userRepository           db.UserRepository
		workRepository           db.WorkRepository
		workHavingTagsRepository db.WorkHavingTagsRepository
		presenter                CreateWorkPresenter
		email                    email.Email
		ctxTimeout               time.Duration
	}
)

func NewCreateWorkInteractor(
	userRepository db.UserRepository,
	workRepository db.WorkRepository,
	workHavingTagsRepository db.WorkHavingTagsRepository,
	workStorage file_storage.WorkSentenceStorageRepository,
	presenter CreateWorkPresenter,
	email email.Email,
	t time.Duration,
) CreateWorkUseCase {
	return createWorkInteractor{
		userRepository:           userRepository,
		workRepository:           workRepository,
		workHavingTagsRepository: workHavingTagsRepository,
		presenter:                presenter,
		email:                    email,
		ctxTimeout:               t,
	}
}

func (i createWorkInteractor) Execute(ctx context.Context, input CreateWorkInput) (CreateWorkOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	if isExists, err := i.userRepository.ExistsById(ctx, input.UserId); err != nil || !isExists {
		if err != nil {
			return CreateWorkOutput{}, err
		}
		return CreateWorkOutput{}, errors.New("ユーザが存在しません")
	}

	id, err := uuid.NewV4()

	if err != nil {
		return CreateWorkOutput{}, err
	}

	sortIndex, err := i.workRepository.GetMaxSortIndex(ctx, input.UserId)

	if err != nil {
		return CreateWorkOutput{}, err
	}

	work := model.NewWorkInit(id.String(), sortIndex)

	wg, err := i.workRepository.FindAll(ctx, input.UserId)
	if err != nil {
		return CreateWorkOutput{}, err
	}

	if err := wg.AddWork(work); err != nil {
		return CreateWorkOutput{}, err
	}

	if err := i.workRepository.Create(ctx, input.UserId, work); err != nil {
		return CreateWorkOutput{}, err
	}

	return i.presenter.Output(work), nil
}
