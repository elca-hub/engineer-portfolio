package user

import (
	"context"
	"devport/domain/model"
	"devport/domain/repo/sql"
	"devport/infra/file_uploader"
	"sync"
	"time"
)

type (
	DeleteUserImageUseCase interface {
		Execute(context.Context, DeleteUserImageInput) (DeleteUserImageOutput, error)
	}

	DeleteUserImageInput struct {
	}

	DeleteUserImagePresenter interface {
		Output() DeleteUserImageOutput
	}

	DeleteUserImageOutput struct {
	}

	deleteUserImageInterator struct {
		sqlRepository sql.UserRepository
		fileUploader  file_uploader.FileUploader
		presenter     UploadUserImagePresenter
		ctxTimeout    time.Duration
	}
)

func NewDeleteUserImageInterator(
	sqlRepository sql.UserRepository,
	fileUploader file_uploader.FileUploader,
	t time.Duration,
) DeleteUserImageUseCase {
	return deleteUserImageInterator{
		sqlRepository: sqlRepository,
		fileUploader:  fileUploader,
		ctxTimeout:    t,
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func (i deleteUserImageInterator) Execute(tx context.Context, input DeleteUserImageInput) (DeleteUserImageOutput, error) {
	var (
		wg                                         sync.WaitGroup
		uploadedIconImageNames                     []*model.FileIconName
		uploadedHeaderImageNames                   []*model.FileIconName
		dbIconNames                                []*model.FileIconName
		dbHeaderNames                              []*model.FileIconName
		iconErr, headerErr, dbIconErr, dbHeaderErr error
	)

	wg.Add(4)
	// アップロード済みのアイコン取得
	go func() {
		defer wg.Done()
		uploadedIconImageNames, iconErr = i.fileUploader.GetFiles(model.ICON_PATH)
	}()

	// アップロード済みのヘッダー取得
	go func() {
		defer wg.Done()
		uploadedHeaderImageNames, headerErr = i.fileUploader.GetFiles(model.HEADER_PATH)
	}()

	// DBに登録されているアイコン取得
	go func() {
		defer wg.Done()
		dbIconNames, dbIconErr = i.sqlRepository.FetchIconNamesAll(tx)
	}()

	// DBに登録されているヘッダー取得
	go func() {
		defer wg.Done()
		dbHeaderNames, dbHeaderErr = i.sqlRepository.FetchHeaderNamesAll(tx)
	}()

	wg.Wait()

	// エラーチェック
	if iconErr != nil {
		return DeleteUserImageOutput{}, iconErr
	}
	if headerErr != nil {
		return DeleteUserImageOutput{}, headerErr
	}
	if dbIconErr != nil {
		return DeleteUserImageOutput{}, dbIconErr
	}
	if dbHeaderErr != nil {
		return DeleteUserImageOutput{}, dbHeaderErr
	}

	dbIconNamesStrings := make([]string, 0)
	for _, v := range dbIconNames {
		dbIconNamesStrings = append(dbIconNamesStrings, v.GetFileName())
	}

	dbHeaderNamesStrings := make([]string, 0)
	for _, v := range dbHeaderNames {
		dbHeaderNamesStrings = append(dbHeaderNamesStrings, v.GetFileName())
	}

	// 削除する画像名を取得
	deletedIconImageNames := make([]string, 0)
	deletedHeaderImageNames := make([]string, 0)

	for _, iconName := range uploadedIconImageNames {
		if !contains(dbIconNamesStrings, iconName.GetFileName()) {
			deletedIconImageNames = append(deletedIconImageNames, iconName.GetObjectName())
		}
	}

	for _, headerName := range uploadedHeaderImageNames {
		if !contains(dbHeaderNamesStrings, headerName.GetFileName()) {
			deletedHeaderImageNames = append(deletedHeaderImageNames, headerName.GetObjectName())
		}
	}

	errChan := make(chan error, len(deletedIconImageNames)+len(deletedHeaderImageNames))

	for _, iconName := range deletedIconImageNames {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			if err := i.fileUploader.DeleteFile(name); err != nil {
				errChan <- err
			}
		}(iconName)
	}

	for _, headerName := range deletedHeaderImageNames {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			if err := i.fileUploader.DeleteFile(name); err != nil {
				errChan <- err
			}
		}(headerName)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return DeleteUserImageOutput{}, err
		}
	}

	return DeleteUserImageOutput{}, nil
}
