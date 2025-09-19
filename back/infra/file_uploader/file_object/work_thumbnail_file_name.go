package file_object

import "fmt"

type WorkThumbnailFileName struct {
	fileName   string
	objectName string
}

const (
	WorkThumbnailUri = "work_thumbnails"
)

func WorkThumbnailWorkPrefix(userId string, workId string) string {
	return fmt.Sprintf("%s/%s/%s/thumbnail/", WorkThumbnailUri, userId, workId)
}

func NewWorkThumbnailFileName(userId string, workId string, fileName string) *WorkThumbnailFileName {
	return &WorkThumbnailFileName{
		fileName:   fileName,
		objectName: fmt.Sprintf("%s%s", WorkThumbnailWorkPrefix(userId, workId), fileName),
	}
}

func (f *WorkThumbnailFileName) GetFileName() string {
	return f.fileName
}

func (f *WorkThumbnailFileName) GetObjectName() string {
	return f.objectName
}
