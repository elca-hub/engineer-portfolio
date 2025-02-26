package user

import (
	"devport/domain/model"
	mock_nosql "devport/domain/repository/mock/nosql"
	mock_sql "devport/domain/repository/mock/sql"
	mock_email "devport/infra/mock/email"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		testEmail, _ := model.NewEmail("test@example.com")
		testModel, _ := model.NewUser(
			model.NewUUID(""),
			"test",
			20,
			testEmail,
			"password0123",
			time.Now(),
			time.Now(),
			model.Unconfirmed,
		)

		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		sqlMock := mock_sql.NewMockUserRepository(ctrl)
		noSqlMock := mock_nosql.NewMockUserRepository(ctrl)
		emailMock := mock_email.NewMockEmail(ctrl)

		sqlMock.EXPECT().Create(gomock.Any()).Return(nil)
		sqlMock.EXPECT().Exists(testEmail).Return(false, nil)
		noSqlMock.EXPECT().StartSession(testEmail)
		emailMock.EXPECT().SendEmail([]string{testEmail.Email()}, "【メール確認のお願い】", gomock.Any())

		interator := NewCreateUserInterator(sqlMock, noSqlMock, emailMock)

		input := CreateUserInput{
			Age:      testModel.Age(),
			Name:     testModel.Name(),
			Email:    testEmail.Email(),
			Password: testModel.Password(),
		}

		output, err := interator.Execute(input)

		assert.NoError(t, err)
		assert.Equal(t, output.Email, testEmail.Email())
	})
}
