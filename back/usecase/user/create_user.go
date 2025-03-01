package user

import (
	"crypto/rand"
	"devport/domain/model"
	"devport/domain/repository/nosql"
	"devport/domain/repository/sql"
	"devport/infra/email"
	"devport/infra/security"
	"errors"
	"fmt"
	"math/big"
	"time"
)

type (
	CreateUserUseCase interface {
		Execute(CreateUserInput) (CreateUserOutput, error)
	}

	CreateUserInput struct {
		Age      int    `json:"age" validate:"required"`
		Name     string `json:"name" validate:"required,max=50,min=1"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=64"`
	}

	CreateUserOutput struct {
		Email string
	}

	createUserInterator struct {
		sqlRepository   sql.UserRepository
		noSqlRepository nosql.UserRepository
		email           email.Email
	}
)

func NewCreateUserInterator(
	sqlRepository sql.UserRepository,
	noSqlRepository nosql.UserRepository,
	email email.Email,
) CreateUserUseCase {
	return createUserInterator{
		sqlRepository:   sqlRepository,
		noSqlRepository: noSqlRepository,
		email:           email,
	}
}

func (i createUserInterator) Execute(input CreateUserInput) (CreateUserOutput, error) {
	userEmail, err := model.NewEmail(input.Email)

	if err != nil {
		return CreateUserOutput{""}, err
	}

	isExists, err := i.sqlRepository.Exists(userEmail) // ユーザが存在するか確認

	if err != nil {
		return CreateUserOutput{""}, err
	}
	if isExists {
		return CreateUserOutput{""}, errors.New("email already exists")
	}

	rawPassword, err := model.NewRawPassword(input.Password)

	if err != nil {
		return CreateUserOutput{""}, err
	}

	hashedPassword := security.HashPassword(rawPassword)

	user, err := model.NewUser(model.NewUUID(""), input.Name, input.Age, userEmail, hashedPassword, time.Now(), time.Now(), model.InConfirmation)

	if err != nil {
		return CreateUserOutput{""}, err
	}

	if err := i.sqlRepository.Create(user); err != nil {
		return CreateUserOutput{""}, err
	}

	// 6桁の確認コードを生成
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))

	if err != nil {
		return CreateUserOutput{""}, err
	}

	if err := i.noSqlRepository.AddConfirmationCode(userEmail, n.Int64()); err != nil {
		return CreateUserOutput{""}, err
	}

	mailSubject := "【ユーザ登録の認証コード送信のお知らせ】"

	mailContent := fmt.Sprintf("初回に登録されるすべてのユーザーに認証コードによるメール確認を行なっています。\n以下の数字を入力して認証を完了してください。\n認証コード:%d", n)

	if err := i.email.SendEmail([]string{input.Email}, mailSubject, mailContent); err != nil {
		return CreateUserOutput{""}, err
	}

	return CreateUserOutput{input.Email}, nil
}
