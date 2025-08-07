package file_object

import "fmt"

type WorkImageFileName struct {
	fileName   string
	objectName string
}

const (
	WorkImageUri = "work_images"
)

func WorkImageWorkPrefix(userId string, workId string) string {
	return fmt.Sprintf("%s/%s/%s/", WorkImageUri, userId, workId)
}

func NewWorkImageFileName(userId string, workId string, fileName string) *WorkImageFileName {
	return &WorkImageFileName{
		fileName:   fileName,
		objectName: fmt.Sprintf("%s%s", WorkImageWorkPrefix(userId, workId), fileName),
	}
}

func (f *WorkImageFileName) GetFileName() string {
	return f.fileName
}

func (f *WorkImageFileName) GetObjectName() string {
	return f.objectName
}