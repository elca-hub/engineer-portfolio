package user_presenter

import (
	usermodel "devport/domain/model"
	"devport/usecase/user"
)

type GetUserInfoPresenter struct{}

func NewGetUserInfoPresenter() *GetUserInfoPresenter {
	return &GetUserInfoPresenter{}
}

func (p *GetUserInfoPresenter) Output(model usermodel.User) user.GetUserInfoOutput {
	emailModel := model.Email()
	var email string
	if emailModel == nil {
		email = ""
	} else {
		email = emailModel.Email()
	}

	skillOutput := make([]struct {
		Name      string `json:"name"`
		Status    string `json:"status"`
		When      string `json:"when"`
		SortIndex int    `json:"sort_index"`
	}, 0, len(model.Skills()))

	for _, skill := range model.Skills() {
		skillOutput = append(skillOutput, struct {
			Name      string `json:"name"`
			Status    string `json:"status"`
			When      string `json:"when"`
			SortIndex int    `json:"sort_index"`
		}{
			Name:      skill.Name(),
			Status:    skill.Status(),
			When:      skill.When().String(),
			SortIndex: skill.SortIndex(),
		})
	}

	externalserviceurlOutput := make([]struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}, 0, len(model.ExternalServiceURLs()))

	for _, externalserviceurl := range model.ExternalServiceURLs() {
		externalserviceurlOutput = append(externalserviceurlOutput, struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}{
			Name: externalserviceurl.Name(),
			Url:  externalserviceurl.Url(),
		})
	}

	return user.GetUserInfoOutput{
		Email:              email,
		Name:               model.Name(),
		IconPath:           model.IconPath(),
		HeaderPath:         model.HeaderPath(),
		BioPath:            model.BioPath(),
		Skills:             skillOutput,
		ExternalServiceUrl: externalserviceurlOutput,
	}
}
