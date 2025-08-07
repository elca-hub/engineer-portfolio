package repository

import (
	"context"
	"database/sql"
	"devport/domain/model"
	"devport/infra/database/sqlboiler/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SqlboilerWorkRepository struct {
	db *sql.DB
}

func NewSqlboilerWorkRepository(db *sql.DB) *SqlboilerWorkRepository {
	return &SqlboilerWorkRepository{
		db: db,
	}
}

func (r *SqlboilerWorkRepository) Create(ctx context.Context, userId string, work *model.Work) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	sqlboilerWork := r.convertToSqlBoilerModel(work, userId)

	if !ok {
		return sqlboilerWork.Insert(ctx, r.db, boil.Infer())
	}

	return sqlboilerWork.Insert(ctx, tx, boil.Infer())
}

func (r *SqlboilerWorkRepository) Update(ctx context.Context, userId string, work *model.Work) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	sqlboilerWork := r.convertToSqlBoilerModel(work, userId)

	if !ok {
		_, err := sqlboilerWork.Update(ctx, r.db, boil.Infer())
		return err
	}

	_, err := sqlboilerWork.Update(ctx, tx, boil.Infer())
	return err
}

func (r *SqlboilerWorkRepository) GetMaxSortIndex(ctx context.Context, userId string) (int, error) {
	works, err := models.Works(
		models.WorkWhere.UserID.EQ(userId),
		qm.OrderBy("sort_index DESC"),
		qm.Limit(1),
	).All(ctx, r.db)
	if err != nil {
		return 0, err
	}

	if len(works) == 0 {
		return 0, nil
	}

	return works[0].SortIndex, nil
}

func (r *SqlboilerWorkRepository) FindById(ctx context.Context, userId string, id string) (*model.Work, error) {
	sqlboilerWork, err := models.Works(models.WorkWhere.ID.EQ(id), models.WorkWhere.UserID.EQ(userId)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return r.convertToDomainModel(ctx, sqlboilerWork)
}

func (r *SqlboilerWorkRepository) FindAll(ctx context.Context, userId string) ([]*model.Work, error) {
	sqlboilerWorks, err := models.Works(models.WorkWhere.UserID.EQ(userId)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	worksDomain := make([]*model.Work, len(sqlboilerWorks))
	for i, work := range sqlboilerWorks {
		worksDomain[i], err = r.convertToDomainModel(ctx, work)
		if err != nil {
			return nil, err
		}
	}
	return worksDomain, nil
}

func (r *SqlboilerWorkRepository) convertToDomainModel(ctx context.Context, sqlboilerWork *models.Work) (*model.Work, error) {
	workUrls, err := models.WorkUrls(models.WorkURLWhere.WorkID.EQ(sqlboilerWork.ID)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	workUrlsDomain := make([]*model.WorkUrl, len(workUrls))

	if workUrls != nil {
		var err error

		for i, url := range workUrls {
			workUrlsDomain[i], err = model.NewWorkUrl(url.ID, url.URL, url.Title)
			if err != nil {
				return nil, err
			}
		}
	}

	workTags, err := models.WorkHavingTags(models.WorkHavingTagWhere.WorkID.EQ(sqlboilerWork.ID)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	workTagsDomain := make([]string, len(workTags))
	for i, tag := range workTags {
		workTag, err := models.FindWorkTag(ctx, r.db, tag.WorkTagID)
		if err != nil {
			return nil, err
		}
		workTagsDomain[i] = workTag.Name
	}

	return model.NewWork(
		sqlboilerWork.ID,
		sqlboilerWork.Title,
		sqlboilerWork.Content,
		sqlboilerWork.GithubRepositoryURL,
		workUrlsDomain,
		workTagsDomain,
		sqlboilerWork.IsDraft,
		sqlboilerWork.SortIndex,
	)
}

func (r *SqlboilerWorkRepository) convertToSqlBoilerModel(work *model.Work, userId string) *models.Work {
	return &models.Work{
		ID:                  work.ID(),
		Title:               work.Title(),
		Content:             work.Content(),
		SortIndex:           work.SortIndex(),
		UserID:              userId,
		GithubRepositoryURL: work.GithubRepositoryUrl(),
	}
}

func (r *SqlboilerWorkRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
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
