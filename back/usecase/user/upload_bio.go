package user

import (
	"context"
	"devport/domain/model"
	"devport/domain/repo/sql"
	"devport/infra/file_uploader"
	"time"
)

type (
	UploadBioUseCase interface {
		Execute(context.Context, UploadBioInput) (UploadBioOutput, error)
	}

	UploadBioInput struct {
		Bio    string `json:"bio" validate:"max=1000"`
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
		bioImageRepo   sql.BioImagesRepository
		ctxTimeout     time.Duration
	}
)

func NewUploadBioInteractor(
	fileUploader file_uploader.FileUploader,
	presenter UploadBioPresenter,
	userRepository sql.UserRepository,
	bioImageRepo sql.BioImagesRepository,
	t time.Duration,
) UploadBioUseCase {
	return uploadBioInteractor{
		fileUploader:   fileUploader,
		presenter:      presenter,
		userRepository: userRepository,
		bioImageRepo:   bioImageRepo,
		ctxTimeout:     t,
	}
}

func (i uploadBioInteractor) Execute(tx context.Context, input UploadBioInput) (UploadBioOutput, error) {
	user, err := i.userRepository.FindById(tx, input.UserId)
	if err != nil {
		return UploadBioOutput{}, err
	}

	bioModel, err := model.NewBio(user.ID(), user.BioPath(), input.Bio)

	if err != nil {
		return UploadBioOutput{}, err
	}

	if err != nil {
		return UploadBioOutput{}, err
	}

	if err := i.fileUploader.UploadFile([]byte(input.Bio), bioModel.ObjectName()); err != nil {
		return UploadBioOutput{}, err
	}

	user.UpdateBioPath(bioModel.ID())

	if err := i.userRepository.Update(tx, user); err != nil {
		return UploadBioOutput{}, err
	}

	return UploadBioOutput{BioName: bioModel.ID()}, nil
}
