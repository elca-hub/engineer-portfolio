package skill

import (
	"context"
	"devport/domain/dto"
	"devport/domain/model"
	"devport/domain/repo/db"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	CreateSkillUseCase interface {
		Execute(context.Context, CreateSkillInput) (CreateSkillOutput, error)
	}

	CreateSkillInput struct {
		Name    string `json:"name" validate:"required"`
		Year    string `json:"year" validate:"required"`
		Comment string `json:"comment"`
		UserId  string `validate:"required"`
	}

	CreateSkillOutput struct {
		Skill *dto.SkillDTO `json:"skill"`
	}

	CreateSkillPresenter interface {
		Output(skill *model.Skill) CreateSkillOutput
	}

	createSkillInteractor struct {
		skillsRepository db.SkillsRepository
		userRepository   db.UserRepository
		presenter        CreateSkillPresenter
		ctxTimeout       time.Duration
	}
)

func NewCreateSkillInteractor(
	skillsRepository db.SkillsRepository,
	userRepository db.UserRepository,
	presenter CreateSkillPresenter,
	t time.Duration,
) CreateSkillUseCase {
	return createSkillInteractor{
		skillsRepository: skillsRepository,
		userRepository:   userRepository,
		presenter:        presenter,
		ctxTimeout:       t,
	}
}

func (i createSkillInteractor) Execute(ctx context.Context, input CreateSkillInput) (CreateSkillOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	var skill *model.Skill

	yearTime, err := time.Parse("2006-01-02", input.Year)
	if err != nil {
		return CreateSkillOutput{}, err
	}

	if isExists, err := i.userRepository.ExistsById(ctx, input.UserId); err != nil || !isExists {
		if err != nil {
			return CreateSkillOutput{}, err
		}
		if !isExists {
			return CreateSkillOutput{}, errors.New("ユーザーが存在しません")
		}
	}

	maxSortIndex, err := i.skillsRepository.GetMaxSortIndex(ctx, input.UserId)
	if err != nil {
		return CreateSkillOutput{}, err
	}
	sortIndex := maxSortIndex + 1

	id, err := uuid.NewRandom()

	if err != nil {
		return CreateSkillOutput{}, err
	}

	skill, err = model.NewSkillWithID(id.String(), input.Name, yearTime, input.Comment, sortIndex)
	if err != nil {
		return CreateSkillOutput{}, err
	}

	err = i.skillsRepository.Create(ctx, skill, input.UserId)

	if err != nil {
		return CreateSkillOutput{}, err
	}

	return i.presenter.Output(skill), nil
}
