package user

import (
	"context"
	"devport/domain/domain_service"
	"devport/domain/repo/db"
	"devport/domain/repo/file_storage"
	"errors"
	"slices"
	"time"

	"golang.org/x/sync/errgroup"
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
		bioService         *domain_service.BioService
		userRepository     db.UserRepository
		bioImageRepo       db.BioImagesRepository
		bioSentenceStorage file_storage.BioSentenceStorageRepository
		bioImageStorage    file_storage.BioImageStorageRepository
		presenter          UploadBioPresenter
		ctxTimeout         time.Duration
	}
)

func NewUploadBioInteractor(
	bioService *domain_service.BioService,
	userRepository db.UserRepository,
	bioImageRepo db.BioImagesRepository,
	bioSentenceStorage file_storage.BioSentenceStorageRepository,
	bioImageStorage file_storage.BioImageStorageRepository,
	presenter UploadBioPresenter,
	t time.Duration,
) UploadBioUseCase {
	return uploadBioInteractor{
		bioService:         bioService,
		presenter:          presenter,
		userRepository:     userRepository,
		bioImageRepo:       bioImageRepo,
		bioSentenceStorage: bioSentenceStorage,
		bioImageStorage:    bioImageStorage,
		ctxTimeout:         t,
	}
}

func (i uploadBioInteractor) Execute(tx context.Context, input UploadBioInput) (UploadBioOutput, error) {
	user, err := i.userRepository.FindById(tx, input.UserId)
	if err != nil {
		return UploadBioOutput{}, err
	}

	uploadedImageName, err := i.bioImageStorage.FindByUserId(tx, user.ID())
	if err != nil {
		return UploadBioOutput{}, err
	}

	imageIds := i.bioService.GetImageIds(user.ID(), input.Bio)
	if i.bioService.IsFullImage() {
		return UploadBioOutput{}, errors.New("自己紹介文に使用している画像が多すぎます")
	}

	for _, bioFileName := range imageIds {
		/**
		array内にtargetが存在するかどうか
		*/
		isExistsInArray := func(array []string, target string) bool {
			for _, item := range array {
				if item == target {
					return true
				}
			}
			return false
		}

		// 画像がアップロードされていない場合はエラー
		if !isExistsInArray(uploadedImageName, bioFileName) {
			return UploadBioOutput{}, errors.New("画像が存在しません")
		}
	}

	if _, err := i.bioSentenceStorage.Upload(tx, user.ID(), input.Bio); err != nil {
		return UploadBioOutput{}, err
	}

	differenceArray := func(a, b []string) []string {
		diff := []string{}
		for _, item := range a {
			if !slices.Contains(b, item) {
				diff = append(diff, item)
			}
		}
		return diff
	}

	if input.IsDeleteImage {
		// ストレージ上から削除する画像の名前を取得
		deleteImageTarget := differenceArray(uploadedImageName, imageIds)

		errG := new(errgroup.Group)
		for _, imageName := range deleteImageTarget {
			imageName := imageName
			errG.Go(func() error {
				return i.bioImageStorage.Delete(tx, user.ID(), imageName)
			})
		}

		dbImages, err := i.bioImageRepo.FindByUserId(tx, user.ID())
		if err != nil {
			return UploadBioOutput{}, err
		}

		deleteImageTargetDB := differenceArray(dbImages, imageIds)

		for _, imageName := range deleteImageTargetDB {
			in := imageName // 確かこうしないとバグる気がした。多分スコープの問題？
			errG.Go(func() error {
				return i.bioImageRepo.Delete(tx, user, in)
			})
		}

		if err := errG.Wait(); err != nil {
			return UploadBioOutput{}, err
		}
	}

	return UploadBioOutput{Bio: input.Bio}, nil
}
