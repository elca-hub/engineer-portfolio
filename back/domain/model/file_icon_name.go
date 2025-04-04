package model

import "fmt"

type FileIconName struct {
	fileName   string
	objectName string
}

func NewFileIconName(fileName string) *FileIconName {
	return &FileIconName{
		fileName:   fileName,
		objectName: fmt.Sprintf("%s/%s", "icon", fileName),
	}
}

func (f *FileIconName) GetFileName() string {
	return f.fileName
}

func (f *FileIconName) GetObjectName() string {
	return f.objectName
}
