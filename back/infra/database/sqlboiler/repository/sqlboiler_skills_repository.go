package repository

import (
	"context"
	"database/sql"
	"devport/domain/model"
	"devport/infra/database/sqlboiler/models"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SqlBoilerSkillsRepository struct {
	db *sql.DB
}

func NewSqlBoilerSkillsRepository(db *sql.DB) *SqlBoilerSkillsRepository {
	return &SqlBoilerSkillsRepository{
		db: db,
	}
}

func (r *SqlBoilerSkillsRepository) Create(ctx context.Context, skill *model.Skill, userID string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	sqlboilerModel := r.convertToSqlBoilerModel(skill, userID)

	if !ok {
		return sqlboilerModel.Insert(ctx, r.db, boil.Infer())
	}

	return sqlboilerModel.Insert(ctx, tx, boil.Infer())
}

func (r *SqlBoilerSkillsRepository) FindByUserID(ctx context.Context, userID string) ([]*model.Skill, error) {
	skills, err := models.Skills(
		models.SkillWhere.UserID.EQ(userID),
		qm.OrderBy("sort_index ASC"),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	result := make([]*model.Skill, len(skills))
	for i, skill := range skills {
		domainSkill, err := r.convertToDomainModel(skill)
		if err != nil {
			return nil, err
		}
		result[i] = domainSkill
	}

	return result, nil
}

func (r *SqlBoilerSkillsRepository) Update(ctx context.Context, skill *model.Skill, userID string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	sqlboilerModel := r.convertToSqlBoilerModel(skill, userID)

	if !ok {
		_, err := sqlboilerModel.Update(ctx, r.db, boil.Infer())
		return err
	}

	_, err := sqlboilerModel.Update(ctx, tx, boil.Infer())
	return err
}

func (r *SqlBoilerSkillsRepository) Delete(ctx context.Context, userID string, id string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	if !ok {
		_, err := models.Skills(models.SkillWhere.ID.EQ(id), models.SkillWhere.UserID.EQ(userID)).DeleteAll(ctx, r.db)
		return err
	}

	_, err := models.Skills(models.SkillWhere.ID.EQ(id)).DeleteAll(ctx, tx)
	return err
}

func (r *SqlBoilerSkillsRepository) convertToSqlBoilerModel(skill *model.Skill, userID string) *models.Skill {
	return &models.Skill{
		Name:      skill.Name(),
		WhenDate:  skill.Year(),
		Comment:   null.StringFrom(skill.Comment()),
		SortIndex: skill.SortIndex(),
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ID:        skill.ID(),
	}
}

func (r *SqlBoilerSkillsRepository) GetMaxSortIndex(ctx context.Context, userID string) (int, error) {
	skills, err := models.Skills(
		models.SkillWhere.UserID.EQ(userID),
		qm.OrderBy("sort_index DESC"),
		qm.Limit(1),
	).All(ctx, r.db)
	if err != nil {
		return 0, err
	}

	if len(skills) == 0 {
		return 0, nil
	}

	return skills[0].SortIndex, nil
}

func (r *SqlBoilerSkillsRepository) convertToDomainModel(sqlboilerSkill *models.Skill) (*model.Skill, error) {
	comment := ""
	if sqlboilerSkill.Comment.Valid {
		comment = sqlboilerSkill.Comment.String
	}

	return model.NewSkillWithID(
		sqlboilerSkill.ID,
		sqlboilerSkill.Name,
		sqlboilerSkill.WhenDate,
		comment,
		sqlboilerSkill.SortIndex,
	)
}

func (r *SqlBoilerSkillsRepository) FindByID(ctx context.Context, userId string, skillId string) (*model.Skill, error) {
	skill, err := models.Skills(
		models.SkillWhere.UserID.EQ(userId),
		models.SkillWhere.ID.EQ(skillId),
	).One(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return r.convertToDomainModel(skill)
}

func (r *SqlBoilerSkillsRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
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
