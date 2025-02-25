package user

import (
	"devport/domain/model"
	"devport/domain/repository/nosql"
	"devport/domain/repository/sql"
	"devport/infra/email"
	"devport/infra/password"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type (
	CreateUserUseCase interface {
		Execute(CreateUserInput) (CreateUserOutput, error)
	}

	CreateUserInput struct {
		Age      int    `json:"age" validate:"required"`
		Name     string `json:"name" validate:"required,max=50"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=64"`
	}

	CreateUserOutput struct {
		Email string
	}

	createUserInterator struct {
		sqlRepository   sql.UserRepository
		noSqlRepository nosql.UserRepository
	}
)

func NewCreateUserInterator(
	sqlRepository sql.UserRepository,
	noSqlRepository nosql.UserRepository,
) CreateUserUseCase {
	return createUserInterator{
		sqlRepository:   sqlRepository,
		noSqlRepository: noSqlRepository,
	}
}

func (i createUserInterator) Execute(input CreateUserInput) (CreateUserOutput, error) {
	hashedPw := password.HashPassword(input.Password)

	validate := validator.New()

	if err := validate.Struct(input); err != nil {
		return CreateUserOutput{""}, err
	}

	userEmail, err := model.NewEmail(input.Email)

	if err != nil {
		return CreateUserOutput{""}, err
	}

	isExists, err := i.sqlRepository.Exists(userEmail) // ユーザが存在するか確認

	if err != nil {
		return CreateUserOutput{""}, err
	}
	if isExists {
		return CreateUserOutput{""}, errors.New("user_presenter already exists")
	}

	user, err := model.NewUser(model.NewUUID(""), input.Name, input.Age, userEmail, hashedPw, time.Now(), time.Now(), model.InConfirmation)

	if err != nil {
		return CreateUserOutput{""}, err
	}

	if err := i.sqlRepository.Create(user); err != nil {
		return CreateUserOutput{""}, err
	}

	token, err := i.noSqlRepository.StartSession(userEmail)

	if err != nil {
		return CreateUserOutput{""}, err
	}

	mailSubject := "【メール確認のお願い】"

	mailContent := fmt.Sprintf("以下のリンクをクリックしてメールアドレスを確認してください。\nhttp://localhost:8080/verification/email?token=%s", token)

	if err := email.SmtpSendMail([]string{input.Email}, mailSubject, mailContent); err != nil {
		return CreateUserOutput{""}, err
	}

	return CreateUserOutput{input.Email}, nil
}
