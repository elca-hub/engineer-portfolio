package dto

import "time"

type UserDTO struct {
	Email              string                  `json:"email"`
	UserId             string                  `json:"user_id"`
	Name               string                  `json:"name"`
	IconName           string                  `json:"icon_name"`
	HeaderIconName     string                  `json:"header_icon_name"`
	BioPath            string                  `json:"bio_path"`
	OrganizationName   string                  `json:"organization_name"`
	Skills             []SkillDTO              `json:"skills"`
	Birthday           string                  `json:"birthday"`
	ExternalServiceUrl []ExternalServiceUrlDTO `json:"external_service_url"`
}

func NewUserDTO(
	email string,
	userId string,
	name string,
	birthday time.Time,
	iconName string,
	headerIconName string,
	bioPath string,
	organizationName string,
	skills []SkillDTO,
	externalServiceUrl []ExternalServiceUrlDTO,
) *UserDTO {
	return &UserDTO{
		Email:              email,
		UserId:             userId,
		Name:               name,
		Birthday:           birthday.Format("2006-01-02"),
		IconName:           iconName,
		HeaderIconName:     headerIconName,
		BioPath:            bioPath,
		OrganizationName:   organizationName,
		Skills:             skills,
		ExternalServiceUrl: externalServiceUrl,
	}
}
