package certification

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/db"
	"time"
)

type (
	UpdateCertificationUseCase interface {
		Execute(context.Context, UpdateCertificationInput) (UpdateCertificationOutput, error)
	}

	UpdateCertificationInput struct {
		ID        string `validate:"required"`
		Name      string `json:"name" validate:"required"`
		Year      string `json:"year" validate:"required"`
		Comment   string `json:"comment"`
		SortIndex int    `json:"sort_index"`
		UserId    string `validate:"required"`
	}

	UpdateCertificationOutput struct {
		Certification *dto.CertificationDTO `json:"certification"`
	}

	UpdateCertificationPresenter interface {
		Output(certification *model.Certification) UpdateCertificationOutput
	}

	updateCertificationInteractor struct {
		certificationsRepository db.CertificationsRepository
		userRepository           db.UserRepository
		presenter                UpdateCertificationPresenter
		ctxTimeout               time.Duration
	}
)

func NewUpdateCertificationInteractor(
	certificationsRepository db.CertificationsRepository,
	userRepository db.UserRepository,
	presenter UpdateCertificationPresenter,
	t time.Duration,
) UpdateCertificationUseCase {
	return updateCertificationInteractor{
		certificationsRepository: certificationsRepository,
		userRepository:           userRepository,
		presenter:                presenter,
		ctxTimeout:               t,
	}
}

func (i updateCertificationInteractor) Execute(ctx context.Context, input UpdateCertificationInput) (UpdateCertificationOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	certification, err := i.certificationsRepository.FindByID(ctx, input.UserId, input.ID)

	if err != nil {
		return UpdateCertificationOutput{}, err
	}

	yearTime, err := time.Parse("2006-01-02", input.Year)
	if err != nil {
		return UpdateCertificationOutput{}, err
	}

	if err := certification.UpdateName(input.Name); err != nil {
		return UpdateCertificationOutput{}, err
	}

	if err := certification.UpdateYear(yearTime); err != nil {
		return UpdateCertificationOutput{}, err
	}

	if err := certification.UpdateComment(input.Comment); err != nil {
		return UpdateCertificationOutput{}, err
	}

	if err := certification.UpdateSortIndex(input.SortIndex); err != nil {
		return UpdateCertificationOutput{}, err
	}

	if err := i.certificationsRepository.Update(ctx, certification, input.UserId); err != nil {
		return UpdateCertificationOutput{}, err
	}

	return i.presenter.Output(certification), nil
}
