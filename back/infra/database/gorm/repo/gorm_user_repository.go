package repo

import (
	"context"
	"devport/domain/model"
	"devport/infra/database/gorm/gorm_model"
	"errors"

	"gorm.io/gorm"
)

type transactionKey struct{}

var transactionContextKey = transactionKey{}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (r GormUserRepository) Create(ctx context.Context, user *model.User) error {
	tx, ok := ctx.Value(transactionContextKey).(*gorm.DB)
	gormUser := r.convertToGormModel(*user)

	if !ok {
		return r.db.Create(&gormUser).Error
	}

	return tx.Create(&gormUser).Error
}

func (r GormUserRepository) Exists(ctx context.Context, email *model.Email) (bool, error) {
	first := r.db.Where("email = ?", email.Email()).First(&gorm_model.User{})

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

	r.db.Model(&gorm_model.User{}).Where("name = ?", name).Count(&counter)

	return counter > 0, nil
}

func (r GormUserRepository) ExistsById(ctx context.Context, id string) (bool, error) {
	var counter int64

	r.db.Model(&gorm_model.User{}).Where("id = ?", id).Count(&counter)

	return counter > 0, nil
}

func (r GormUserRepository) Update(ctx context.Context, user *model.User) error {
	tx, ok := ctx.Value(transactionContextKey).(*gorm.DB)
	gormUser := r.convertToGormModel(*user)

	if !ok {
		return r.db.Save(&gormUser).Error
	}

	return tx.Save(&gormUser).Error
}

func (r GormUserRepository) FindByEmail(ctx context.Context, email *model.Email) (*model.User, error) {
	var gormUser gorm_model.User

	if err := r.db.Where("email = ?", email.Email()).First(&gormUser).Error; err != nil {
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

	if err := r.db.Where("id = ?", id).First(&gormUser).Error; err != nil {
		return nil, err
	}

	user, err := r.convertToDomainModel(gormUser)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r GormUserRepository) FetchIconNamesAll(ctx context.Context) ([]*model.FileIconName, error) {
	var gormUsers []gorm_model.User

	if err := r.db.Where("icon_path <> ''").Find(&gormUsers).Error; err != nil {
		return nil, err
	}

	res := make([]*model.FileIconName, len(gormUsers))

	for i, user := range gormUsers {
		var err error
		res[i], err = model.NewFileIconName(user.IconPath, model.ICON_PATH)

		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (r GormUserRepository) FetchHeaderNamesAll(ctx context.Context) ([]*model.FileIconName, error) {
	var gormUsers []gorm_model.User

	if err := r.db.Where("header_path <> ''").Find(&gormUsers).Error; err != nil {
		return nil, err
	}

	res := make([]*model.FileIconName, len(gormUsers))

	for i, user := range gormUsers {
		var err error
		res[i], err = model.NewFileIconName(user.HeaderPath, model.HEADER_PATH)

		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (r GormUserRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	tx := r.db.Begin()

	transactionCtx := context.WithValue(ctx, transactionContextKey, tx)

	if err := fn(transactionCtx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r GormUserRepository) convertToGormModel(user model.User) gorm_model.User {
	email := user.Email()

	return gorm_model.User{
		ID:               user.ID(),
		Name:             user.Name(),
		Birthday:         user.Birthday(),
		Email:            email.Email(),
		IconPath:         user.IconName(),
		HeaderPath:       user.HeaderIconName(),
		BioPath:          user.BioPath(),
		OrganizationName: user.OrganizationName(),
		OccupationName:   user.OccupationName(),
		Place:            user.Place(),
		CreatedAt:        user.CreatedAt(),
		UpdatedAt:        user.UpdatedAt(),
	}
}

func (r GormUserRepository) convertToDomainModel(gormUser gorm_model.User) (*model.User, error) {
	userEmail, err := model.NewEmail(gormUser.Email)

	if err != nil {
		return nil, err
	}

	gormSkills := gormUser.Skills

	skills := make([]uint, len(gormSkills))

	for i, gormSkill := range gormSkills {
		skills[i] = gormSkill.ID
	}

	gormExternalServiceUrl := gormUser.ExternalServiceUrls

	externalServiceUrls := make([]string, len(gormExternalServiceUrl))

	for i, gormExternalServiceUrl := range gormExternalServiceUrl {
		externalServiceUrls[i] = gormExternalServiceUrl.ID
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
