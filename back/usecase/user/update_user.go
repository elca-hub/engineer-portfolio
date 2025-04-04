package user

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/sql"
	"devport/infra/file_uploader"
	"fmt"
	"github.com/google/uuid"
	"mime/multipart"
	"time"
)

type (
	UpdateUserUseCase interface {
		Execute(context.Context, UpdateUserInput) (UpdateUserOutput, error)
	}

	UpdateUserInput struct {
		UserId     string                `validate:"required"`
		Icon       multipart.File        `validate:"required"`
		IconHeader *multipart.FileHeader `validate:"required"`
	}

	UpdateUserPresenter interface {
		Output(user *model.User) UpdateUserOutput
	}

	UpdateUserOutput struct {
		User dto.UserDTO `json:"user"`
	}

	updateUserInterator struct {
		sqlRepository sql.UserRepository
		fileUploader  file_uploader.FileUploader
		presenter     UpdateUserPresenter
		ctxTimeout    time.Duration
	}
)

func NewUpdateUserInterator(
	sqlRepository sql.UserRepository,
	fileUploader file_uploader.FileUploader,
	presenter UpdateUserPresenter,
	t time.Duration,
) UpdateUserUseCase {
	return updateUserInterator{
		sqlRepository: sqlRepository,
		fileUploader:  fileUploader,
		presenter:     presenter,
		ctxTimeout:    t,
	}
}

func (i updateUserInterator) Execute(tx context.Context, input UpdateUserInput) (UpdateUserOutput, error) {
	ctx, cancel := context.WithTimeout(tx, i.ctxTimeout)
	defer cancel()

	user, err := i.sqlRepository.FindById(ctx, input.UserId)

	if err != nil {
		return UpdateUserOutput{}, err
	}

	err = i.sqlRepository.WithTransaction(ctx, func(tx context.Context) error {
		iconFile, err := model.NewFileIcon(input.Icon, input.IconHeader)
		if err != nil {
			return err
		}

		fileName, _ := uuid.NewUUID()

		iconPathName, err := i.fileUploader.UploadFile(iconFile.GetFile().Bytes(), "devport", fmt.Sprintf("%s/%s.%s", user.ID(), fileName.String(), iconFile.GetExtension()))

		if err != nil {
			return err
		}

		user.SetIconPath(iconPathName)

		if err := i.sqlRepository.Update(tx, user); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return UpdateUserOutput{}, err
	}

	return i.presenter.Output(user), nil
}
