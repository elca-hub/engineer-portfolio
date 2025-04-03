package user

import (
	"context"
	"devport/domain/model"
	"devport/domain/repo/nosql"
	"devport/domain/repo/sql"
	"devport/infra/email"
	"errors"
	"time"
)

type (
	CreateUserUseCase interface {
		Execute(context.Context, CreateUserInput) (CreateUserOutput, error)
	}

	CreateUserInput struct {
		Birthday string `json:"birthday" validate:"required"`
		Name     string `json:"name" validate:"required,max=50,min=1"`
		Email    string `json:"email" validate:"required,email"`
		UserId   string `json:"user_id" validate:"required,max=50,min=1"`
	}

	CreateUserOutput struct {
		Email string `json:"email"`
	}

	CreateUserPresenter interface {
		Output(email string) CreateUserOutput
	}

	createUserInterator struct {
		sqlRepository   sql.UserRepository
		noSqlRepository nosql.UserRepository
		presenter       CreateUserPresenter
		email           email.Email
		ctxTimeout      time.Duration
	}
)

func NewCreateUserInterator(
	sqlRepository sql.UserRepository,
	noSqlRepository nosql.UserRepository,
	presenter CreateUserPresenter,
	email email.Email,
	t time.Duration,
) CreateUserUseCase {
	return createUserInterator{
		sqlRepository:   sqlRepository,
		noSqlRepository: noSqlRepository,
		presenter:       presenter,
		email:           email,
		ctxTimeout:      t,
	}
}

func (i createUserInterator) Execute(ctx context.Context, input CreateUserInput) (CreateUserOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	err := i.sqlRepository.WithTransaction(ctx, func(tx context.Context) error {
		e, err := model.NewEmail(input.Email)

		if err != nil {
			return err
		}

		isExistsMail, err := i.sqlRepository.Exists(tx, e)

		if err != nil {
			return err
		}

		if isExistsMail {
			return errors.New("メールアドレスは既に存在します")
		}

		isExistsName, err := i.sqlRepository.ExistsByName(tx, input.Name)

		if err != nil {
			return err
		}

		if isExistsName {
			return errors.New("ユーザ名は既に存在します")
		}

		isExistsId, err := i.sqlRepository.ExistsById(tx, input.UserId)

		if err != nil {
			return err
		}

		if isExistsId {
			return errors.New("IDはすでに存在しています")
		}

		jst, _ := time.LoadLocation("Asia/Tokyo")
		birthDay, err := time.ParseInLocation("2006-01-02", input.Birthday, jst)

		if err != nil {
			return err
		}

		user, err := model.NewUser(
			input.UserId,
			input.Name,
			birthDay,
			e,
			"",
			"",
			"",
			[]*model.Skill{},
			[]*model.ExternalServiceUrl{},
		)

		if err != nil {
			return err
		}

		if err := i.sqlRepository.Create(tx, user); err != nil {
			return err
		}

		mailObject := "【新規登録完了のお知らせ】"

		vars := map[string]string{"Name": user.Name()}
		files := []string{"infra/email/template/register.tpl"}

		if err := i.email.SendEmail(input.Email, mailObject, vars, files...); err != nil {
			return err
		}

		return nil
	})

	if err != nil {

		return i.presenter.Output(""), err
	}

	return i.presenter.Output(input.Email), nil
}
