package repository

import (
	"context"
	"database/sql"
	"devport/domain/model"
	"devport/infra/database/sqlboiler/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type SqlBoilerBioImagesRepository struct {
	db *sql.DB
}

func NewSqlBoilerBioImagesRepository(db *sql.DB) *SqlBoilerBioImagesRepository {
	return &SqlBoilerBioImagesRepository{db: db}
}

func (r *SqlBoilerBioImagesRepository) Create(ctx context.Context, userId string, fileName string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	tar := &models.BioImage{
		UserID:   userId,
		FileName: fileName,
	}

	if !ok {
		return tar.Insert(ctx, r.db, boil.Infer())
	}

	return tar.Insert(ctx, tx, boil.Infer())
}

func (r *SqlBoilerBioImagesRepository) Delete(ctx context.Context, user *model.User, fileName string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	tar, err := models.BioImages(models.BioImageWhere.UserID.EQ(user.ID()), models.BioImageWhere.FileName.EQ(fileName)).One(ctx, r.db)

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

func (r *SqlBoilerBioImagesRepository) DeleteAllByUserId(ctx context.Context, user *model.User) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	if !ok {
		fileNames, err := models.BioImages(models.BioImageWhere.UserID.EQ(user.ID())).All(ctx, r.db)
		if err != nil {
			return err
		}

		if _, err := fileNames.DeleteAll(ctx, r.db); err != nil {
			return err
		}
	}

	fileNames, err := models.BioImages(models.BioImageWhere.UserID.EQ(user.ID())).All(ctx, tx)
	if err != nil {
		return err
	}

	if _, err := fileNames.DeleteAll(ctx, r.db); err != nil {
		return err
	}

	return nil
}

func (r *SqlBoilerBioImagesRepository) FindByUserId(ctx context.Context, userId string) ([]string, error) {
	fileNames, err := models.BioImages(models.BioImageWhere.UserID.EQ(userId)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	res := make([]string, len(fileNames))
	for i, fileName := range fileNames {
		res[i] = fileName.FileName
	}
	return res, nil
}

func (r *SqlBoilerBioImagesRepository) IsExistsFileName(ctx context.Context, fileName string) (bool, error) {
	exists, err := models.BioImages(models.BioImageWhere.FileName.EQ(fileName)).Exists(ctx, r.db)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *SqlBoilerBioImagesRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
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
