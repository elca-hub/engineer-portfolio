package work_presenter

import (
	"devport/usecase/work"
)

type UploadWorkThumbnailPresenter struct{}

func NewUploadWorkThumbnailPresenter() work.UploadWorkThumbnailPresenter {
	return UploadWorkThumbnailPresenter{}
}

func (p UploadWorkThumbnailPresenter) Output(imageUrl string) work.UploadWorkThumbnailOutput {
	return work.UploadWorkThumbnailOutput{
		ImageUrl: imageUrl,
	}
}
