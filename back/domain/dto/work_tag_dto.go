package dto

type WorkTagDTO struct {
	Name string `json:"name"`
}

func NewWorkTagDTO(tag string) *WorkTagDTO {
	return &WorkTagDTO{
		Name: tag,
	}
}
