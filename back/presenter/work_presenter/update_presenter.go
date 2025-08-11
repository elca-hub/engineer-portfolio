package work_presenter

import (
	"devport/domain/dto"
	"devport/domain/model"
	"devport/usecase/work"
)

type UpdateWorkPresenter struct{}

func NewUpdateWorkPresenter() work.UpdateWorkPresenter {
	return UpdateWorkPresenter{}
}

func (p UpdateWorkPresenter) Output(workModel *model.Work) work.UpdateWorkOutput {
	return work.UpdateWorkOutput{
		Work: dto.NewWorkDTO(workModel),
	}
}
