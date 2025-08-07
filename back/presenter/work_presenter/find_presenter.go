package work_presenter

import (
	"devport/domain/dto"
	"devport/domain/model"
	"devport/usecase/work"
)

type FindWorkPresenter struct{}

func NewFindWorkPresenter() work.FindWorkPresenter {
	return FindWorkPresenter{}
}

func (p FindWorkPresenter) Output(workModel *model.Work) work.FindWorkOutput {
	return work.FindWorkOutput{
		Work: dto.NewWorkDTO(workModel),
	}
}
