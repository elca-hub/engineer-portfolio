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

func (r *SqlBoilerUserRepository) FindByEmail(ctx context.Context, email *model.Email, bio *model.Bio) (*model.User, error) {
	user, err := models.Users(models.UserWhere.Email.EQ(email.Email()), qm.Load("Skills"), qm.Load("Certifications"), qm.Load("ExternalServiceUrls")).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return r.convertToDomainModel(user, bio)
}

func (r *SqlBoilerUserRepository) FindById(ctx context.Context, id string, bio *model.Bio) (*model.User, error) {
	user, err := models.Users(models.UserWhere.ID.EQ(id), qm.Load("Skills"), qm.Load("Certifications"), qm.Load("ExternalServiceUrls")).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return r.convertToDomainModel(user, bio)
}

func (r *SqlBoilerUserRepository) FetchIconNamesAll(ctx context.Context) ([]string, error) {
	users, err := models.Users(models.UserWhere.IconPath.IsNotNull()).All(ctx, r.db)

	if err != nil {
		return nil, err
	}

	iconNames := make([]string, len(users))

	// 取得時にis not nullをしているのでそのまま取得してok
	for i, user := range users {
		iconNames[i] = user.IconPath.String
	}

	return iconNames, nil
}

func (r *SqlBoilerUserRepository) FetchHeaderNamesAll(ctx context.Context) ([]string, error) {
	users, err := models.Users(models.UserWhere.HeaderPath.IsNotNull()).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	headerNames := make([]string, len(users))

	for i, user := range users {
		headerNames[i] = user.HeaderPath.String
	}

	return headerNames, nil
}

func (r *SqlBoilerUserRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
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
		Birthday:         user.Birthday(),
		CreatedAt:        user.CreatedAt(),
		UpdatedAt:        user.UpdatedAt(),
	}
}

func (r *SqlBoilerUserRepository) convertToDomainModel(sqlboilerUser *models.User, bio *model.Bio) (*model.User, error) {
	userEmail, err := model.NewEmail(sqlboilerUser.Email)
	if err != nil {
		return nil, err
	}

	// Skills変換
	sqlboilerSkills := sqlboilerUser.R.Skills
	skills := make([]*model.Skill, len(sqlboilerSkills))
	for i, sqlboilerSkill := range sqlboilerSkills {
		comment := ""
		if sqlboilerSkill.Comment.Valid {
			comment = sqlboilerSkill.Comment.String
		}
		skill, err := model.NewSkillWithID(
			sqlboilerSkill.ID,
			sqlboilerSkill.Name,
			sqlboilerSkill.WhenDate,
			comment,
			sqlboilerSkill.SortIndex,
		)
		if err != nil {
			return nil, err
		}
		skills[i] = skill
	}

	// Certifications変換
	sqlboilerCertifications := sqlboilerUser.R.Certifications
	certifications := make([]*model.Certification, len(sqlboilerCertifications))
	for i, sqlboilerCertification := range sqlboilerCertifications {
		comment := ""
		if sqlboilerCertification.Comment.Valid {
			comment = sqlboilerCertification.Comment.String
		}
		certification, err := model.NewCertificationWithID(
			sqlboilerCertification.ID,
			sqlboilerCertification.Name,
			sqlboilerCertification.WhenDate,
			comment,
			sqlboilerCertification.SortIndex,
		)
		if err != nil {
			return nil, err
		}
		certifications[i] = certification
	}

	// ExternalServiceUrls変換
	sqlboilerExternalServiceUrls := sqlboilerUser.R.ExternalServiceUrls
	externalServiceUrls := make([]*model.ExternalServiceUrl, len(sqlboilerExternalServiceUrls))
	for i, sqlboilerExternalServiceUrl := range sqlboilerExternalServiceUrls {
		externalServiceUrl, err := model.NewExternalServiceUrl(
			sqlboilerExternalServiceUrl.ID,
			sqlboilerExternalServiceUrl.ServiceType,
			sqlboilerExternalServiceUrl.URL,
		)
		if err != nil {
			return nil, err
		}
		externalServiceUrls[i] = externalServiceUrl
	}

	user, err := model.NewUser(
		sqlboilerUser.ID,
		sqlboilerUser.Name,
		sqlboilerUser.Birthday,
		userEmail,
		sqlboilerUser.IconPath.String,
		sqlboilerUser.HeaderPath.String,
		bio,
		sqlboilerUser.OrganizationName.String,
		sqlboilerUser.OccupationName.String,
		sqlboilerUser.Place.String,
		sqlboilerUser.CreatedAt,
		sqlboilerUser.UpdatedAt,
		skills,
		certifications,
		externalServiceUrls,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
