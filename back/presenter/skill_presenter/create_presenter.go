package skill_presenter

import (
	"devport/domain/dto"
	"devport/domain/model"
	"devport/usecase/skill"
)

type SkillPresenter struct{}

func NewCreateSkillPresenter() skill.CreateSkillPresenter {
	return SkillPresenter{}
}

func (p SkillPresenter) Output(s *model.Skill) skill.CreateSkillOutput {
	skillDto := dto.NewSkillDTO(s)
	return skill.CreateSkillOutput{
		Skill: skillDto,
	}
}