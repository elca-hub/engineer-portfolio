package repository

import (
	"context"
	"database/sql"
	"devport/domain/model"
	"devport/infra/database/sqlboiler/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)


type SqlBoilerWorkImagesRepository struct {
	db *sql.DB
}

func NewSqlBoilerWorkImagesRepository(db *sql.DB) *SqlBoilerWorkImagesRepository {
	return &SqlBoilerWorkImagesRepository{db: db}
}

func (r *SqlBoilerWorkImagesRepository) Create(ctx context.Context, workId string, fileName string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	tar := &models.WorkImage{
		WorkID:   workId,
		FileName: fileName,
	}

	if !ok {
		return tar.Insert(ctx, r.db, boil.Infer())
	}

	return tar.Insert(ctx, tx, boil.Infer())
}

func (r *SqlBoilerWorkImagesRepository) Delete(ctx context.Context, work *model.Work, fileName string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	tar, err := models.WorkImages(models.WorkImageWhere.WorkID.EQ(work.ID()), models.WorkImageWhere.FileName.EQ(fileName)).One(ctx, r.db)

	if err != nil {
		return err
	}

	if !ok {
		if _, err := tar.Delete(ctx, r.db); err != nil {
			return err
		}

		return nil
	}

	if _, err := tar.Delete(ctx, tx); err != nil {
		return err
	}

	return nil
}


func (r *SqlBoilerWorkImagesRepository) FindByWorkId(ctx context.Context, workId string) ([]string, error) {
	fileNames, err := models.WorkImages(models.WorkImageWhere.WorkID.EQ(workId)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	res := make([]string, len(fileNames))
	for i, fileName := range fileNames {
		res[i] = fileName.FileName
	}
	return res, nil
}


func (r *SqlBoilerWorkImagesRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
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