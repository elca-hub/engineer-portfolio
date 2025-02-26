package user

import (
	usermodel "devport/domain/model"
	mock_nosql "devport/domain/repository/mock/nosql"
	mock_sql "devport/domain/repository/mock/sql"
	mock_email "devport/infra/mock/email"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func beforeAction(t *testing.T, i CreateUserInput) (
	CreateUserUseCase,
	*mock_sql.MockUserRepository,
	*mock_nosql.MockUserRepository,
	*mock_email.MockEmail,
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sqlMock := mock_sql.NewMockUserRepository(ctrl)
	noSqlMock := mock_nosql.NewMockUserRepository(ctrl)
	emailMock := mock_email.NewMockEmail(ctrl)

	uc := NewCreateUserInterator(sqlMock, noSqlMock, emailMock)

	return uc, sqlMock, noSqlMock, emailMock
}

func TestCreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		i := CreateUserInput{
			Age:      18,
			Name:     "test",
			Email:    "test@example.com",
			Password: "security",
		}

		testEmail, _ := usermodel.NewEmail(i.Email)

		uc, sqlMock, noSqlMock, emailMock := beforeAction(t, i)

		sqlMock.EXPECT().Create(gomock.Any()).Return(nil)
		sqlMock.EXPECT().Exists(testEmail).Return(false, nil)
		noSqlMock.EXPECT().StartSession(testEmail)
		emailMock.EXPECT().SendEmail([]string{i.Email}, gomock.Any(), gomock.Any())

		res, err := uc.Execute(i)

		assert.NoError(t, err)

		assert.Equal(t, i.Email, res.Email)
	})

	})
}
