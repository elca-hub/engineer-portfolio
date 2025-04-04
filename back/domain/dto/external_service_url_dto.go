package dto

type ExternalServiceUrlDTO struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func NewExternalServiceUrlDTO(name string, url string) *ExternalServiceUrlDTO {
	return &ExternalServiceUrlDTO{
		Name: name,
		Url:  url,
	}
}
