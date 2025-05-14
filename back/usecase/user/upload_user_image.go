package user

import (
	"context"
	"devport/domain/model"
	"devport/infra/file_uploader"
	"mime/multipart"
	"time"

	"golang.org/x/sync/errgroup"
)

type (
	UploadUserImageUseCase interface {
		Execute(context.Context, UploadUserImageInput) (UploadUserImageOutput, error)
	}

	UploadUserImageInput struct {
		Icon         multipart.File
		IconHeader   *multipart.FileHeader
		Header       multipart.File
		HeaderHeader *multipart.FileHeader
	}

	UploadUserImagePresenter interface {
		Output(iconName string, headerIconName string) UploadUserImageOutput
	}

	UploadUserImageOutput struct {
		IconName       string `json:"icon_name"`
		HeaderIconName string `json:"header_icon_name"`
	}

	uploadUserImageInterator struct {
		fileUploader file_uploader.FileUploader
		presenter    UploadUserImagePresenter
		ctxTimeout   time.Duration
	}
)

func NewUploadUserImageInterator(
	fileUploader file_uploader.FileUploader,
	presenter UploadUserImagePresenter,
	t time.Duration,
) UploadUserImageUseCase {
	return uploadUserImageInterator{
		fileUploader: fileUploader,
		presenter:    presenter,
		ctxTimeout:   t,
	}
}

func (i uploadUserImageInterator) Execute(tx context.Context, input UploadUserImageInput) (UploadUserImageOutput, error) {
	// 並行処理用のグループを作成
	g := new(errgroup.Group)

	var iconName string
	var headerIconName string

	iconNameChan := make(chan string, 1)
	headerIconNameChan := make(chan string, 1)

	// アイコン画像のアップロード
	if input.Icon != nil && input.IconHeader != nil {
		g.Go(func() error {
			iconFile, err := model.NewBucketFile(input.Icon, input.IconHeader, model.ICON_PATH)
			if err != nil {
				return err
			}

			if err := i.fileUploader.UploadFile(iconFile.GetFile(), iconFile.GetFileName().GetObjectName()); err != nil {
				return err
			}
			iconNameChan <- iconFile.GetFileName().GetFileName()
			return nil
		})
	}

	// ヘッダー画像のアップロード
	if input.Header != nil && input.HeaderHeader != nil {
		g.Go(func() error {
			headerFile, err := model.NewBucketFile(input.Header, input.HeaderHeader, model.HEADER_PATH)
			if err != nil {
				return err
			}

			if err := i.fileUploader.UploadFile(headerFile.GetFile(), headerFile.GetFileName().GetObjectName()); err != nil {
				return err
			}
			headerIconNameChan <- headerFile.GetFileName().GetFileName()
			return nil
		})
	}

	// 並行処理の完了を待機
	if err := g.Wait(); err != nil {
		return UploadUserImageOutput{}, err
	}

	select {
	case iconName = <-iconNameChan:
	default:
		iconName = ""
	}

	select {
	case headerIconName = <-headerIconNameChan:
	default:
		headerIconName = ""
	}

	return i.presenter.Output(iconName, headerIconName), nil
}
