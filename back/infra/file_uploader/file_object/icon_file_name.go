package file_object

import "fmt"

type IconFileName struct {
	fileName   string
	objectName string
}

func NewIconFileName(fileName string) *IconFileName {
	return &IconFileName{
		fileName:   fileName,
		objectName: fmt.Sprintf("icon/%s", fileName),
	}
}

func (f *IconFileName) GetFileName() string {
	return f.fileName
}

func (f *IconFileName) GetObjectName() string {
	return f.objectName
}
