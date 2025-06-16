package seed

import (
	"context"
	"database/sql"
	"devport/infra/database/sqlboiler/models"
	"fmt"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var mockUsers = []struct {
	userId   string
	name     string
	email    string
	birthday time.Time
}{
	{
		userId:   "mock_user_1",
		name:     "mock1",
		email:    "devport_mock_1@example.com",
		birthday: time.Date(2004, 11, 16, 0, 0, 0, 0, time.UTC),
	},
	{
		userId:   "mock_user_2",
		name:     "mock2",
		email:    "devport_mock_2@example.com",
		birthday: time.Date(2004, 11, 16, 0, 0, 0, 0, time.UTC),
	},
}

func CreateUserSeed(ctx context.Context, db *sql.DB) {
	for _, user := range mockUsers {
		exists, err := models.UserExists(ctx, db, user.userId)
		if err != nil {
			fmt.Errorf("error when checking if user exists: %v", err)
		}
		if exists {
			continue
		}

		user := models.User{
			ID:       user.userId,
			Name:     user.name,
			Email:    user.email,
			Birthday: user.birthday,
		}

		user.Insert(ctx, db, boil.Infer())
	}
}
