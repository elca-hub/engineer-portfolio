package user

import (
	"context"
	"devport/domain/repo/sql"
	"devport/infra/file_uploader"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type (
	UploadBioUseCase interface {
		Execute(context.Context, UploadBioInput) (UploadBioOutput, error)
	}

	UploadBioInput struct {
		Bio    string `json:"bio" validate:"required,max=1000,min=1"`
		UserId string `validate:"required"`
	}

	UploadBioPresenter interface {
		Output(bioName string) UploadBioOutput
	}

	UploadBioOutput struct {
		BioName string `json:"bio_name"`
	}

	uploadBioInteractor struct {
		fileUploader   file_uploader.FileUploader
		presenter      UploadBioPresenter
		userRepository sql.UserRepository
		ctxTimeout     time.Duration
	}
)

func NewUploadBioInteractor(
	fileUploader file_uploader.FileUploader,
	presenter UploadBioPresenter,
	userRepository sql.UserRepository,
	t time.Duration,
) UploadBioUseCase {
	return uploadBioInteractor{
		fileUploader:   fileUploader,
		presenter:      presenter,
		userRepository: userRepository,
		ctxTimeout:     t,
	}
}

func (i uploadBioInteractor) Execute(tx context.Context, input UploadBioInput) (UploadBioOutput, error) {
	user, err := i.userRepository.FindById(tx, input.UserId)
	if err != nil {
		return UploadBioOutput{}, err
	}

	var bioId string

	// 新規にbioをアップロードする場合はuuidを生成
	if user.BioPath() != "" {
		bioId = user.BioPath()
	} else {
		bioId = uuid.New().String()
	}

	objectName := fmt.Sprintf("bio/%s/%s.md", user.ID(), bioId)

	if err := i.fileUploader.UploadFile([]byte(input.Bio), objectName); err != nil {
		return UploadBioOutput{}, err
	}

	user.UpdateBioPath(bioId)

	if err := i.userRepository.Update(tx, user); err != nil {
		return UploadBioOutput{}, err
	}

	return UploadBioOutput{BioName: bioId}, nil
}
