package file_object

import "fmt"

type HeaderObjectName struct {
	fileName   string
	objectName string
}

func NewHeaderObjectName(fileName string) *HeaderObjectName {
	return &HeaderObjectName{
		objectName: fmt.Sprintf("header/%s", fileName),
	}
}

func (f *HeaderObjectName) GetObjectName() string {
	return f.objectName
}
