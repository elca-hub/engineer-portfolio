package skill_presenter

import (
	"devport/usecase/skill"
)

type UpdateSkillsSortPresenter struct{}

func NewUpdateSkillsSortPresenter() skill.UpdateSkillsSortPresenter {
	return UpdateSkillsSortPresenter{}
}

func (p UpdateSkillsSortPresenter) Output() skill.UpdateSkillsSortOutput {
	return skill.UpdateSkillsSortOutput{}
}
