package repository

import (
	"context"
	"devport/domain/model"
	"devport/domain/repository"
	"devport/infra/database/gorm/gorm_model"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db repository.SQL
}

func NewGormUserRepository(db repository.SQL) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (r GormUserRepository) Create(tx *gorm.DB, user *model.User) error {
	gormUser := convertToGormModel(*user)

	return tx.Create(&gormUser).Error
}

func (r GormUserRepository) Exists(tx *gorm.DB, email *model.Email) (bool, error) {
	var counter int64

	tx.Model(&gorm_model.User{}).Where("email = ?", email.Email()).Count(&counter)

	return counter > 0, nil
}

func (r GormUserRepository) ExistsByName(tx *gorm.DB, name string) (bool, error) {
	var counter int64

	tx.Model(&gorm_model.User{}).Where("name = ?", name).Count(&counter)

	return counter > 0, nil
}

func (r GormUserRepository) Update(tx *gorm.DB, user *model.User) error {
	gormUser := convertToGormModel(*user)

	return tx.Save(&gormUser).Error
}

func (r GormUserRepository) FindByEmail(tx *gorm.DB, email *model.Email) (*model.User, error) {
	var gormUser gorm_model.User

	if err := tx.Where("email = ?", email.Email()).First(&gormUser).Error; err != nil {
		return nil, err
	}

	user, err := convertToDomainModel(gormUser)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r GormUserRepository) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	tx, err := r.db.BeginTx(ctx)

	if err != nil {
		return err
	}

	if err := fn(tx.Tx()); err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}

func convertToGormModel(user model.User) gorm_model.User {
	email := user.Email()

	return gorm_model.User{
		ID:        user.ID().ID(),
		Name:      user.Name(),
		Birthday:  user.Birthday(),
		Email:     email.Email(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}
}

func convertToDomainModel(gormUser gorm_model.User) (*model.User, error) {
	userEmail, err := model.NewEmail(gormUser.Email)

	if err != nil {
		return nil, err
	}

	user, err := model.NewUser(
		model.NewUUID(gormUser.ID),
		gormUser.Name,
		gormUser.Birthday,
		userEmail,
		gormUser.CreatedAt,
		gormUser.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
