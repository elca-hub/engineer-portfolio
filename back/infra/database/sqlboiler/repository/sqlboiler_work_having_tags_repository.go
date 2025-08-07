package repository

import (
	"context"
	"database/sql"
	"devport/domain/model"
	"devport/infra/database/sqlboiler/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type SqlboilerWorkHavingTagsRepository struct {
	db *sql.DB
}

func NewSqlboilerWorkHavingTagsRepository(db *sql.DB) *SqlboilerWorkHavingTagsRepository {
	return &SqlboilerWorkHavingTagsRepository{
		db: db,
	}
}

func (r *SqlboilerWorkHavingTagsRepository) DeleteByWorkId(ctx context.Context, workId string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	if !ok {
		_, err := models.WorkHavingTags(
			models.WorkHavingTagWhere.WorkID.EQ(workId),
		).DeleteAll(ctx, r.db)
		return err
	}

	_, err := models.WorkHavingTags(
		models.WorkHavingTagWhere.WorkID.EQ(workId),
	).DeleteAll(ctx, tx)
	return err
}

func (r *SqlboilerWorkHavingTagsRepository) CreateByWorkId(ctx context.Context, workId string, tagIds []string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	workHavingTags := make([]*models.WorkHavingTag, len(tagIds))
	for i, tagId := range tagIds {
		workHavingTags[i] = &models.WorkHavingTag{
			WorkID:    workId,
			WorkTagID: tagId,
		}
	}

	// TODO: n+1問題
	if !ok {
		for _, workHavingTag := range workHavingTags {
			err := workHavingTag.Insert(ctx, r.db, boil.Infer())
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, workHavingTag := range workHavingTags {
		err := workHavingTag.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *SqlboilerWorkHavingTagsRepository) Find(ctx context.Context, workId string) ([]*model.WorkTag, error) {
	workHavingTags, err := models.WorkHavingTags(
		models.WorkHavingTagWhere.WorkID.EQ(workId),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return r.convertToDomainModel(ctx, workHavingTags)
}

func (r *SqlboilerWorkHavingTagsRepository) UpdateByWorkId(ctx context.Context, workId string, tagIds []string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	if !ok {
		for _, tagId := range tagIds {
			tagModel := &models.WorkHavingTag{
				WorkID:    workId,
				WorkTagID: tagId,
			}

			_, err := tagModel.Update(ctx, r.db, boil.Infer())

			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, tagId := range tagIds {
		tagModel := &models.WorkHavingTag{
			WorkID:    workId,
			WorkTagID: tagId,
		}

		_, err := tagModel.Update(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *SqlboilerWorkHavingTagsRepository) convertToDomainModel(ctx context.Context, workHavingTags []*models.WorkHavingTag) ([]*model.WorkTag, error) {
	workTags := make([]*model.WorkTag, len(workHavingTags))
	for i, workHavingTag := range workHavingTags {
		workTag, err := workHavingTag.WorkTag().One(ctx, r.db)
		if err != nil {
			return nil, err
		}
		workTags[i], err = model.NewWorkTag(
			workTag.ID,
			workTag.Name,
		)
		if err != nil {
			return nil, err
		}
	}

	return workTags, nil
}

func (r *SqlboilerWorkHavingTagsRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
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
