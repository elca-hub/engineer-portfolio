package work_presenter

import (
	usecase "devport/usecase/work"
)

type CreateWorkPresenter struct{}

func NewCreateWorkPresenter() *CreateWorkPresenter {
	return &CreateWorkPresenter{}
}

func (p *CreateWorkPresenter) Output(title string) usecase.CreateWorkOutput {
	return usecase.CreateWorkOutput{
		Title: title,
	}
}
