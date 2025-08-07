package work_presenter

import (
	"devport/domain/dto"
	"devport/domain/model"
	"devport/usecase/work"
)

type FetchWorkPresenter struct{}

func NewFetchWorkPresenter() work.FetchWorkPresenter {
	return FetchWorkPresenter{}
}

func (p FetchWorkPresenter) Output(works []*model.Work) work.FetchWorkOutput {
	dtos := make([]*dto.WorkDTO, len(works))
	for i, work := range works {
		dtos[i] = dto.NewWorkDTO(work)
	}

	return work.FetchWorkOutput{
		Works: dtos,
	}
}
