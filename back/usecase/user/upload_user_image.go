package user

import (
	"context"
	"devport/domain/repo/db"
	"devport/domain/repo/file_storage"
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
		UserId       string `json:"user_id" validate:"required,max=50,min=1"`
	}

	UploadUserImagePresenter interface {
		Output(iconName string, headerIconName string) UploadUserImageOutput
	}

	UploadUserImageOutput struct {
		IconName       string `json:"icon_name"`
		HeaderIconName string `json:"header_icon_name"`
	}

	uploadUserImageInterator struct {
		userRepository db.UserRepository
		iconStorage    file_storage.IconImageStorage
		headerStorage  file_storage.HeaderImageStorage
		presenter      UploadUserImagePresenter
		ctxTimeout     time.Duration
	}
)

func NewUploadUserImageInterator(
	userRepository db.UserRepository,
	iconStorage file_storage.IconImageStorage,
	headerStorage file_storage.HeaderImageStorage,
	presenter UploadUserImagePresenter,
	t time.Duration,
) UploadUserImageUseCase {
	return uploadUserImageInterator{
		userRepository: userRepository,
		iconStorage:    iconStorage,
		headerStorage:  headerStorage,
		presenter:      presenter,
		ctxTimeout:     t,
	}
}

func (i uploadUserImageInterator) Execute(tx context.Context, input UploadUserImageInput) (UploadUserImageOutput, error) {
	user, err := i.userRepository.FindById(tx, input.UserId, nil)

	if err != nil {
		return UploadUserImageOutput{}, err
	}

	// 並行処理用のグループを作成
	g := new(errgroup.Group)

	// バッファ付きチャネルを作成
	iconNameChan := make(chan string, 1)
	headerIconNameChan := make(chan string, 1)

	// アイコン画像のアップロード
	if input.Icon != nil && input.IconHeader != nil {
		g.Go(func() error {
			defer close(iconNameChan)
			// 既にアップロードしている場合はその画像を削除
			if user.IconName() != "" {
				if err := i.iconStorage.Delete(tx, user.IconName()); err != nil {
					return err
				}
			}

			newFileName, err := i.iconStorage.Upload(tx, input.Icon, input.IconHeader)
			if err != nil {
				return err
			}

			iconNameChan <- newFileName
			return nil
		})
	} else {
		close(iconNameChan)
	}

	// ヘッダー画像のアップロード
	if input.Header != nil && input.HeaderHeader != nil {
		g.Go(func() error {
			defer close(headerIconNameChan)
			if user.HeaderIconName() != "" {
				if err := i.headerStorage.Delete(tx, user.HeaderIconName()); err != nil {
					return err
				}
			}

			newFileName, err := i.headerStorage.Upload(tx, input.Header, input.HeaderHeader)
			if err != nil {
				return err
			}

			headerIconNameChan <- newFileName
			return nil
		})
	} else {
		close(headerIconNameChan)
	}

	// 並行処理の完了を待機
	if err := g.Wait(); err != nil {
		return UploadUserImageOutput{}, err
	}

	var newIconName, newHeaderIconName string

	// チャネルからの読み取りを待機
	if input.Icon != nil && input.IconHeader != nil {
		if name, ok := <-iconNameChan; ok {
			newIconName = name
		} else {
			newIconName = user.IconName()
		}
	} else {
		newIconName = user.IconName()
	}

	if input.Header != nil && input.HeaderHeader != nil {
		if name, ok := <-headerIconNameChan; ok {
			newHeaderIconName = name
		} else {
			newHeaderIconName = user.HeaderIconName()
		}
	} else {
		newHeaderIconName = user.HeaderIconName()
	}

	user.UpdateIconName(newIconName)
	user.UpdateHeaderIconName(newHeaderIconName)

	if err := i.userRepository.Update(tx, user); err != nil {
		return UploadUserImageOutput{}, err
	}

	return i.presenter.Output(newIconName, newHeaderIconName), nil
}
