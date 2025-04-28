package dto

import (
	"devport/domain/model"
)

type UserDTO struct {
	Email            string `json:"email"`
	UserId           string `json:"user_id"`
	Name             string `json:"name"`
	IconName         string `json:"icon_name"`
	HeaderIconName   string `json:"header_icon_name"`
	BioPath          string `json:"bio_path"`
	OrganizationName string `json:"organization_name"`
	OccupationName   string `json:"occupation_name"`
	Place            string `json:"place"`
	Birthday         string `json:"birthday"`
}

func NewUserDTO(userModel *model.User) *UserDTO {
	return &UserDTO{
		UserId:           userModel.ID(),
		Email:            userModel.Email().Email(),
		Name:             userModel.Name(),
		IconName:         userModel.IconName(),
		HeaderIconName:   userModel.HeaderIconName(),
		Birthday:         userModel.Birthday().Format("2006-01-02"),
		BioPath:          userModel.BioPath(),
		OrganizationName: userModel.OrganizationName(),
		OccupationName:   userModel.OccupationName(),
		Place:            userModel.Place(),
	}
}
