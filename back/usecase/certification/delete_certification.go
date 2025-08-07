package certification

import (
	"context"
	"devport/domain/repo/db"
	"errors"
	"time"
)

type (
	DeleteCertificationUseCase interface {
		Execute(context.Context, DeleteCertificationInput) (DeleteCertificationOutput, error)
	}

	DeleteCertificationInput struct {
		ID     string `validate:"required"`
		UserId string `validate:"required"`
	}

	DeleteCertificationOutput struct{}

	DeleteCertificationPresenter interface {
		Output() DeleteCertificationOutput
	}

	deleteCertificationInteractor struct {
		certificationsRepository db.CertificationsRepository
		userRepository           db.UserRepository
		presenter                DeleteCertificationPresenter
		ctxTimeout               time.Duration
	}
)

func NewDeleteCertificationInteractor(
	certificationsRepository db.CertificationsRepository,
	userRepository db.UserRepository,
	presenter DeleteCertificationPresenter,
	t time.Duration,
) DeleteCertificationUseCase {
	return deleteCertificationInteractor{
		certificationsRepository: certificationsRepository,
		userRepository:           userRepository,
		presenter:                presenter,
		ctxTimeout:               t,
	}
}

func (i deleteCertificationInteractor) Execute(ctx context.Context, input DeleteCertificationInput) (DeleteCertificationOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	if isExists, err := i.userRepository.ExistsById(ctx, input.UserId); err != nil || !isExists {
		if err != nil {
			return DeleteCertificationOutput{}, err
		}
		if !isExists {
			return DeleteCertificationOutput{}, errors.New("ユーザーが存在しません")
		}
	}

	err := i.certificationsRepository.Delete(ctx, input.UserId, input.ID)

	if err != nil {
		return DeleteCertificationOutput{}, err
	}

	return i.presenter.Output(), nil
}
