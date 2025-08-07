package repository

import (
	"context"
	"database/sql"
	"devport/domain/model"
	"devport/infra/database/sqlboiler/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type SqlboilerWorkUrlRepository struct {
	db *sql.DB
}

func NewSqlboilerWorkUrlRepository(db *sql.DB) *SqlboilerWorkUrlRepository {
	return &SqlboilerWorkUrlRepository{
		db: db,
	}
}

func (r *SqlboilerWorkUrlRepository) CreateByWorkId(ctx context.Context, workId string, workUrl []*model.WorkUrl) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	sqlboilerWorkUrls := make([]*models.WorkURL, len(workUrl))
	for i, url := range workUrl {
		sqlboilerWorkUrls[i] = r.convertToSqlBoilerModel(url, workId)
	}

	if !ok {
		for _, url := range sqlboilerWorkUrls {
			err := url.Insert(ctx, r.db, boil.Infer())
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, url := range sqlboilerWorkUrls {
		err := url.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *SqlboilerWorkUrlRepository) DeleteByWorkId(ctx context.Context, workId string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	if !ok {
		_, err := models.WorkUrls(models.WorkURLWhere.WorkID.EQ(workId)).DeleteAll(ctx, r.db)
		return err
	}

	_, err := models.WorkUrls(models.WorkURLWhere.WorkID.EQ(workId)).DeleteAll(ctx, tx)
	return err
}

func (r *SqlboilerWorkUrlRepository) Delete(ctx context.Context, id string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	if !ok {
		_, err := models.WorkUrls(models.WorkURLWhere.ID.EQ(id)).DeleteAll(ctx, r.db)
		return err
	}

	_, err := models.WorkUrls(models.WorkURLWhere.ID.EQ(id)).DeleteAll(ctx, tx)
	return err
}

func (r *SqlboilerWorkUrlRepository) convertToSqlBoilerModel(workUrl *model.WorkUrl, workId string) *models.WorkURL {
	return &models.WorkURL{
		ID:     workUrl.ID(),
		URL:    workUrl.Url(),
		Title:  workUrl.Title(),
		WorkID: workId,
	}
}

func (r *SqlboilerWorkUrlRepository) convertToDomainModel(ctx context.Context, sqlboilerWorkUrl *models.WorkURL) (*model.WorkUrl, error) {
	return model.NewWorkUrl(
		sqlboilerWorkUrl.ID,
		sqlboilerWorkUrl.URL,
		sqlboilerWorkUrl.Title,
	)
}

func (r *SqlboilerWorkUrlRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	if err := fn(context.WithValue(ctx, transactionContextKey, tx)); err != nil {
		return err
	}

	committed = true

	return tx.Commit()
}
