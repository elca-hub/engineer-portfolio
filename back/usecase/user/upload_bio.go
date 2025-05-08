package user

import (
	"context"
	"devport/domain/model"
	"devport/domain/repo/sql"
	"devport/infra/file_uploader"
	"errors"
	"fmt"
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

	for _, imageId := range bioModel.ImageIds() {
		fileIconName, err := model.NewFileIconName(imageId, model.BIO_IMAGE_PATH)
		if err != nil {
			return UploadBioOutput{}, err
		}

		exists, err := i.bioImageRepo.IsExistsFileName(tx, fileIconName)
		if err != nil {
			return UploadBioOutput{}, err
		}

		if !exists {
			return UploadBioOutput{}, errors.New("画像が存在しません")
		}
	}

	if err := i.fileUploader.UploadFile([]byte(input.Bio), bioModel.ObjectName()); err != nil {
		return UploadBioOutput{}, err
	}

	user.UpdateBioPath(bioModel.ID())

	if err := i.userRepository.Update(tx, user); err != nil {
		return UploadBioOutput{}, err
	}

	if err := i.bioImageRepo.WithTransaction(tx, func(tx context.Context) error {
		uploadedFiles, err := i.bioImageRepo.FindByUserId(tx, user)

		if err != nil {
			return err
		}
		difference := func(a, b []*model.FileIconName) ([]*model.FileIconName, error) {
			aString := make([]string, len(a))
			bString := make([]string, len(b))
			for i, item := range a {
				aString[i] = item.GetFileName()
			}
			for i, item := range b {
				bString[i] = item.GetFileName()
			}

			// aStringからbStringを引いた配列を返す
			diff := make([]string, 0)
			for _, item := range aString {
				if !contains(bString, item) {
					diff = append(diff, item)
				}
			}

			diffFiles := make([]*model.FileIconName, len(diff))
			for i, item := range diff {
				diffFiles[i], err = model.NewFileIconName(item, model.BIO_IMAGE_PATH)
				if err != nil {
					return nil, err
				}
			}

			return diffFiles, nil
		}

		bioImages := bioModel.ImageIds()

		bioImageFiles := make([]*model.FileIconName, len(bioImages))

		for i, imageId := range bioImages {
			bioImageFiles[i], err = model.NewFileIconName(imageId, model.BIO_IMAGE_PATH)
			if err != nil {
				return err
			}
		}

		diff, err := difference(uploadedFiles, bioImageFiles)
		if err != nil {
			return err
		}

		for _, d := range diff {
			fmt.Println(d.GetFileName())
			i.fileUploader.DeleteFile(d.GetObjectName())
			i.bioImageRepo.Delete(tx, user, d)
		}

		return nil
	}); err != nil {
		return UploadBioOutput{}, err
	}

	return UploadBioOutput{BioName: bioModel.ID()}, nil
}
