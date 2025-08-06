package skill_presenter

import (
	"devport/domain/dto"
	"devport/domain/model"
	"devport/usecase/skill"
)

type UpdateSkillPresenter struct{}

func NewUpdateSkillPresenter() skill.UpdateSkillPresenter {
	return UpdateSkillPresenter{}
}

func (p UpdateSkillPresenter) Output(s *model.Skill) skill.UpdateSkillOutput {
	skillDto := dto.NewSkillDTO(s)
	return skill.UpdateSkillOutput{
		Skill: skillDto,
	}
}