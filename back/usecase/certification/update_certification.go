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

	var certification *model.Certification

	_, err := i.userRepository.FindById(ctx, input.UserId, nil)
	if err != nil {
		return UpdateCertificationOutput{}, err
	}

	yearTime, err := time.Parse("2006-01-02", input.Year)
	if err != nil {
		return UpdateCertificationOutput{}, err
	}

	certification, err = model.NewCertification(input.Name, yearTime, input.Comment, input.SortIndex)
	if err != nil {
		return UpdateCertificationOutput{}, err
	}

	err = i.certificationsRepository.Update(ctx, certification, input.UserId)

	if err != nil {
		return UpdateCertificationOutput{}, err
	}

	return i.presenter.Output(certification), nil
}