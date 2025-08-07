package work_presenter

import (
	"devport/domain/dto"
	"devport/domain/model"
	"devport/usecase/work"
)

type CreateWorkPresenter struct{}

func NewCreateWorkPresenter() work.CreateWorkPresenter {
	return CreateWorkPresenter{}
}

func (p CreateWorkPresenter) Output(workModel *model.Work) work.CreateWorkOutput {
	return work.CreateWorkOutput{
		Work: dto.NewWorkDTO(workModel),
	}
}