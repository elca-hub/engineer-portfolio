package seed

import (
	"context"
	"database/sql"
	"devport/infra/database/sqlboiler/models"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func CreateUserSeed(ctx context.Context, db *sql.DB) {
	userId := "mock_user_id"
	exists, err := models.UserExists(ctx, db, userId)
	if err != nil {
		panic(err)
	}

	if exists {
		return
	}

	user := models.User{
		ID:       userId,
		Name:     "mock mock",
		Email:    "devport_mock@example.com",
		Birthday: time.Date(2004, 11, 16, 0, 0, 0, 0, time.UTC),
	}

	user.Insert(ctx, db, boil.Infer())
}
