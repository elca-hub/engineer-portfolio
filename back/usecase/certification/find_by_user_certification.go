package certification

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/db"
	"time"
)

type (
	FindByUserCertificationUseCase interface {
		Execute(context.Context, FindByUserCertificationInput) (FindByUserCertificationOutput, error)
	}

	FindByUserCertificationInput struct {
		UserId string `validate:"required"`
	}

	FindByUserCertificationOutput struct {
		Certifications []*dto.CertificationDTO `json:"certifications"`
	}

	FindByUserCertificationPresenter interface {
		Output(certifications []*model.Certification) FindByUserCertificationOutput
	}

	findByUserCertificationInteractor struct {
		certificationsRepository db.CertificationsRepository
		userRepository           db.UserRepository
		presenter                FindByUserCertificationPresenter
		ctxTimeout               time.Duration
	}
)

func NewFindByUserCertificationInteractor(
	certificationsRepository db.CertificationsRepository,
	userRepository db.UserRepository,
	presenter FindByUserCertificationPresenter,
	t time.Duration,
) FindByUserCertificationUseCase {
	return findByUserCertificationInteractor{
		certificationsRepository: certificationsRepository,
		userRepository:           userRepository,
		presenter:                presenter,
		ctxTimeout:               t,
	}
}

func (i findByUserCertificationInteractor) Execute(ctx context.Context, input FindByUserCertificationInput) (FindByUserCertificationOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	_, err := i.userRepository.FindById(ctx, input.UserId, nil)
	if err != nil {
		return FindByUserCertificationOutput{}, err
	}

	certifications, err := i.certificationsRepository.FindByUserID(ctx, input.UserId)
	if err != nil {
		return FindByUserCertificationOutput{}, err
	}

	return i.presenter.Output(certifications), nil
}