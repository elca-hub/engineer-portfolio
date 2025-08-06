package skill

import (
	"context"
	"devport/domain/repo/db"
	"time"
)

type (
	UpdateSkillsSortUseCase interface {
		Execute(context.Context, UpdateSkillsSortInput) (UpdateSkillsSortOutput, error)
	}

	UpdateSkillsSortInput struct {
		UserId string `validate:"required"`

		SortList []struct {
			ID        string `validate:"required" json:"id"`
			SortIndex int    `json:"sort_index"`
		} `json:"sort_list" validate:"required,dive"`
	}

	UpdateSkillsSortOutput struct {
	}

	UpdateSkillsSortPresenter interface {
		Output() UpdateSkillsSortOutput
	}

	updateSkillsSortInteractor struct {
		skillsRepository db.SkillsRepository
		userRepository   db.UserRepository
		presenter        UpdateSkillsSortPresenter
		ctxTimeout       time.Duration
	}
)

func NewUpdateSkillsSortInteractor(
	skillsRepository db.SkillsRepository,
	userRepository db.UserRepository,
	presenter UpdateSkillsSortPresenter,
	t time.Duration,
) UpdateSkillsSortUseCase {
	return updateSkillsSortInteractor{
		skillsRepository: skillsRepository,
		userRepository:   userRepository,
		presenter:        presenter,
		ctxTimeout:       t,
	}
}

func (i updateSkillsSortInteractor) Execute(ctx context.Context, input UpdateSkillsSortInput) (UpdateSkillsSortOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	err := i.skillsRepository.WithTransaction(ctx, func(ctx context.Context) error {
		for _, skillSort := range input.SortList {
			skill, err := i.skillsRepository.FindByID(ctx, input.UserId, skillSort.ID)

			if err != nil {
				return err
			}

			if err := skill.UpdateSortIndex(skillSort.SortIndex); err != nil {
				return err
			}

			if err := i.skillsRepository.Update(ctx, skill, input.UserId); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return UpdateSkillsSortOutput{}, err
	}

	return i.presenter.Output(), nil
}
