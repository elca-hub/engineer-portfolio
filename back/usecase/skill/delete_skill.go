package skill

import (
	"context"
	"devport/domain/repo/db"
	"time"
)

type (
	DeleteSkillUseCase interface {
		Execute(context.Context, DeleteSkillInput) error
	}

	DeleteSkillInput struct {
		ID     string `validate:"required"`
		UserId string `validate:"required"`
	}

	deleteSkillInteractor struct {
		skillsRepository db.SkillsRepository
		userRepository   db.UserRepository
		ctxTimeout       time.Duration
	}
)

func NewDeleteSkillInteractor(
	skillsRepository db.SkillsRepository,
	userRepository db.UserRepository,
	t time.Duration,
) DeleteSkillUseCase {
	return deleteSkillInteractor{
		skillsRepository: skillsRepository,
		userRepository:   userRepository,
		ctxTimeout:       t,
	}
}

func (i deleteSkillInteractor) Execute(ctx context.Context, input DeleteSkillInput) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	_, err := i.userRepository.FindById(ctx, input.UserId, nil)
	if err != nil {
		return err
	}

	return i.skillsRepository.Delete(ctx, input.ID)
}