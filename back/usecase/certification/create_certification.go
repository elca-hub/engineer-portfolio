package certification

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/db"
	"time"
)

type (
	CreateCertificationUseCase interface {
		Execute(context.Context, CreateCertificationInput) (CreateCertificationOutput, error)
	}

	CreateCertificationInput struct {
		Name    string `json:"name" validate:"required"`
		Year    string `json:"year" validate:"required"`
		Comment string `json:"comment"`
		UserId  string `validate:"required"`
	}

	CreateCertificationOutput struct {
		Certification *dto.CertificationDTO `json:"certification"`
	}

	CreateCertificationPresenter interface {
		Output(certification *model.Certification) CreateCertificationOutput
	}

	createCertificationInteractor struct {
		certificationsRepository db.CertificationsRepository
		userRepository           db.UserRepository
		presenter                CreateCertificationPresenter
		ctxTimeout               time.Duration
	}
)

func NewCreateCertificationInteractor(
	certificationsRepository db.CertificationsRepository,
	userRepository db.UserRepository,
	presenter CreateCertificationPresenter,
	t time.Duration,
) CreateCertificationUseCase {
	return createCertificationInteractor{
		certificationsRepository: certificationsRepository,
		userRepository:           userRepository,
		presenter:                presenter,
		ctxTimeout:               t,
	}
}

func (i createCertificationInteractor) Execute(ctx context.Context, input CreateCertificationInput) (CreateCertificationOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	var certification *model.Certification

	yearTime, err := time.Parse("2006-01-02", input.Year)
	if err != nil {
		return CreateCertificationOutput{}, err
	}

	_, err = i.userRepository.FindById(ctx, input.UserId, nil)
	if err != nil {
		return CreateCertificationOutput{}, err
	}

	maxSortIndex, err := i.certificationsRepository.GetMaxSortIndex(ctx, input.UserId)
	if err != nil {
		return CreateCertificationOutput{}, err
	}
	sortIndex := maxSortIndex + 1

	certification, err = model.NewCertification(input.Name, yearTime, input.Comment, sortIndex)
	if err != nil {
		return CreateCertificationOutput{}, err
	}

	err = i.certificationsRepository.Create(ctx, certification, input.UserId)

	if err != nil {
		return CreateCertificationOutput{}, err
	}

	return i.presenter.Output(certification), nil
}