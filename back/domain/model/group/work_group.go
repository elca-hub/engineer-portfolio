package group

import (
	"devport/domain/model"
	"fmt"
)

const (
	MAX_PUBLISH_WORK_COUNT = 50
	MAX_DRAFT_WORK_COUNT   = 100
)

type WorkGroup struct {
	works       []*model.Work
	publicWorks []*model.Work
	draftWorks  []*model.Work
}

func NewWorkGroup(works []*model.Work) (*WorkGroup, error) {
	publishWorks := make([]*model.Work, 0, MAX_PUBLISH_WORK_COUNT)
	draftWorks := make([]*model.Work, 0, MAX_DRAFT_WORK_COUNT)

	for _, work := range works {
		if !work.IsDraft() {
			publishWorks = append(publishWorks, work)
		} else {
			draftWorks = append(draftWorks, work)
		}
	}

	if len(publishWorks) > MAX_PUBLISH_WORK_COUNT {
		return nil, fmt.Errorf("公開作品事例は%d件までです", MAX_PUBLISH_WORK_COUNT)
	}

	if len(draftWorks) > MAX_DRAFT_WORK_COUNT {
		return nil, fmt.Errorf("下書き作品事例は%d件までです", MAX_DRAFT_WORK_COUNT)
	}

	return &WorkGroup{
		works:       publishWorks,
		publicWorks: publishWorks,
		draftWorks:  draftWorks,
	}, nil
}

func (w *WorkGroup) Works() []*model.Work {
	return w.works
}

func (w *WorkGroup) PublicWorks() []*model.Work {
	return w.publicWorks
}

func (w *WorkGroup) DraftWorks() []*model.Work {
	return w.draftWorks
}

func (w *WorkGroup) AddWork(work *model.Work) error {
	if !work.IsDraft() {
		if len(w.publicWorks) >= MAX_PUBLISH_WORK_COUNT {
			return fmt.Errorf("公開作品事例は%d件までです", MAX_PUBLISH_WORK_COUNT)
		}

		w.publicWorks = append(w.publicWorks, work)
	} else {
		if len(w.draftWorks) >= MAX_DRAFT_WORK_COUNT {
			return fmt.Errorf("下書き作品事例は%d件までです", MAX_DRAFT_WORK_COUNT)
		}
		w.draftWorks = append(w.draftWorks, work)
	}

	w.works = append(w.works, work)

	return nil
}
