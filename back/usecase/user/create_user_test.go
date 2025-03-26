package user

import (
	usermodel "devport/domain/model"
	mock_nosql "devport/domain/repo/mock/nosql"
	mock_sql "devport/domain/repo/mock/sql"
	mock_email "devport/infra/mock/email"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
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

	uc := NewCreateUserInterator(sqlMock, noSqlMock, emailMock, 5)

	return uc, sqlMock, noSqlMock, emailMock
}

func TestCreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		i := CreateUserInput{
			Birthday: "1990-01-01",
			Name:     "test",
			Email:    "test@example.com",
		}

		testEmail, _ := usermodel.NewEmail(i.Email)

		uc, sqlMock, noSqlMock, emailMock := beforeAction(t, i)

		sqlMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
		sqlMock.EXPECT().Exists(gomock.Any(), testEmail).Return(false, nil)
		sqlMock.EXPECT().ExistsByName(gomock.Any(), i.Name).Return(false, nil)
		sqlMock.EXPECT().WithTransaction(gomock.Any(), gomock.Any()).Return(nil)
		noSqlMock.EXPECT().AddConfirmationCode(gomock.Any(), gomock.Any()).Return(nil)
		emailMock.EXPECT().SendEmail([]string{i.Email}, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())

		res, err := uc.Execute(t.Context(), i)

		assert.NoError(t, err)

		assert.Equal(t, i.Email, res.Email)
	})

	t.Run("failures", func(t *testing.T) {
		t.Run("Email", func(t *testing.T) {
			cases := map[string]struct {
				email   string
				isExist bool
			}{
				"empty": {
					email:   "",
					isExist: false,
				},
				"invalid": {
					email:   "test",
					isExist: false,
				},
				"already exists": {
					email:   "test@example.com",
					isExist: true,
				},
			}

			for name, c := range cases {
				t.Run(name, func(t *testing.T) {
					t.Parallel()
					i := CreateUserInput{
						Birthday: "1990-01-01",
						Name:     "test",
						Email:    c.email,
					}

					testEmail, _ := usermodel.NewEmail(i.Email)

					uc, sqlMock, _, _ := beforeAction(t, i)

					sqlMock.EXPECT().Exists(gomock.Any(), testEmail).Return(c.isExist, nil)
					sqlMock.EXPECT().WithTransaction(gomock.Any(), gomock.Any()).Return(nil)

					_, err := uc.Execute(t.Context(), i)

					assert.Error(t, err)
				})
			}
		})

		t.Run("UserName", func(t *testing.T) {
			tooLongName := "a"
			for range usermodel.MaxNameLen {
				tooLongName += "a"
			}

			cases := map[string]struct {
				name string
			}{
				"empty": {
					name: "",
				},
				"too long": {
					name: tooLongName,
				},
				"special characters": {
					name: "test@",
				},
			}

			for name, c := range cases {
				t.Run(name, func(t *testing.T) {
					t.Parallel()
					i := CreateUserInput{
						Birthday: "1990-01-01",
						Name:     c.name,
						Email:    "test@example.com",
					}

					uc, sqlMock, _, _ := beforeAction(t, i)

					sqlMock.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(false, nil)
					sqlMock.EXPECT().ExistsByName(gomock.Any(), i.Name).Return(false, nil)
					sqlMock.EXPECT().WithTransaction(gomock.Any(), gomock.Any()).Return(nil)

					_, err := uc.Execute(t.Context(), i)

					assert.Error(t, err)
				})
			}
		})
	})
}
