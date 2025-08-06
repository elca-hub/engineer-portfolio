package certification

import (
	"context"
	"devport/domain/repo/db"
	"time"
)

type (
	UpdateCertificationsSortUseCase interface {
		Execute(context.Context, UpdateCertificationsSortInput) (UpdateCertificationSortOutput, error)
	}

	UpdateCertificationsSortInput struct {
		UserId string `validate:"required"`

		SortList []struct {
			ID        string `validate:"required" json:"id"`
			SortIndex int    `json:"sort_index"`
		} `json:"sort_list" validate:"required,dive"`
	}

	UpdateCertificationSortOutput struct {
	}

	UpdateCertificationSortPresenter interface {
		Output() UpdateCertificationSortOutput
	}

	updateCertificationSortInteractor struct {
		certificationsRepository db.CertificationsRepository
		userRepository           db.UserRepository
		presenter                UpdateCertificationSortPresenter
		ctxTimeout               time.Duration
	}
)

func NewUpdateCertificationSortInteractor(
	certificationsRepository db.CertificationsRepository,
	userRepository db.UserRepository,
	presenter UpdateCertificationSortPresenter,
	t time.Duration,
) UpdateCertificationsSortUseCase {
	return updateCertificationSortInteractor{
		certificationsRepository: certificationsRepository,
		userRepository:           userRepository,
		presenter:                presenter,
		ctxTimeout:               t,
	}
}

func (i updateCertificationSortInteractor) Execute(ctx context.Context, input UpdateCertificationsSortInput) (UpdateCertificationSortOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	err := i.certificationsRepository.WithTransaction(ctx, func(ctx context.Context) error {
		for _, certificationSort := range input.SortList {
			certification, err := i.certificationsRepository.FindByID(ctx, input.UserId, certificationSort.ID)

			if err != nil {
				return err
			}

			if err := certification.UpdateSortIndex(certificationSort.SortIndex); err != nil {
				return err
			}

			if err := i.certificationsRepository.Update(ctx, certification, input.UserId); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return UpdateCertificationSortOutput{}, err
	}

	return i.presenter.Output(), nil
}
