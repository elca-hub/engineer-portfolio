package work

import (
	"context"
	"devport/domain/repo/db"
	"devport/domain/repo/file_storage"
	"mime/multipart"
	"time"
)

type (
	UploadWorkThumbnailUseCase interface {
		Execute(context.Context, UploadWorkThumbnailInput) (UploadWorkThumbnailOutput, error)
	}

	UploadWorkThumbnailInput struct {
		UserId string `validate:"required"`
		WorkId string `validate:"required"`

		Image       multipart.File
		ImageHeader *multipart.FileHeader
	}

	UploadWorkThumbnailPresenter interface {
		Output(imageUrl string) UploadWorkThumbnailOutput
	}

	UploadWorkThumbnailOutput struct {
		ImageUrl string `json:"image_url"`
	}

	uploadWorkThumbnailInteractor struct {
		workThumbnailStorage file_storage.WorkThumbnailStorageRepository
		presenter            UploadWorkThumbnailPresenter
		workRepository       db.WorkRepository
		ctxTimeout           time.Duration
	}
)

func NewUploadWorkThumbnailInteractor(
	workThumbnailStorage file_storage.WorkThumbnailStorageRepository,
	workRepository db.WorkRepository,
	presenter UploadWorkThumbnailPresenter,
	t time.Duration,
) UploadWorkThumbnailUseCase {
	return uploadWorkThumbnailInteractor{
		workRepository:       workRepository,
		workThumbnailStorage: workThumbnailStorage,
		presenter:            presenter,
		ctxTimeout:           t,
	}
}

func (i uploadWorkThumbnailInteractor) Execute(tx context.Context, input UploadWorkThumbnailInput) (UploadWorkThumbnailOutput, error) {
	work, err := i.workRepository.FindById(tx, input.UserId, input.WorkId)
	if err != nil {
		return UploadWorkThumbnailOutput{}, err
	}

	var (
		imagePath string
		fileName  string
	)

	if err := i.workRepository.WithTransaction(tx, func(ttx context.Context) error {
		currentThumbnail, err := i.workThumbnailStorage.FindByWorkId(ttx, input.UserId, input.WorkId)
		if err != nil {
			return err
		}

		if len(currentThumbnail) > 0 {
			err := i.workThumbnailStorage.Delete(ttx, input.UserId, input.WorkId, currentThumbnail[0])
			if err != nil {
				return err
			}
		}

		imagePath, fileName, err = i.workThumbnailStorage.Upload(ttx, input.UserId, input.WorkId, input.Image, input.ImageHeader)

		if err != nil {
			return err
		}

		work.UpdateThumbnailImageUrl(&fileName)

		if err := i.workRepository.Update(ttx, input.UserId, work); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return UploadWorkThumbnailOutput{}, err
	}

	return i.presenter.Output(imagePath), nil
}
