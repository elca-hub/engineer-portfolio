package skill

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/db"
	"time"
)

type (
	FindByUserSkillUseCase interface {
		Execute(context.Context, FindByUserSkillInput) (FindByUserSkillOutput, error)
	}

	FindByUserSkillInput struct {
		UserId string `validate:"required"`
	}

	FindByUserSkillOutput struct {
		Skills []*dto.SkillDTO `json:"skills"`
	}

	FindByUserSkillPresenter interface {
		Output(skills []*model.Skill) FindByUserSkillOutput
	}

	findByUserSkillInteractor struct {
		skillsRepository db.SkillsRepository
		userRepository   db.UserRepository
		presenter        FindByUserSkillPresenter
		ctxTimeout       time.Duration
	}
)

func NewFindByUserSkillInteractor(
	skillsRepository db.SkillsRepository,
	userRepository db.UserRepository,
	presenter FindByUserSkillPresenter,
	t time.Duration,
) FindByUserSkillUseCase {
	return findByUserSkillInteractor{
		skillsRepository: skillsRepository,
		userRepository:   userRepository,
		presenter:        presenter,
		ctxTimeout:       t,
	}
}

func (i findByUserSkillInteractor) Execute(ctx context.Context, input FindByUserSkillInput) (FindByUserSkillOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	_, err := i.userRepository.FindById(ctx, input.UserId, nil)
	if err != nil {
		return FindByUserSkillOutput{}, err
	}

	skills, err := i.skillsRepository.FindByUserID(ctx, input.UserId)
	if err != nil {
		return FindByUserSkillOutput{}, err
	}

	return i.presenter.Output(skills), nil
}