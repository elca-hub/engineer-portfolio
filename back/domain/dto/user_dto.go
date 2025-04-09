package dto

import (
	"devport/domain/model"
)

type UserDTO struct {
	Email              string                   `json:"email"`
	UserId             string                   `json:"user_id"`
	Name               string                   `json:"name"`
	IconName           string                   `json:"icon_name"`
	HeaderIconName     string                   `json:"header_icon_name"`
	BioPath            string                   `json:"bio_path"`
	OrganizationName   string                   `json:"organization_name"`
	OccupationName     string                   `json:"occupation_name"`
	Skills             []*SkillDTO              `json:"skills"`
	Birthday           string                   `json:"birthday"`
	ExternalServiceUrl []*ExternalServiceUrlDTO `json:"external_service_url"`
}

func NewUserDTO(userModel *model.User) *UserDTO {
	skillsDto := make([]*SkillDTO, len(userModel.Skills()))

	for i, skillModel := range userModel.Skills() {
		skillsDto[i] = NewSkillDTO(skillModel)
	}

	externalDto := make([]*ExternalServiceUrlDTO, len(userModel.ExternalServiceURLs()))

	for i, externalModel := range userModel.ExternalServiceURLs() {
		externalDto[i] = NewExternalServiceUrlDTO(externalModel)
	}

	return &UserDTO{
		UserId:             userModel.ID(),
		Email:              userModel.Email().Email(),
		Name:               userModel.Name(),
		IconName:           userModel.IconName(),
		HeaderIconName:     userModel.HeaderIconName(),
		Birthday:           userModel.Birthday().Format("2006-01-02"),
		BioPath:            userModel.BioPath(),
		OrganizationName:   userModel.OrganizationName(),
		OccupationName:     userModel.OccupationName(),
		Skills:             skillsDto,
		ExternalServiceUrl: externalDto,
	}
}
