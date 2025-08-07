package work

import (
	"context"
	"devport/domain/domain_service"
	"devport/domain/repo/db"
	"devport/domain/repo/file_storage"
	"fmt"
	"mime/multipart"
	"time"
)

type (
	UploadWorkImageUseCase interface {
		Execute(context.Context, UploadWorkImageInput) (UploadWorkImageOutput, error)
	}

	UploadWorkImageInput struct {
		UserId string `validate:"required"`
		WorkId string `validate:"required"`

		Image       multipart.File
		ImageHeader *multipart.FileHeader
	}

	UploadWorkImagePresenter interface {
		Output(imageUrl string) UploadWorkImageOutput
	}

	UploadWorkImageOutput struct {
		ImageUrl string `json:"image_url"`
	}

	uploadWorkImageInteractor struct {
		userRepository      db.UserRepository
		workImageRepository db.WorkImagesRepository
		workImageStorage    file_storage.WorkImageStorageRepository
		presenter           UploadWorkImagePresenter
		ctxTimeout          time.Duration
	}
)

func NewUploadWorkImageInteractor(
	userRepository db.UserRepository,
	workImageStorage file_storage.WorkImageStorageRepository,
	workImageRepo db.WorkImagesRepository,
	presenter UploadWorkImagePresenter,
	t time.Duration,
) UploadWorkImageUseCase {
	return uploadWorkImageInteractor{
		workImageRepository: workImageRepo,
		userRepository:      userRepository,
		workImageStorage:    workImageStorage,
		presenter:           presenter,
		ctxTimeout:          t,
	}
}

func (i uploadWorkImageInteractor) Execute(tx context.Context, input UploadWorkImageInput) (UploadWorkImageOutput, error) {
	workImages, err := i.workImageRepository.FindByWorkId(tx, input.WorkId)

	if err != nil {
		return UploadWorkImageOutput{}, err
	}

	if len(workImages) == domain_service.MAX_IMAGE_LEN {
		return UploadWorkImageOutput{}, fmt.Errorf("画像は%d枚以上アップロードすることができません", domain_service.MAX_IMAGE_LEN)
	}

	var imagePath string
	var fileName string

	if err := i.workImageRepository.WithTransaction(tx, func(ttx context.Context) error {
		imagePath, fileName, err = i.workImageStorage.Upload(ttx, input.UserId, input.WorkId, input.Image, input.ImageHeader)

		if err != nil {
			return err
		}

		if err := i.workImageRepository.Create(ttx, input.WorkId, fileName); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return UploadWorkImageOutput{}, err
	}

	return i.presenter.Output(imagePath), nil
}
