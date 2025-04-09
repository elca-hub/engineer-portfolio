package dto

import "devport/domain/model"

type SkillDTO struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	When      string `json:"when"`
	SortIndex int    `json:"sort_index"`
}

func NewSkillDTO(skill *model.Skill) *SkillDTO {
	return &SkillDTO{
		Name:      skill.Name(),
		Status:    skill.Status(),
		When:      skill.When().Format("2006-01-02"),
		SortIndex: skill.SortIndex(),
	}
}
