package work_presenter

import (
	"devport/usecase/work"
)

type UploadWorkImagePresenter struct{}

func NewUploadWorkImagePresenter() work.UploadWorkImagePresenter {
	return UploadWorkImagePresenter{}
}

func (p UploadWorkImagePresenter) Output(imageUrl string) work.UploadWorkImageOutput {
	return work.UploadWorkImageOutput{
		ImageUrl: imageUrl,
	}
}