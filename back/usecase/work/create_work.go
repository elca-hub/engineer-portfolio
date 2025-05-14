package work

import (
	"context"
	"devport/domain/model"
	"devport/domain/repo/db"
	"devport/infra/file_uploader"
	"time"

	"github.com/google/uuid"
)

type (
	CreateWorkUseCase interface {
		Execute(context.Context, CreateWorkInput) (CreateWorkOutput, error)
	}

	CreateWorkInput struct {
		UserID    string `validate:"required"`
		Title     string `json:"title" validate:"required,max=255,min=1"`
		Content   string `json:"content" validate:"required"`
		GithubURL string `json:"github_url"`
	}

	CreateWorkOutput struct {
		Title string `json:"title"`
	}

	CreateWorkPresenter interface {
		Output(title string) CreateWorkOutput
	}

	createWorkInteractor struct {
		workRepository db.WorkRepository
		userRepository db.UserRepository
		fileUploader   file_uploader.FileUploader
		presenter      CreateWorkPresenter
		ctxTimeout     time.Duration
	}
)

func NewCreateWorkInteractor(
	workRepository db.WorkRepository,
	userRepository db.UserRepository,
	fileUploader file_uploader.FileUploader,
	presenter CreateWorkPresenter,
	t time.Duration,
) CreateWorkUseCase {
	return createWorkInteractor{
		workRepository: workRepository,
		userRepository: userRepository,
		fileUploader:   fileUploader,
		presenter:      presenter,
		ctxTimeout:     t,
	}
}

func (i createWorkInteractor) Execute(ctx context.Context, input CreateWorkInput) (CreateWorkOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	user, err := i.userRepository.FindById(ctx, input.UserID)

	if err != nil {
		return CreateWorkOutput{}, err
	}

	var work *model.Work

	err = i.workRepository.WithTransaction(ctx, func(ctx context.Context) error {
		var err error

		workId := uuid.New().String()

		work, err = model.NewWork(workId, user.ID(), input.Title, input.Content, nil, -1, input.GithubURL, time.Now(), time.Now())

		if err != nil {
			return err
		}

		if err := i.fileUploader.UploadFile([]byte(work.Content()), work.ContentFileName().GetObjectName()); err != nil {
			return err
		}

		return i.workRepository.Create(ctx, work)
	})

	if err != nil {
		return CreateWorkOutput{}, err
	}

	return i.presenter.Output(work.Title()), nil
}
