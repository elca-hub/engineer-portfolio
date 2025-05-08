package user

import (
	"context"
	"devport/domain/model"
	"devport/domain/repo/sql"
	"devport/infra/file_uploader"
	"fmt"
	"mime/multipart"
	"time"
)

type (
	UploadBioImageUseCase interface {
		Execute(context.Context, UploadBioImageInput) (UploadBioImageOutput, error)
	}

	UploadBioImageInput struct {
		Image       multipart.File
		ImageHeader *multipart.FileHeader
		UserId      string `validate:"required"`
	}

	UploadBioImagePresenter interface {
		Output(imageId string) UploadBioImageOutput
	}

	UploadBioImageOutput struct {
		ImageId string `json:"image_id"`
	}

	uploadBioImageInteractor struct {
		fileUploader   file_uploader.FileUploader
		userRepository sql.UserRepository
		bioImageRepo   sql.BioImagesRepository
		presenter      UploadBioImagePresenter
		ctxTimeout     time.Duration
	}
)

func NewUploadBioImageInteractor(
	fileUploader file_uploader.FileUploader,
	presenter UploadBioImagePresenter,
	userRepository sql.UserRepository,
	bioImageRepo sql.BioImagesRepository,
	t time.Duration,
) UploadBioImageUseCase {
	return uploadBioImageInteractor{
		fileUploader:   fileUploader,
		bioImageRepo:   bioImageRepo,
		userRepository: userRepository,
		presenter:      presenter,
		ctxTimeout:     t,
	}
}

func (i uploadBioImageInteractor) Execute(tx context.Context, input UploadBioImageInput) (UploadBioImageOutput, error) {
	user, err := i.userRepository.FindById(tx, input.UserId)
	if err != nil {
		return UploadBioImageOutput{}, err
	}

	image, err := model.NewFileIcon(input.Image, input.ImageHeader, model.BIO_IMAGE_PATH)
	if err != nil {
		return UploadBioImageOutput{}, err
	}

	bioImages, err := i.bioImageRepo.FindByUserId(tx, user)

	if err != nil {
		return UploadBioImageOutput{}, err
	}

	if len(bioImages) == model.MaxBioImagesLen {
		return UploadBioImageOutput{}, fmt.Errorf("画像は%d枚以上アップロードすることができません", model.MaxBioImagesLen)
	}

	if err := i.bioImageRepo.WithTransaction(tx, func(tx context.Context) error {
		if err := i.bioImageRepo.Create(tx, user, image.GetFileName()); err != nil {
			return err
		}

		if err := i.fileUploader.UploadFile(image.GetFile(), image.GetFileName().GetObjectName()); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return UploadBioImageOutput{}, err
	}

	return UploadBioImageOutput{ImageId: image.GetFileName().GetFileName()}, nil
}
