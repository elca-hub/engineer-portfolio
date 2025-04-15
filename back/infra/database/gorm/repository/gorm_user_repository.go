package repository

import (
	"context"
	"devport/adapter/repository"
	"devport/domain/model"
	"devport/infra/database/gorm/gorm_model"
	"errors"

	"gorm.io/gorm"
)

type transactionKey struct{}

var transactionContextKey = transactionKey{}

type GormUserRepository struct {
	db repository.SQL
}

func NewGormUserRepository(db repository.SQL) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (r GormUserRepository) Create(ctx context.Context, user *model.User) error {
	tx, ok := ctx.Value(transactionContextKey).(*gorm.DB)
	gormUser := r.convertToGormModel(*user)

	if !ok {
		return r.db.Execute(ctx).Create(&gormUser).Error
	}

	return tx.Create(&gormUser).Error
}

func (r GormUserRepository) Exists(ctx context.Context, email *model.Email) (bool, error) {
	first := r.db.Execute(ctx).Where("email = ?", email.Email()).First(&gorm_model.User{})

	if errors.Is(first.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if first.Error != nil {
		return false, first.Error
	}

	return true, nil
}

func (r GormUserRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var counter int64

	r.db.Execute(ctx).Model(&gorm_model.User{}).Where("name = ?", name).Count(&counter)

	return counter > 0, nil
}

func (r GormUserRepository) ExistsById(ctx context.Context, id string) (bool, error) {
	var counter int64

	r.db.Execute(ctx).Model(&gorm_model.User{}).Where("id = ?", id).Count(&counter)

	return counter > 0, nil
}

func (r GormUserRepository) Update(ctx context.Context, user *model.User) error {
	tx, ok := ctx.Value(transactionContextKey).(*gorm.DB)
	gormUser := r.convertToGormModel(*user)

	if !ok {
		return r.db.Execute(ctx).Save(&gormUser).Error
	}

	return tx.Save(&gormUser).Error
}

func (r GormUserRepository) FindByEmail(ctx context.Context, email *model.Email) (*model.User, error) {
	var gormUser gorm_model.User

	if err := r.db.Execute(ctx).Where("email = ?", email.Email()).First(&gormUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	user, err := r.convertToDomainModel(gormUser)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r GormUserRepository) FindById(ctx context.Context, id string) (*model.User, error) {
	var gormUser gorm_model.User

	if err := r.db.Execute(ctx).Where("id = ?", id).First(&gormUser).Error; err != nil {
		return nil, err
	}

	user, err := r.convertToDomainModel(gormUser)

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

	transactionCtx := context.WithValue(ctx, transactionContextKey, tx.Tx())

	if err := fn(transactionCtx); err != nil {
		tx.Tx().Rollback()
		return err
	}

	return tx.Tx().Commit().Error
}

func (r GormUserRepository) convertToGormModel(user model.User) gorm_model.User {
	email := user.Email()

	modelSkills := user.Skills()

	skills := make([]gorm_model.Skill, len(modelSkills))

	for i, modelSkill := range modelSkills {
		skills[i] = gorm_model.Skill{
			Name:      modelSkill.Name(),
			Status:    modelSkill.Status(),
			When:      modelSkill.When(),
			SortIndex: modelSkill.SortIndex(),
		}
	}

	modelExternalServiceUrls := user.ExternalServiceURLs()

	externalServiceUrls := make([]gorm_model.ExternalServiceUrl, len(modelExternalServiceUrls))

	for i, modelExternalServiceUrl := range modelExternalServiceUrls {
		externalServiceUrls[i] = gorm_model.ExternalServiceUrl{
			UserId: user.ID(),
			Url:    modelExternalServiceUrl.Url(),
		}
	}

	return gorm_model.User{
		ID:                  user.ID(),
		Name:                user.Name(),
		Birthday:            user.Birthday(),
		Email:               email.Email(),
		IconPath:            user.IconName(),
		HeaderPath:          user.HeaderIconName(),
		BioPath:             user.BioPath(),
		OrganizationName:    user.OrganizationName(),
		OccupationName:      user.OccupationName(),
		Place:               user.Place(),
		CreatedAt:           user.CreatedAt(),
		UpdatedAt:           user.UpdatedAt(),
		Skills:              skills,
		ExternalServiceUrls: externalServiceUrls,
	}
}

func (r GormUserRepository) convertToDomainModel(gormUser gorm_model.User) (*model.User, error) {
	userEmail, err := model.NewEmail(gormUser.Email)

	if err != nil {
		return nil, err
	}

	gormSkills := gormUser.Skills

	skills := make([]*model.Skill, len(gormSkills))

	for i, gormSkill := range gormSkills {
		skills[i], err = model.NewSkill(gormSkill.Name, gormSkill.Status, gormSkill.When, gormSkill.SortIndex)

		if err != nil {
			return nil, err
		}
	}

	gormExternalServiceUrl := gormUser.ExternalServiceUrls

	externalServiceUrls := make([]*model.ExternalServiceUrl, len(gormExternalServiceUrl))

	for i, gormExternalServiceUrl := range gormExternalServiceUrl {
		externalServiceUrls[i], err = model.NewExternalServiceUrl(gormExternalServiceUrl.ServiceType, gormExternalServiceUrl.Url)

		if err != nil {
			return nil, err
		}
	}

	user, err := model.NewUser(
		gormUser.ID,
		gormUser.Name,
		gormUser.Birthday,
		userEmail,
		gormUser.IconPath,
		gormUser.HeaderPath,
		gormUser.BioPath,
		gormUser.OrganizationName,
		gormUser.OccupationName,
		gormUser.Place,
		gormUser.CreatedAt,
		gormUser.UpdatedAt,
		skills,
		externalServiceUrls,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
