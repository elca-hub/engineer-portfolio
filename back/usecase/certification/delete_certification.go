package certification

import (
	"context"
	"devport/domain/repo/db"
	"time"
)

type (
	DeleteCertificationUseCase interface {
		Execute(context.Context, DeleteCertificationInput) error
	}

	DeleteCertificationInput struct {
		ID     string `validate:"required"`
		UserId string `validate:"required"`
	}

	deleteCertificationInteractor struct {
		certificationsRepository db.CertificationsRepository
		userRepository           db.UserRepository
		ctxTimeout               time.Duration
	}
)

func NewDeleteCertificationInteractor(
	certificationsRepository db.CertificationsRepository,
	userRepository db.UserRepository,
	t time.Duration,
) DeleteCertificationUseCase {
	return deleteCertificationInteractor{
		certificationsRepository: certificationsRepository,
		userRepository:           userRepository,
		ctxTimeout:               t,
	}
}

func (i deleteCertificationInteractor) Execute(ctx context.Context, input DeleteCertificationInput) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	return i.certificationsRepository.Delete(ctx, input.ID)
}
