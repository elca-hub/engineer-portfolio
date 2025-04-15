package repository

import (
	"context"
	"devport/adapter/repository"
	"devport/domain/model"
	"devport/infra/database/gorm/gorm_model"

	"gorm.io/gorm"
)

type GormExternalServiceRepository struct {
	db repository.SQL
}

func NewGormExternalServiceRepository(db repository.SQL) *GormExternalServiceRepository {
	return &GormExternalServiceRepository{
		db: db,
	}
}

func (r GormExternalServiceRepository) Create(ctx context.Context, user *model.User, externalServiceUrl *model.ExternalServiceUrl) error {
	tx, ok := ctx.Value(transactionContextKey).(*gorm.DB)

	gormExternalServiceUrl := r.convertToGormModel(externalServiceUrl, user.ID())

	if !ok {
		return r.db.Execute(ctx).Create(&gormExternalServiceUrl).Error
	}

	return tx.Create(&gormExternalServiceUrl).Error
}

func (r GormExternalServiceRepository) Delete(ctx context.Context, user *model.User, externalServiceUrl *model.ExternalServiceUrl) error {
	tx, ok := ctx.Value(transactionContextKey).(*gorm.DB)

	gormExternalServiceUrl := r.convertToGormModel(externalServiceUrl, user.ID())

	if !ok {
		return r.db.Execute(ctx).Where("user_id = ?", user.ID()).Where("id = ?", externalServiceUrl.ID()).Delete(&gormExternalServiceUrl).Error
	}

	return tx.Where("user_id = ?", user.ID()).Where("id = ?", externalServiceUrl.ID()).Delete(&gormExternalServiceUrl).Error
}

func (r GormExternalServiceRepository) FindByUserId(ctx context.Context, user *model.User) ([]*model.ExternalServiceUrl, error) {
	gormExternalServiceUrls, err := r.db.Execute(ctx).Where("user_id = ?", user.ID()).Find(&gorm_model.ExternalServiceUrl{}).Rows()
	if err != nil {
		return nil, err
	}

	externalServiceUrls := make([]*model.ExternalServiceUrl, 0)

	for gormExternalServiceUrls.Next() {
		var gormExternalServiceUrl gorm_model.ExternalServiceUrl
		err = r.db.ScanRows(ctx, gormExternalServiceUrls, &gormExternalServiceUrl)
		if err != nil {
			return nil, err
		}

		externalServiceUrl, err := model.NewExternalServiceUrl(gormExternalServiceUrl.Url, gormExternalServiceUrl.ID)
		if err != nil {
			return nil, err
		}

		externalServiceUrls = append(externalServiceUrls, externalServiceUrl)
	}

	return externalServiceUrls, nil
}

func (r GormExternalServiceRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	tx, err := r.db.BeginTx(ctx)

	if err != nil {
		return err
	}

	transactionCtx := context.WithValue(ctx, transactionContextKey, tx.Tx())

	if err := fn(transactionCtx); err != nil {
		tx.Tx().Rollback()
		return err
	}

	return tx.Tx().Commit().Error
}

func (r GormExternalServiceRepository) convertToGormModel(externalServiceUrl *model.ExternalServiceUrl, userId string) gorm_model.ExternalServiceUrl {
	return gorm_model.ExternalServiceUrl{
		UserId: userId,
		Url:    externalServiceUrl.Url(),
	}
}
