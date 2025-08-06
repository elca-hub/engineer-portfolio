package skill_presenter

import (
	"devport/domain/dto"
	"devport/domain/model"
	"devport/usecase/skill"
)

type FindByUserSkillPresenter struct{}

func NewFindByUserSkillPresenter() skill.FindByUserSkillPresenter {
	return FindByUserSkillPresenter{}
}

func (p FindByUserSkillPresenter) Output(skills []*model.Skill) skill.FindByUserSkillOutput {
	skillDtos := make([]*dto.SkillDTO, len(skills))
	for i, s := range skills {
		skillDtos[i] = dto.NewSkillDTO(s)
	}
	return skill.FindByUserSkillOutput{
		Skills: skillDtos,
	}
}