package user

import (
	usermodel "devport/domain/model"
	mock_nosql "devport/domain/repository/mock/nosql"
	mock_sql "devport/domain/repository/mock/sql"
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

	uc := NewCreateUserInterator(sqlMock, noSqlMock, emailMock)

	return uc, sqlMock, noSqlMock, emailMock
}

func TestCreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		i := CreateUserInput{
			Birthday:             "1990-01-01",
			Name:                 "test",
			Email:                "test@example.com",
			Password:             "Security_1234",
			PasswordConfirmation: "Security_1234",
		}

		testEmail, _ := usermodel.NewEmail(i.Email)

		uc, sqlMock, noSqlMock, emailMock := beforeAction(t, i)

		sqlMock.EXPECT().Create(gomock.Any()).Return(nil)
		sqlMock.EXPECT().Exists(testEmail).Return(false, nil)
		sqlMock.EXPECT().ExistsByName(i.Name).Return(false, nil)
		noSqlMock.EXPECT().AddConfirmationCode(gomock.Any(), gomock.Any()).Return(nil)
		emailMock.EXPECT().SendEmail([]string{i.Email}, gomock.Any(), gomock.Any())

		res, err := uc.Execute(i)

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
						Birthday:             "1990-01-01",
						Name:                 "test",
						Email:                c.email,
						Password:             "Security_1234",
						PasswordConfirmation: "Security_1234",
					}

					testEmail, _ := usermodel.NewEmail(i.Email)

					uc, sqlMock, _, _ := beforeAction(t, i)

					sqlMock.EXPECT().Exists(testEmail).Return(c.isExist, nil)

					_, err := uc.Execute(i)

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
						Birthday:             "1990-01-01",
						Name:                 c.name,
						Email:                "test@example.com",
						Password:             "Security_1234",
						PasswordConfirmation: "Security_1234",
					}

					uc, sqlMock, _, _ := beforeAction(t, i)

					sqlMock.EXPECT().Exists(gomock.Any()).Return(false, nil)
					sqlMock.EXPECT().ExistsByName(i.Name).Return(false, nil)

					_, err := uc.Execute(i)

					assert.Error(t, err)
				})
			}
		})

		t.Run("Password", func(t *testing.T) {
			cases := map[string]struct {
				password             string
				PasswordConfirmation string
			}{
				"empty": {
					password:             "",
					PasswordConfirmation: "",
				},
				"too short": {
					password:             "1234567",
					PasswordConfirmation: "1234567",
				},
				"wrong confirmation password": {
					password:             "Security_1234",
					PasswordConfirmation: "Security_12345",
				},
			}

			for name, c := range cases {
				t.Run(name, func(t *testing.T) {
					t.Parallel()
					i := CreateUserInput{
						Birthday:             "1990-01-01",
						Name:                 "test",
						Email:                "test@example.com",
						Password:             c.password,
						PasswordConfirmation: c.PasswordConfirmation,
					}

					uc, sqlMock, _, _ := beforeAction(t, i)

					sqlMock.EXPECT().Exists(gomock.Any()).Return(false, nil)
					sqlMock.EXPECT().ExistsByName(i.Name).Return(false, nil)

					_, err := uc.Execute(i)

					assert.Error(t, err)
				})
			}
		})
	})
}
