package repo

import (
	"context"
	"devport/domain/model"
	"devport/infra/database/gorm/gorm_model"
	"errors"

	"gorm.io/gorm"
)

type GormExternalServiceRepository struct {
	db *gorm.DB
}

func NewGormExternalServiceRepository(db *gorm.DB) *GormExternalServiceRepository {
	return &GormExternalServiceRepository{
		db: db,
	}
}

func (r GormExternalServiceRepository) Create(ctx context.Context, user *model.User, externalServiceUrl *model.ExternalServiceUrl) error {
	tx, ok := ctx.Value(transactionContextKey).(*gorm.DB)

	gormExternalServiceUrl := r.convertToGormModel(externalServiceUrl, user.ID())

	if !ok {
		return r.db.Create(&gormExternalServiceUrl).Error
	}

	return tx.Create(&gormExternalServiceUrl).Error
}

func (r GormExternalServiceRepository) Delete(ctx context.Context, user *model.User, externalServiceUrl *model.ExternalServiceUrl) error {
	tx, ok := ctx.Value(transactionContextKey).(*gorm.DB)

	gormExternalServiceUrl := r.convertToGormModel(externalServiceUrl, user.ID())

	if !ok {
		return r.db.Where("user_id = ?", user.ID()).Where("id = ?", externalServiceUrl.ID()).Delete(&gormExternalServiceUrl).Error
	}

	return tx.Where("user_id = ?", user.ID()).Where("id = ?", externalServiceUrl.ID()).Delete(&gormExternalServiceUrl).Error
}

func (r GormExternalServiceRepository) FindByUserId(ctx context.Context, user *model.User) ([]*model.ExternalServiceUrl, error) {
	var gormExternalServiceUrls []gorm_model.ExternalServiceUrl
	// urlが存在するかどうか
	err := r.db.Where("user_id = ?", user.ID()).Where("url <> ''").Find(&gormExternalServiceUrls).Error

	if err != nil {
		return nil, err
	}

	externalServiceUrls := make([]*model.ExternalServiceUrl, len(gormExternalServiceUrls))
	for i, gormExternalServiceUrl := range gormExternalServiceUrls {
		externalServiceUrl, err := r.convertToModel(&gormExternalServiceUrl)
		if err != nil {
			return nil, err
		}
		externalServiceUrls[i] = externalServiceUrl
	}

	return externalServiceUrls, nil
}

func (r GormExternalServiceRepository) FindByServiceType(ctx context.Context, user *model.User, serviceType int) (*model.ExternalServiceUrl, error) {
	var gormExternalServiceUrl gorm_model.ExternalServiceUrl
	err := r.db.Where("user_id = ?", user.ID()).Where("service_type = ?", serviceType).First(&gormExternalServiceUrl).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // ここを変える！！！
		}
		return nil, err
	}

	return r.convertToModel(&gormExternalServiceUrl)
}

func (r GormExternalServiceRepository) FindById(ctx context.Context, id string) (*model.ExternalServiceUrl, error) {
	var gormExternalServiceUrl gorm_model.ExternalServiceUrl
	err := r.db.Where("id = ?", id).First(&gormExternalServiceUrl).Error
	if err != nil {
		return nil, err
	}

	return r.convertToModel(&gormExternalServiceUrl)
}

func (r GormExternalServiceRepository) Update(ctx context.Context, user *model.User, externalServiceUrl *model.ExternalServiceUrl) error {
	tx, ok := ctx.Value(transactionContextKey).(*gorm.DB)

	gormExternalServiceUrl := r.convertToGormModel(externalServiceUrl, user.ID())

	if !ok {
		return r.db.Model(&gorm_model.ExternalServiceUrl{}).Where("id = ?", externalServiceUrl.ID()).Where("user_id = ?", user.ID()).Updates(gormExternalServiceUrl).Error
	}

	return tx.Model(&gorm_model.ExternalServiceUrl{}).Where("id = ?", externalServiceUrl.ID()).Where("user_id = ?", user.ID()).Updates(gormExternalServiceUrl).Error
}

func (r GormExternalServiceRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	tx := r.db.Begin()

	transactionCtx := context.WithValue(ctx, transactionContextKey, tx)

	if err := fn(transactionCtx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r GormExternalServiceRepository) convertToModel(gormExternalServiceUrl *gorm_model.ExternalServiceUrl) (*model.ExternalServiceUrl, error) {
	return model.NewExternalServiceUrl(gormExternalServiceUrl.ID, gormExternalServiceUrl.ServiceType, gormExternalServiceUrl.Url)
}

func (r GormExternalServiceRepository) convertToGormModel(externalServiceUrl *model.ExternalServiceUrl, userId string) gorm_model.ExternalServiceUrl {
	return gorm_model.ExternalServiceUrl{
		ID:          externalServiceUrl.ID(),
		UserId:      userId,
		Url:         externalServiceUrl.Url(),
		ServiceType: externalServiceUrl.ServiceType(),
	}
}
