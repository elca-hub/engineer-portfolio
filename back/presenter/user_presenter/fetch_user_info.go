package user_presenter

import (
	"devport/domain/dto"
	usermodel "devport/domain/model"
	"devport/usecase/user"
	usecase "devport/usecase/user"
)

type FetchUserInfoPresenter struct{}

func NewFetchUserInfoPresenter() *FetchUserInfoPresenter {
	return &FetchUserInfoPresenter{}
}

func (p *FetchUserInfoPresenter) Output(user usermodel.User) user.FetchUserInfoOutput {
	skills := make([]dto.SkillDTO, 0, len(user.Skills()))

	for _, skill := range user.Skills() {
		skills = append(skills, dto.SkillDTO{
			Name:      skill.Name(),
			Status:    skill.Status(),
			When:      skill.When().String(),
			SortIndex: skill.SortIndex(),
		})
	}

	externalServiceURLs := make([]dto.ExternalServiceUrlDTO, 0, len(user.ExternalServiceURLs()))
	for _, externalServiceURL := range user.ExternalServiceURLs() {
		externalServiceURLs = append(externalServiceURLs, dto.ExternalServiceUrlDTO{
			Name: externalServiceURL.Name(),
			Url:  externalServiceURL.Url(),
		})
	}

	userDto := dto.NewUserDTO(
		user.Email().Email(),
		user.ID(),
		user.Name(),
		user.Birthday(),
		user.IconName(),
		user.HeaderIconName(),
		user.BioPath(),
		user.OrganizationName(),
		skills,
		externalServiceURLs,
	)

	return usecase.FetchUserInfoOutput{
		User: userDto,
	}
}
