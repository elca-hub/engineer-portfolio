package file_object

import "fmt"

type IconObjectName struct {
	fileName   string
	objectName string
}

func NewIconObjectName(fileName string) *IconObjectName {
	return &IconObjectName{
		objectName: fmt.Sprintf("icon/%s", fileName),
	}
}

func (f *IconObjectName) GetObjectName() string {
	return f.objectName
}

func (f *IconObjectName) GetFileName() string {
	return f.fileName
}

func (f *IconObjectName) GetFilePattern() uint {
	return IconObjectType
}
