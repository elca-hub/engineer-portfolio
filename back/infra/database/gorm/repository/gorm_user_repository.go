package repository

import (
	"context"
	"devport/adapter/repository"
	"devport/domain/model"
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

func (r GormUserRepository) Create(ctx context.Context, user *model.User) error {
	tx, ok := ctx.Value("TransactionContextKey").(*gorm.DB)
	gormUser := convertToGormModel(*user)

	if !ok {
		return r.db.Execute(ctx).Create(&gormUser).Error
	}

	return tx.Create(&gormUser).Error
}

func (r GormUserRepository) Exists(ctx context.Context, email *model.Email) (bool, error) {
	var counter int64

	r.db.Execute(ctx).Model(&gorm_model.User{}).Where("email = ?", email.Email()).Count(&counter)

	return counter > 0, nil
}

func (r GormUserRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var counter int64

	r.db.Execute(ctx).Model(&gorm_model.User{}).Where("name = ?", name).Count(&counter)

	return counter > 0, nil
}

func (r GormUserRepository) Update(ctx context.Context, user *model.User) error {
	tx, ok := ctx.Value("TransactionContextKey").(*gorm.DB)
	gormUser := convertToGormModel(*user)

	if !ok {
		return r.db.Execute(ctx).Save(&gormUser).Error
	}

	return tx.Save(&gormUser).Error
}

func (r GormUserRepository) FindByEmail(ctx context.Context, email *model.Email) (*model.User, error) {
	var gormUser gorm_model.User

	if err := r.db.Execute(ctx).Where("email = ?", email.Email()).First(&gormUser).Error; err != nil {
		return nil, err
	}

	user, err := convertToDomainModel(gormUser)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r GormUserRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	tx, err := r.db.BeginTx(ctx)

	if err != nil {
		return err
	}

	transactionCtx := context.WithValue(ctx, "TransactionContextKey", tx.Tx())

	if err := fn(transactionCtx); err != nil {
		tx.Tx().Rollback()

		return err
	}

	return tx.Tx().Commit().Error
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
