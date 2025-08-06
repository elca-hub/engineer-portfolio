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

type SqlBoilerCertificationRepository struct {
	db *sql.DB
}

func NewSqlBoilerCertificationRepository(db *sql.DB) *SqlBoilerCertificationRepository {
	return &SqlBoilerCertificationRepository{
		db: db,
	}
}

func (r *SqlBoilerCertificationRepository) Create(ctx context.Context, certification *model.Certification, userID string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	sqlboilerModel := r.convertToSqlBoilerModel(certification, userID)

	if !ok {
		return sqlboilerModel.Insert(ctx, r.db, boil.Infer())
	}

	return sqlboilerModel.Insert(ctx, tx, boil.Infer())
}

func (r *SqlBoilerCertificationRepository) FindByUserID(ctx context.Context, userID string) ([]*model.Certification, error) {
	certifications, err := models.Certifications(
		models.CertificationWhere.UserID.EQ(userID),
		qm.OrderBy("sort_index ASC"),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	result := make([]*model.Certification, len(certifications))
	for i, cert := range certifications {
		domainCert, err := r.convertToDomainModel(cert)
		if err != nil {
			return nil, err
		}
		result[i] = domainCert
	}

	return result, nil
}

func (r *SqlBoilerCertificationRepository) Update(ctx context.Context, certification *model.Certification, userID string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	sqlboilerModel := r.convertToSqlBoilerModel(certification, userID)

	if !ok {
		_, err := sqlboilerModel.Update(ctx, r.db, boil.Infer())
		return err
	}

	_, err := sqlboilerModel.Update(ctx, tx, boil.Infer())
	return err
}

func (r *SqlBoilerCertificationRepository) Delete(ctx context.Context, id string) error {
	tx, ok := ctx.Value(transactionContextKey).(*sql.Tx)

	if !ok {
		_, err := models.Certifications(models.CertificationWhere.ID.EQ(id)).DeleteAll(ctx, r.db)
		return err
	}

	_, err := models.Certifications(models.CertificationWhere.ID.EQ(id)).DeleteAll(ctx, tx)
	return err
}

func (r *SqlBoilerCertificationRepository) convertToSqlBoilerModel(certification *model.Certification, userID string) *models.Certification {
	return &models.Certification{
		Name:      certification.Name(),
		WhenDate:  certification.Year(),
		Comment:   null.StringFrom(certification.Comment()),
		SortIndex: certification.SortIndex(),
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (r *SqlBoilerCertificationRepository) GetMaxSortIndex(ctx context.Context, userID string) (int, error) {
	certifications, err := models.Certifications(
		models.CertificationWhere.UserID.EQ(userID),
		qm.OrderBy("sort_index DESC"),
		qm.Limit(1),
	).All(ctx, r.db)
	if err != nil {
		return 0, err
	}

	if len(certifications) == 0 {
		return 0, nil
	}

	return certifications[0].SortIndex, nil
}

func (r *SqlBoilerCertificationRepository) convertToDomainModel(sqlboilerCertification *models.Certification) (*model.Certification, error) {
	comment := ""
	if sqlboilerCertification.Comment.Valid {
		comment = sqlboilerCertification.Comment.String
	}

	return model.NewCertificationWithID(
		sqlboilerCertification.ID,
		sqlboilerCertification.Name,
		sqlboilerCertification.WhenDate,
		comment,
		sqlboilerCertification.SortIndex,
	)
}