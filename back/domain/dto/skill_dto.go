package dto

type SkillDTO struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	When      string `json:"when"`
	SortIndex int    `json:"sort_index"`
}

func NewSkillDTO(name string, status string, when string, sortIndex int) *SkillDTO {
	return &SkillDTO{
		Name:      name,
		Status:    status,
		When:      when,
		SortIndex: sortIndex,
	}
}
