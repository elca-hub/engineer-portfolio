package repository

import (
	"context"
	"database/sql"
	"devport/domain/model"
	"devport/infra/database/sqlboiler/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type transactionKey struct{}

var transactionContextKey = transactionKey{}

type SqlBoilerUserRepository struct {
	db *sql.DB
}

func NewSqlBoilerUserRepository(db *sql.DB) *SqlBoilerUserRepository {
	return &SqlBoilerUserRepository{
		db: db,
	}
}

func (r *SqlBoilerUserRepository) Create(ctx context.Context, user *model.User) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	model := r.convertToSqlBoilerModel(user)

	if !ok {
		return model.Insert(ctx, r.db, boil.Infer())
	}

	return model.Insert(ctx, tx, boil.Infer())
}

func (r *SqlBoilerUserRepository) Exists(ctx context.Context, email *model.Email) (bool, error) {
	exists, err := models.Users(models.UserWhere.Email.EQ(email.Email())).Exists(ctx, r.db)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *SqlBoilerUserRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	exists, err := models.Users(models.UserWhere.Name.EQ(name)).Exists(ctx, r.db)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *SqlBoilerUserRepository) ExistsById(ctx context.Context, id string) (bool, error) {
	exists, err := models.Users(models.UserWhere.ID.EQ(id)).Exists(ctx, r.db)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *SqlBoilerUserRepository) Update(ctx context.Context, user *model.User) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	model := r.convertToSqlBoilerModel(user)

	if !ok {
		_, err := model.Update(ctx, r.db, boil.Infer())
		return err
	}

	_, err := model.Update(ctx, tx, boil.Infer())
	return err
}

func (r *SqlBoilerUserRepository) FindByEmail(ctx context.Context, email *model.Email) (*model.User, error) {
	user, err := models.Users(models.UserWhere.Email.EQ(email.Email()), qm.Load("Skills"), qm.Load("ExternalServiceUrls")).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return r.convertToDomainModel(user)
}

func (r *SqlBoilerUserRepository) FindById(ctx context.Context, id string) (*model.User, error) {
	user, err := models.Users(models.UserWhere.ID.EQ(id), qm.Load("Skills"), qm.Load("ExternalServiceUrls")).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return r.convertToDomainModel(user)
}

func (r *SqlBoilerUserRepository) FetchIconNamesAll(ctx context.Context) ([]*model.FileIconName, error) {
	users, err := models.Users(models.UserWhere.IconPath.IsNotNull()).All(ctx, r.db)

	if err != nil {
		return nil, err
	}

	iconNames := make([]*model.FileIconName, len(users))

	for i, user := range users {
		iconNames[i], err = model.NewFileName(user.IconPath.String, model.ICON_PATH)
		if err != nil {
			return nil, err
		}
	}

	return iconNames, nil
}

func (r *SqlBoilerUserRepository) FetchHeaderNamesAll(ctx context.Context) ([]*model.FileIconName, error) {
	users, err := models.Users(models.UserWhere.HeaderPath.IsNotNull()).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	headerNames := make([]*model.FileIconName, len(users))

	for i, user := range users {
		headerNames[i], err = model.NewFileName(user.HeaderPath.String, model.HEADER_PATH)
		if err != nil {
			return nil, err
		}
	}

	return headerNames, nil
}

func (r *SqlBoilerUserRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)
	if !ok {
		return fn(ctx)
	}
	return fn(context.WithValue(ctx, transactionContextKey, tx))
}

func (r *SqlBoilerUserRepository) convertToSqlBoilerModel(user *model.User) *models.User {
	return &models.User{
		ID:               user.ID(),
		Name:             user.Name(),
		Email:            user.Email().Email(),
		IconPath:         null.StringFrom(user.IconName()),
		HeaderPath:       null.StringFrom(user.HeaderIconName()),
		OrganizationName: null.StringFrom(user.OrganizationName()),
		OccupationName:   null.StringFrom(user.OccupationName()),
		Place:            null.StringFrom(user.Place()),
		BioPath:          null.StringFrom(user.BioPath()),
		Birthday:         user.Birthday(),
		CreatedAt:        user.CreatedAt(),
		UpdatedAt:        user.UpdatedAt(),
	}
}

func (r *SqlBoilerUserRepository) convertToDomainModel(sqlboilerUser *models.User) (*model.User, error) {
	userEmail, err := model.NewEmail(sqlboilerUser.Email)

	if err != nil {
		return nil, err
	}

	sqlboilerSkills := sqlboilerUser.R.Skills

	skills := make([]uint, len(sqlboilerSkills))

	// idを取得
	for i, sqlboilerSkill := range sqlboilerSkills {
		skills[i] = uint(sqlboilerSkill.ID)
	}

	sqlboilerExternalServiceUrl := sqlboilerUser.R.ExternalServiceUrls

	externalServiceUrls := make([]string, len(sqlboilerExternalServiceUrl))

	for i, sqlboilerExternalServiceUrl := range sqlboilerExternalServiceUrl {
		externalServiceUrls[i] = sqlboilerExternalServiceUrl.ID
	}

	user, err := model.NewUser(
		sqlboilerUser.ID,
		sqlboilerUser.Name,
		sqlboilerUser.Birthday,
		userEmail,
		sqlboilerUser.IconPath.String,
		sqlboilerUser.HeaderPath.String,
		sqlboilerUser.BioPath.String,
		sqlboilerUser.OrganizationName.String,
		sqlboilerUser.OccupationName.String,
		sqlboilerUser.Place.String,
		sqlboilerUser.CreatedAt,
		sqlboilerUser.UpdatedAt,
		skills,
		externalServiceUrls,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
