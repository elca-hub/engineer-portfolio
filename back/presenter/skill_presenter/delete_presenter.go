package skill_presenter

import (
	"devport/usecase/skill"
)

type DeleteSkillPresenter struct{}

func NewDeleteSkillPresenter() skill.DeleteSkillPresenter {
	return DeleteSkillPresenter{}
}

func (p DeleteSkillPresenter) Output() skill.DeleteSkillOutput {
	return skill.DeleteSkillOutput{}
}
