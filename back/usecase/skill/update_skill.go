package skill

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/db"
	"time"
)

type (
	UpdateSkillUseCase interface {
		Execute(context.Context, UpdateSkillInput) (UpdateSkillOutput, error)
	}

	UpdateSkillInput struct {
		UserId string `validate:"required"`

		ID        string `validate:"required"`
		Name      string `json:"name" validate:"required"`
		Year      string `json:"year" validate:"required"`
		Comment   string `json:"comment"`
		SortIndex int    `json:"sort_index"`
	}

	UpdateSkillOutput struct {
		Skill *dto.SkillDTO `json:"skill"`
	}

	UpdateSkillPresenter interface {
		Output(skill *model.Skill) UpdateSkillOutput
	}

	updateSkillInteractor struct {
		skillsRepository db.SkillsRepository
		userRepository   db.UserRepository
		presenter        UpdateSkillPresenter
		ctxTimeout       time.Duration
	}
)

func NewUpdateSkillInteractor(
	skillsRepository db.SkillsRepository,
	userRepository db.UserRepository,
	presenter UpdateSkillPresenter,
	t time.Duration,
) UpdateSkillUseCase {
	return updateSkillInteractor{
		skillsRepository: skillsRepository,
		userRepository:   userRepository,
		presenter:        presenter,
		ctxTimeout:       t,
	}
}

func (i updateSkillInteractor) Execute(ctx context.Context, input UpdateSkillInput) (UpdateSkillOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	println("input values:", input.UserId, input.ID, input.Name, input.Year, input.Comment, input.SortIndex)

	var skill *model.Skill

	skill, err := i.skillsRepository.FindByID(ctx, input.UserId, input.ID)
	if err != nil {
		return UpdateSkillOutput{}, err
	}

	yearTime, err := time.Parse("2006-01-02", input.Year)
	if err != nil {
		return UpdateSkillOutput{}, err
	}

	if err := skill.UpdateName(input.Name); err != nil {
		return UpdateSkillOutput{}, err
	}

	if err := skill.UpdateYear(yearTime); err != nil {
		return UpdateSkillOutput{}, err
	}

	if err := skill.UpdateComment(input.Comment); err != nil {
		return UpdateSkillOutput{}, err
	}

	if err := skill.UpdateSortIndex(input.SortIndex); err != nil {
		return UpdateSkillOutput{}, err
	}

	if err := i.skillsRepository.Update(ctx, skill, input.UserId); err != nil {
		return UpdateSkillOutput{}, err
	}

	return i.presenter.Output(skill), nil
}
