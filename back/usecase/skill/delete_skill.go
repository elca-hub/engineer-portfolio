package skill

import (
	"context"
	"devport/domain/repo/db"
	"errors"
	"time"
)

type (
	DeleteSkillUseCase interface {
		Execute(context.Context, DeleteSkillInput) (DeleteSkillOutput, error)
	}

	DeleteSkillInput struct {
		ID     string `validate:"required"`
		UserId string `validate:"required"`
	}

	DeleteSkillOutput struct{}

	DeleteSkillPresenter interface {
		Output() DeleteSkillOutput
	}

	deleteSkillInteractor struct {
		skillsRepository db.SkillsRepository
		userRepository   db.UserRepository
		presenter        DeleteSkillPresenter
		ctxTimeout       time.Duration
	}
)

func NewDeleteSkillInteractor(
	skillsRepository db.SkillsRepository,
	userRepository db.UserRepository,
	presenter DeleteSkillPresenter,
	t time.Duration,
) DeleteSkillUseCase {
	return deleteSkillInteractor{
		skillsRepository: skillsRepository,
		userRepository:   userRepository,
		presenter:        presenter,
		ctxTimeout:       t,
	}
}

func (i deleteSkillInteractor) Execute(ctx context.Context, input DeleteSkillInput) (DeleteSkillOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	if isExists, err := i.userRepository.ExistsById(ctx, input.UserId); err != nil || !isExists {
		if err != nil {
			return DeleteSkillOutput{}, err
		}
		return DeleteSkillOutput{}, errors.New("user not found")
	}

	err := i.skillsRepository.Delete(ctx, input.UserId, input.ID)

	if err != nil {
		return DeleteSkillOutput{}, err
	}

	return i.presenter.Output(), nil
}
