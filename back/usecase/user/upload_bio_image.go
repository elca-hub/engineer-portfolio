package user

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
	UploadBioImageUseCase interface {
		Execute(context.Context, UploadBioImageInput) (UploadBioImageOutput, error)
	}

	UploadBioImageInput struct {
		Image       multipart.File
		ImageHeader *multipart.FileHeader
		UserId      string `validate:"required"`
	}

	UploadBioImagePresenter interface {
		Output(imageUrl string) UploadBioImageOutput
	}

	UploadBioImageOutput struct {
		ImageUrl string `json:"image_url"`
	}

	uploadBioImageInteractor struct {
		userRepository     db.UserRepository
		bioImageRepository db.BioImagesRepository
		bioImageStorage    file_storage.BioImageStorageRepository
		presenter          UploadBioImagePresenter
		ctxTimeout         time.Duration
	}
)

func NewUploadBioImageInteractor(
	userRepository db.UserRepository,
	bioImageStorage file_storage.BioImageStorageRepository,
	bioImageRepo db.BioImagesRepository,
	presenter UploadBioImagePresenter,
	t time.Duration,
) UploadBioImageUseCase {
	return uploadBioImageInteractor{
		bioImageRepository: bioImageRepo,
		userRepository:     userRepository,
		bioImageStorage:    bioImageStorage,
		presenter:          presenter,
		ctxTimeout:         t,
	}
}

func (i uploadBioImageInteractor) Execute(tx context.Context, input UploadBioImageInput) (UploadBioImageOutput, error) {
	bioImages, err := i.bioImageRepository.FindByUserId(tx, input.UserId)

	if err != nil {
		return UploadBioImageOutput{}, err
	}

	if len(bioImages) == domain_service.MAX_IMAGE_LEN {
		return UploadBioImageOutput{}, fmt.Errorf("画像は%d枚以上アップロードすることができません", domain_service.MAX_IMAGE_LEN)
	}

	var imagePath string
	var fileName string

	if err := i.bioImageRepository.WithTransaction(tx, func(ttx context.Context) error {
		imagePath, fileName, err = i.bioImageStorage.Upload(ttx, input.UserId, input.Image, input.ImageHeader)

		if err != nil {
			return err
		}

		if err := i.bioImageRepository.Create(ttx, input.UserId, fileName); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return UploadBioImageOutput{}, err
	}

	return i.presenter.Output(imagePath), nil
}
