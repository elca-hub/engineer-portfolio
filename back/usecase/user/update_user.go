package user

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/sql"
	"devport/infra/file_uploader"
	"mime/multipart"
	"time"

	"golang.org/x/sync/errgroup"
)

type (
	UpdateUserUseCase interface {
		Execute(context.Context, UpdateUserInput) (UpdateUserOutput, error)
	}

	UpdateUserInput struct {
		UserId              string `validate:"required"`
		Name                string `validate:"required,max=50"`
		Birthday            string `json:"birthday" validate:"required"`
		Email               string `validate:"required,email"`
		OrganizationName    string `validate:"max=50"`
		OccupationName      string `validate:"max=50"`
		Place               string `validate:"max=50"`
		Icon                multipart.File
		IconHeader          *multipart.FileHeader
		Header              multipart.File
		HeaderHeader        *multipart.FileHeader
		BioPath             string
		Skills              []*model.Skill
		ExternalServiceURLs []*model.ExternalServiceUrl
	}

	UpdateUserPresenter interface {
		Output(user *model.User) UpdateUserOutput
	}

	UpdateUserOutput struct {
		User *dto.UserDTO `json:"user"`
	}

	updateUserInterator struct {
		sqlRepository sql.UserRepository
		fileUploader  file_uploader.FileUploader
		presenter     UpdateUserPresenter
		ctxTimeout    time.Duration
	}
)

func NewUpdateUserInterator(
	sqlRepository sql.UserRepository,
	fileUploader file_uploader.FileUploader,
	presenter UpdateUserPresenter,
	t time.Duration,
) UpdateUserUseCase {
	return updateUserInterator{
		sqlRepository: sqlRepository,
		fileUploader:  fileUploader,
		presenter:     presenter,
		ctxTimeout:    t,
	}
}

func (i updateUserInterator) Execute(tx context.Context, input UpdateUserInput) (UpdateUserOutput, error) {
	ctx, cancel := context.WithTimeout(tx, i.ctxTimeout)
	defer cancel()

	user, err := i.sqlRepository.FindById(ctx, input.UserId)
	if err != nil {
		return UpdateUserOutput{}, err
	}

	email, err := model.NewEmail(input.Email)
	if err != nil {
		return UpdateUserOutput{}, err
	}

	err = i.sqlRepository.WithTransaction(ctx, func(tx context.Context) error {
		// 並行処理用のグループを作成
		g := new(errgroup.Group)

		var iconName string
		var headerIconName string

		iconNameChan := make(chan string, 1)
		headerIconNameChan := make(chan string, 1)

		// アイコン画像のアップロード
		if input.Icon != nil && input.IconHeader != nil {
			g.Go(func() error {
				iconFile, err := model.NewFileIcon(input.Icon, input.IconHeader, model.ICON_PATH)
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
				headerFile, err := model.NewFileIcon(input.Header, input.HeaderHeader, model.HEADER_PATH)
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
			return err
		}

		select {
		case iconName = <-iconNameChan:
		default:
			iconName = user.IconName()
		}

		select {
		case headerIconName = <-headerIconNameChan:
		default:
			headerIconName = user.HeaderIconName()
		}

		jst, _ := time.LoadLocation("Asia/Tokyo")
		birthDay, err := time.ParseInLocation("2006-01-02", input.Birthday, jst)

		if err != nil {
			return err
		}

		updatedUser, err := model.NewUser(
			user.ID(),
			input.Name,
			birthDay,
			email,
			iconName,
			headerIconName,
			input.BioPath,
			input.OrganizationName,
			input.OccupationName,
			input.Place,
			user.CreatedAt(),
			time.Now(),
			input.Skills,
			input.ExternalServiceURLs,
		)
		if err != nil {
			return err
		}

		return i.sqlRepository.Update(tx, updatedUser)
	})

	if err != nil {
		return UpdateUserOutput{}, err
	}

	return i.presenter.Output(user), nil
}
