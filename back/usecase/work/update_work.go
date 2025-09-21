package work

import (
	"context"
	"devport/domain/domain_service"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/db"
	"devport/domain/repo/file_storage"
	"errors"
	"slices"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/sync/errgroup"
)

type (
	UpdateWorkUseCase interface {
		Execute(context.Context, UpdateWorkInput) (UpdateWorkOutput, error)
	}

	UpdateWorkInput struct {
		UserId string `validate:"required"`

		Work *dto.WorkDTO `json:"work" validate:"required"`

		IsDeleteImage bool
	}

	UpdateWorkOutput struct {
		Work *dto.WorkDTO `json:"work"`
	}

	UpdateWorkPresenter interface {
		Output(work *model.Work) UpdateWorkOutput
	}

	updateWorkInteractor struct {
		workSentenceService *domain_service.WorkSentenceService
		workUrlService      *domain_service.WorkUrlService

		userRepository           db.UserRepository
		workRepository           db.WorkRepository
		workHavingTagsRepository db.WorkHavingTagsRepository
		workTagRepository        db.WorkTagRepository
		workUrlRepository        db.WorkUrlRepository
		workSentenceStorage      file_storage.WorkSentenceStorageRepository
		workImageStorage         file_storage.WorkImageStorageRepository
		workImageRepository      db.WorkImagesRepository
		presenter                UpdateWorkPresenter
		ctxTimeout               time.Duration
	}
)

func NewUpdateWorkInteractor(
	workSentenceService *domain_service.WorkSentenceService,
	workUrlService *domain_service.WorkUrlService,
	userRepository db.UserRepository,
	workRepository db.WorkRepository,
	workHavingTagsRepository db.WorkHavingTagsRepository,
	workTagRepository db.WorkTagRepository,
	workUrlRepository db.WorkUrlRepository,
	workSentenceStorage file_storage.WorkSentenceStorageRepository,
	workImageStorage file_storage.WorkImageStorageRepository,
	workImageRepository db.WorkImagesRepository,
	presenter UpdateWorkPresenter,
	t time.Duration,
) UpdateWorkUseCase {
	return updateWorkInteractor{
		workSentenceService:      workSentenceService,
		workUrlService:           workUrlService,
		userRepository:           userRepository,
		workRepository:           workRepository,
		workHavingTagsRepository: workHavingTagsRepository,
		workTagRepository:        workTagRepository,
		workUrlRepository:        workUrlRepository,
		workSentenceStorage:      workSentenceStorage,
		workImageStorage:         workImageStorage,
		workImageRepository:      workImageRepository,
		presenter:                presenter,
		ctxTimeout:               t,
	}
}

func (i updateWorkInteractor) Execute(ctx context.Context, input UpdateWorkInput) (UpdateWorkOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	errG := new(errgroup.Group)

	if isExists, err := i.userRepository.ExistsById(ctx, input.UserId); err != nil || !isExists {
		if err != nil {
			return UpdateWorkOutput{}, err
		}
		return UpdateWorkOutput{}, errors.New("ユーザが存在しません")
	}

	work, err := i.workRepository.FindById(ctx, input.UserId, input.Work.ID)

	if err != nil {
		return UpdateWorkOutput{}, err
	}

	currentTags, err := i.workHavingTagsRepository.Find(ctx, work.ID())
	if err != nil {
		return UpdateWorkOutput{}, err
	}

	convertToTagIds := func(tags []*model.WorkTag) []string {
		tagIds := make([]string, len(tags))
		for i, tag := range tags {
			tagIds[i] = tag.ID()
		}
		return tagIds
	}

	currentTagIds := convertToTagIds(currentTags)

	if err := work.UpdateTags(currentTagIds); err != nil {
		return UpdateWorkOutput{}, err
	}

	work, err = i.updateModel(work, input.Work)
	if err != nil {
		return UpdateWorkOutput{}, err
	}

	errG.Go(func() error {
		if err := i.workRepository.Update(ctx, input.UserId, work); err != nil {
			return err
		}

		return nil
	})

	/*
		タグの更新
	*/
	errG.Go(func() error {
		if err := i.workHavingTagsRepository.DeleteByWorkId(ctx, work.ID()); err != nil {
			return err
		}

		appendTagIds := []string{}

		// タグを追加
		for _, tag := range input.Work.Tags {
			isExists, err := i.workTagRepository.Exists(ctx, tag.Name)
			if err != nil {
				return err
			}

			if isExists {
				tag, err := i.workTagRepository.FindByName(ctx, tag.Name)
				if err != nil {
					return err
				}

				appendTagIds = append(appendTagIds, tag.ID())
				continue
			}

			// タグが存在しない場合新規作成
			id, err := uuid.NewV4()

			if err != nil {
				return err
			}

			appendTag, err := model.NewWorkTag(id.String(), tag.Name)
			if err != nil {
				return err
			}

			if err := i.workTagRepository.Create(ctx, appendTag); err != nil {
				return err
			}

			appendTagIds = append(appendTagIds, appendTag.ID())
		}

		if err := i.workHavingTagsRepository.CreateByWorkId(ctx, work.ID(), appendTagIds); err != nil {
			return err
		}

		err = i.workTagRepository.WithTransaction(ctx, func(ctx context.Context) error {
			if err := i.workTagRepository.DeleteNoUsed(ctx); err != nil {
				return err
			}
			return nil
		})

		if err != nil {
			return err
		}

		return nil
	})

	/*
		外部サービスの更新
	*/
	errG.Go(func() error {
		if err := i.workUrlRepository.DeleteByWorkId(ctx, work.ID()); err != nil {
			return err
		}

		esuDto := make([]*model.WorkUrl, len(input.Work.ExternalServiceUrls))
		for num, url := range input.Work.ExternalServiceUrls {
			id, err := uuid.NewV4()
			if err != nil {
				return err
			}

			var urlModel *model.WorkUrl

			if url.Title != "" {
				urlModel, err = model.NewWorkUrl(id.String(), url.Url, url.Title)
				if err != nil {
					return err
				}
			} else {
				title, err := i.workUrlService.FetchTitle(url.Url)

				if err != nil {
					return err
				}

				urlModel, err = model.NewWorkUrl(id.String(), url.Url, title)
				if err != nil {
					return err
				}
			}

			esuDto[num] = urlModel
			if err != nil {
				return err
			}
		}

		if err := i.workUrlRepository.CreateByWorkId(ctx, work.ID(), esuDto); err != nil {
			return err
		}

		if err := work.UpdateExternalServiceUrls(esuDto); err != nil {
			return err
		}

		return nil
	})

	/*
		画像の更新
	*/
	errG.Go(func() error {
		uploadedImageName, err := i.workImageStorage.FindByWorkId(ctx, input.UserId, work.ID())
		if err != nil {
			return err
		}

		imageIds := i.workSentenceService.GetImageIds(input.UserId, work.ID(), work.Content())

		if i.workSentenceService.IsFullImage() {
			return errors.New("使用している画像が多いです")
		}

		for _, imageId := range imageIds {
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
			if !isExistsInArray(uploadedImageName, imageId) {
				return errors.New("画像が存在しません")
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
						return i.workImageStorage.Delete(ctx, input.UserId, work.ID(), imageName)
					})
				}

				dbImages, err := i.workImageRepository.FindByWorkId(ctx, work.ID())
				if err != nil {
					return err
				}

				deleteImageTargetDB := differenceArray(dbImages, imageIds)

				for _, imageName := range deleteImageTargetDB {
					in := imageName // 確かこうしないとバグる気がした。多分スコープの問題？
					errG.Go(func() error {
						return i.workImageRepository.Delete(ctx, work, in)
					})
				}

				if err := errG.Wait(); err != nil {
					return err
				}
			}
		}

		return nil
	})

	/*
		文章の更新
	*/
	errG.Go(func() error {
		if _, err := i.workSentenceStorage.Upload(ctx, input.UserId, work.ID(), work.Content()); err != nil {
			return err
		}
		return nil
	})

	if err := errG.Wait(); err != nil {
		return UpdateWorkOutput{}, err
	}

	return i.presenter.Output(work), nil
}

/*
モデルの更新
*/
func (i updateWorkInteractor) updateModel(work *model.Work, input *dto.WorkDTO) (*model.Work, error) {
	if err := work.UpdateTitle(input.Title); err != nil {
		return nil, err
	}
	if err := work.UpdateContent(input.Content); err != nil {
		return nil, err
	}

	if err := work.UpdateGithubRepositoryUrl(input.GithubRepositoryUrl); err != nil {
		return nil, err
	}

	if err := work.UpdatePublishStatus(model.PublishStatus(input.PublishStatus)); err != nil {
		return nil, err
	}

	tagNames := make([]string, len(input.Tags))
	for i, tag := range input.Tags {
		tagNames[i] = tag.Name
	}

	if err := work.UpdateTags(tagNames); err != nil {
		return nil, err
	}

	return work, nil
}
