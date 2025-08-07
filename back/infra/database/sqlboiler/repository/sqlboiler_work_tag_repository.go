package repository

import (
	"context"
	"database/sql"
	"devport/domain/model"
	"devport/infra/database/sqlboiler/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type SqlboilerWorkTagRepository struct {
	db *sql.DB
}

func NewSqlboilerWorkTagRepository(db *sql.DB) *SqlboilerWorkTagRepository {
	return &SqlboilerWorkTagRepository{
		db: db,
	}
}

func (r *SqlboilerWorkTagRepository) Create(ctx context.Context, workTag *model.WorkTag) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	sqlboilerWorkTag := r.convertToSqlBoilerModel(workTag)

	if !ok {
		return sqlboilerWorkTag.Insert(ctx, r.db, boil.Infer())
	}

	return sqlboilerWorkTag.Insert(ctx, tx, boil.Infer())
}

func (r *SqlboilerWorkTagRepository) Exists(ctx context.Context, name string) (bool, error) {
	return models.WorkTags(models.WorkTagWhere.Name.EQ(name)).Exists(ctx, r.db)
}

func (r *SqlboilerWorkTagRepository) Update(ctx context.Context, workTag *model.WorkTag) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	sqlboilerWorkTag := r.convertToSqlBoilerModel(workTag)

	if !ok {
		_, err := sqlboilerWorkTag.Update(ctx, r.db, boil.Infer())
		return err
	}

	_, err := sqlboilerWorkTag.Update(ctx, tx, boil.Infer())
	return err
}

func (r *SqlboilerWorkTagRepository) FindByName(ctx context.Context, name string) (*model.WorkTag, error) {
	sqlboilerWorkTag, err := models.WorkTags(models.WorkTagWhere.Name.EQ(name)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return r.convertToDomainModel(ctx, sqlboilerWorkTag)
}

func (r *SqlboilerWorkTagRepository) FindById(ctx context.Context, id string) (*model.WorkTag, error) {
	sqlboilerWorkTag, err := models.WorkTags(models.WorkTagWhere.ID.EQ(id)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return r.convertToDomainModel(ctx, sqlboilerWorkTag)
}

func (r *SqlboilerWorkTagRepository) Delete(ctx context.Context, id string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	if !ok {
		_, err := models.WorkTags(models.WorkTagWhere.ID.EQ(id)).DeleteAll(ctx, r.db)
		return err
	}

	_, err := models.WorkTags(models.WorkTagWhere.ID.EQ(id)).DeleteAll(ctx, tx)
	return err
}

func (r *SqlboilerWorkTagRepository) DeleteNoUsed(ctx context.Context) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	query := `
		SELECT t.id
		FROM work_having_tags w
		LEFT OUTER JOIN work_tags t ON w.work_tag_id <> t.id
	`

	var workIds []string

	// SQLを実行
	if ok {
		rows, err := tx.QueryContext(ctx, query)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var workId string

			if err := rows.Scan(&workId); err != nil {
				continue
			}
			workIds = append(workIds, workId)
		}

		if err := rows.Err(); err != nil {
			return err
		}
	} else {
		rows, err := r.db.QueryContext(ctx, query)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var workId string
			if err := rows.Scan(&workId); err != nil {
				return err
			}
			workIds = append(workIds, workId)
		}

		if err := rows.Err(); err != nil {
			return err
		}
	}

	// 抽出されたIDに対応するworksを削除
	if len(workIds) > 0 {
		if ok {
			_, err := models.WorkTags(models.WorkTagWhere.ID.IN(workIds)).DeleteAll(ctx, tx)

			if err != nil {
				return err
			}
		} else {
			_, err := models.WorkTags(models.WorkTagWhere.ID.IN(workIds)).DeleteAll(ctx, r.db)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *SqlboilerWorkTagRepository) convertToSqlBoilerModel(workTag *model.WorkTag) *models.WorkTag {
	return &models.WorkTag{
		ID:   workTag.ID(),
		Name: workTag.Name(),
	}
}

func (r *SqlboilerWorkTagRepository) convertToDomainModel(ctx context.Context, sqlboilerWorkTag *models.WorkTag) (*model.WorkTag, error) {
	return model.NewWorkTag(
		sqlboilerWorkTag.ID,
		sqlboilerWorkTag.Name,
	)
}

func (r *SqlboilerWorkTagRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
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
