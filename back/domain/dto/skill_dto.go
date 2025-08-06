package dto

import "devport/domain/model"

type SkillDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Year      string `json:"year"`
	Comment   string `json:"comment"`
	SortIndex int    `json:"sort_index"`
}

func NewSkillDTO(skill *model.Skill) *SkillDTO {
	return &SkillDTO{
		ID:        skill.ID(),
		Name:      skill.Name(),
		Year:      skill.Year().Format("2006-01-02"),
		Comment:   skill.Comment(),
		SortIndex: skill.SortIndex(),
	}
}
