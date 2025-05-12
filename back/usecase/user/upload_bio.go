package user

import (
	"context"
	"database/sql"
	"devport/domain/model"
	repo_sql "devport/domain/repo/sql"
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
		Bio           string `json:"bio" validate:"max=5000"`
		UserId        string `validate:"required"`
		IsDeleteImage bool
	}

	UploadBioPresenter interface {
		Output(bioName string) UploadBioOutput
	}

	UploadBioOutput struct {
		Bio string `json:"bio"`
	}

	uploadBioInteractor struct {
		fileUploader   file_uploader.FileUploader
		presenter      UploadBioPresenter
		userRepository repo_sql.UserRepository
		bioImageRepo   repo_sql.BioImagesRepository
		ctxTimeout     time.Duration
	}
)

func NewUploadBioInteractor(
	fileUploader file_uploader.FileUploader,
	presenter UploadBioPresenter,
	userRepository repo_sql.UserRepository,
	bioImageRepo repo_sql.BioImagesRepository,
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

	objectImages, err := i.fileUploader.GetFiles(model.BIO_IMAGE_PATH)
	if err != nil {
		return UploadBioOutput{}, err
	}

	dbImagePaths, err := i.bioImageRepo.FindByUserId(tx, user)
	if err != nil {
		return UploadBioOutput{}, err
	}

	for _, imageId := range bioModel.ImageIds() {
		fileIconName, err := model.NewFileIconName(imageId, model.BIO_IMAGE_PATH)
		if err != nil {
			return UploadBioOutput{}, err
		}

		isExistsInArray := func(array []*model.FileIconName, target string) bool {
			for _, item := range array {
				if item.GetFileName() == target {
					return true
				}
			}
			return false
		}

		// 画像がアップロードされていない場合はエラー
		if !isExistsInArray(objectImages, fileIconName.GetFileName()) {
			return UploadBioOutput{}, errors.New("画像が存在しません")
		}

		if !isExistsInArray(dbImagePaths, fileIconName.GetFileName()) {
			if err := i.bioImageRepo.Create(tx, user, fileIconName); err != nil {
				return UploadBioOutput{}, err
			}
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

		diff, err := difference(objectImages, bioImageFiles)
		if err != nil {
			return err
		}

		fmt.Printf("diff: %v\n", diff)

		for _, d := range diff {
			if input.IsDeleteImage {
				i.fileUploader.DeleteFile(d.GetObjectName())
			}
			if err := i.bioImageRepo.Delete(tx, user, d); err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					continue
				}
				return err
			}
		}

		return nil
	}); err != nil {
		return UploadBioOutput{}, err
	}

	return UploadBioOutput{Bio: input.Bio}, nil
}
