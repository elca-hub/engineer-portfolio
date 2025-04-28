package model

import (
	"errors"
	"fmt"
	"strings"
)

type FileIconName struct {
	fileName   string
	objectName string
}

const (
	ICON_PATH = iota
	HEADER_PATH
)

func NewFileIconName(fileName string, path int) (*FileIconName, error) {
	if strings.HasPrefix(fileName, "icon/") || strings.HasPrefix(fileName, "header/") {
		fileNameTmp := fileName
		// fileNameから'icon/'や'header/'を削除
		fileNameTmp = strings.TrimPrefix(fileNameTmp, "icon/")
		fileNameTmp = strings.TrimPrefix(fileNameTmp, "header/")

		return &FileIconName{
			fileName:   fileNameTmp,
			objectName: fileName,
		}, nil
	}

	switch path {
	case ICON_PATH:
		return &FileIconName{
			fileName:   fileName,
			objectName: fmt.Sprintf("%s/%s", "icon", fileName),
		}, nil
	case HEADER_PATH:
		return &FileIconName{
			fileName:   fileName,
			objectName: fmt.Sprintf("%s/%s", "header", fileName),
		}, nil
	default:
		return nil, errors.New("invalid path")
	}
}

func (f *FileIconName) GetFileName() string {
	return f.fileName
}

func (f *FileIconName) GetObjectName() string {
	return f.objectName
}
