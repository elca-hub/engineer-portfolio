package dto

import "devport/domain/model"

type CertificationDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Year      string `json:"year"`
	Comment   string `json:"comment"`
	SortIndex int    `json:"sort_index"`
}

func NewCertificationDTO(certification *model.Certification) *CertificationDTO {
	return &CertificationDTO{
		ID:        certification.ID(),
		Name:      certification.Name(),
		Year:      certification.Year().Format("2006-01-02"),
		Comment:   certification.Comment(),
		SortIndex: certification.SortIndex(),
	}
}