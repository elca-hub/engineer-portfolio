package dto

type UserDTO struct {
	Email              string                  `json:"email"`
	UserId             string                  `json:"user_id"`
	Name               string                  `json:"name"`
	IconName           string                  `json:"icon_name"`
	HeaderPath         string                  `json:"header_path"`
	BioPath            string                  `json:"bio_path"`
	Skills             []SkillDTO              `json:"skills"`
	ExternalServiceUrl []ExternalServiceUrlDTO `json:"external_service_url"`
}

func NewUserDTO(
	email string,
	userId string,
	name string,
	iconName string,
	headerPath string,
	bioPath string,
	skills []SkillDTO,
	externalServiceUrl []ExternalServiceUrlDTO,
) *UserDTO {
	return &UserDTO{
		Email:              email,
		UserId:             userId,
		Name:               name,
		IconName:           iconName,
		HeaderPath:         headerPath,
		BioPath:            bioPath,
		Skills:             skills,
		ExternalServiceUrl: externalServiceUrl,
	}
}
