package repository

import (
	"context"
	"database/sql"
	"devport/domain/model"
	"devport/domain/repo/db"
	"devport/infra/database/sqlboiler/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type SqlBoilerWorkRepository struct {
	db *sql.DB
}

func NewSqlBoilerWorkRepository(db *sql.DB) *SqlBoilerWorkRepository {
	return &SqlBoilerWorkRepository{db: db}
}

func (r *SqlBoilerWorkRepository) Create(ctx context.Context, work *model.Work) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	model := r.convertToSqlBoilerModel(work)

	if !ok {
		return model.Insert(ctx, r.db, boil.Infer())
	}

	return model.Insert(ctx, tx, boil.Infer())
}

func (r *SqlBoilerWorkRepository) FindById(ctx context.Context, id string) (*db.WorkFindResponse, error) {
	work, err := models.FindWork(ctx, r.db, id)
	if err != nil {
		return nil, err
	}
	return &db.WorkFindResponse{
		ID:        work.ID,
		UserID:    work.UserID,
		Title:     work.Title,
		SortIndex: work.SortIndex,
		GitHubURL: work.GithubURL,
	}, nil
}

func (r *SqlBoilerWorkRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)
	if !ok {
		return fn(ctx)
	}
	return fn(context.WithValue(ctx, transactionContextKey, tx))
}

func (r *SqlBoilerWorkRepository) convertToSqlBoilerModel(work *model.Work) *models.Work {
	return &models.Work{
		ID:        work.ID(),
		UserID:    work.UserID(),
		Title:     work.Title(),
		SortIndex: work.SortIndex(),
		GithubURL: work.GithubURL(),
		CreatedAt: work.CreatedAt(),
		UpdatedAt: work.UpdatedAt(),
	}
}
