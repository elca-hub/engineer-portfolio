package repository

import (
	"context"
	"database/sql"
	"devport/domain/model"
	"devport/infra/database/sqlboiler/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SqlBoilerExternalServiceUrlRepository struct {
	db *sql.DB
}

func NewSqlBoilerExternalServiceUrlRepository(db *sql.DB) *SqlBoilerExternalServiceUrlRepository {
	return &SqlBoilerExternalServiceUrlRepository{db: db}
}

func (r *SqlBoilerExternalServiceUrlRepository) Create(ctx context.Context, user *model.User, externalServiceUrl *model.ExternalServiceUrl) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	model := r.convertToSqlBoilerModel(externalServiceUrl, user.ID())

	if !ok {
		return model.Insert(ctx, r.db, boil.Infer())
	}

	return model.Insert(ctx, tx, boil.Infer())
}

func (r *SqlBoilerExternalServiceUrlRepository) Update(ctx context.Context, user *model.User, externalServiceUrl *model.ExternalServiceUrl) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	model := r.convertToSqlBoilerModel(externalServiceUrl, user.ID())

	if !ok {
		_, err := model.Update(ctx, r.db, boil.Infer())
		return err
	}

	_, err := model.Update(ctx, tx, boil.Infer())
	return err
}

func (r *SqlBoilerExternalServiceUrlRepository) Delete(ctx context.Context, user *model.User, externalServiceUrl *model.ExternalServiceUrl) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	model := r.convertToSqlBoilerModel(externalServiceUrl, user.ID())

	if !ok {
		_, err := model.Delete(ctx, r.db)
		return err
	}

	_, err := model.Delete(ctx, tx)
	return err
}

func (r *SqlBoilerExternalServiceUrlRepository) FindByUserId(ctx context.Context, user *model.User) ([]*model.ExternalServiceUrl, error) {
	externalServiceUrls, err := models.ExternalServiceUrls(qm.Where("user_id = ?", user.ID())).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	res := make([]*model.ExternalServiceUrl, len(externalServiceUrls))

	for i, externalServiceUrl := range externalServiceUrls {
		res[i], err = r.convertToDomainModel(externalServiceUrl)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (r *SqlBoilerExternalServiceUrlRepository) FindByServiceType(ctx context.Context, user *model.User, serviceType int) (*model.ExternalServiceUrl, error) {
	externalServiceUrl, err := models.ExternalServiceUrls(models.ExternalServiceURLWhere.UserID.EQ(user.ID()), models.ExternalServiceURLWhere.ServiceType.EQ(serviceType)).One(ctx, r.db)

	if err != nil {
		return nil, err
	}

	return r.convertToDomainModel(externalServiceUrl)
}

func (r *SqlBoilerExternalServiceUrlRepository) IsExistsByServiceType(ctx context.Context, user *model.User, serviceType int) (bool, error) {
	exists, err := models.ExternalServiceUrls(models.ExternalServiceURLWhere.UserID.EQ(user.ID()), models.ExternalServiceURLWhere.ServiceType.EQ(serviceType)).Exists(ctx, r.db)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *SqlBoilerExternalServiceUrlRepository) FindById(ctx context.Context, id string) (*model.ExternalServiceUrl, error) {
	externalServiceUrl, err := models.ExternalServiceUrls(qm.Where("id = ?", id)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return r.convertToDomainModel(externalServiceUrl)
}

func (r *SqlBoilerExternalServiceUrlRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)
	if !ok {
		return fn(ctx)
	}
	return fn(context.WithValue(ctx, transactionContextKey, tx))
}

func (r *SqlBoilerExternalServiceUrlRepository) convertToSqlBoilerModel(externalServiceUrl *model.ExternalServiceUrl, userId string) *models.ExternalServiceURL {
	return &models.ExternalServiceURL{
		ID:          externalServiceUrl.ID(),
		UserID:      userId,
		ServiceType: externalServiceUrl.ServiceType(),
		URL:         externalServiceUrl.Url(),
	}
}

func (r *SqlBoilerExternalServiceUrlRepository) convertToDomainModel(externalServiceUrl *models.ExternalServiceURL) (*model.ExternalServiceUrl, error) {
	model, err := model.NewExternalServiceUrl(externalServiceUrl.ID, externalServiceUrl.ServiceType, externalServiceUrl.URL)
	if err != nil {
		return nil, err
	}

	return model, nil
}
