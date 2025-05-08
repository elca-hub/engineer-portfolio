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
	BIO_IMAGE_PATH
)

func NewFileIconName(fileName string, path int) (*FileIconName, error) {
	prefixes := []string{"icon/", "header/", "bio_images/"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(fileName, prefix) {
			return &FileIconName{
				fileName:   strings.TrimPrefix(fileName, prefix),
				objectName: fileName,
			}, nil
		}
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
	case BIO_IMAGE_PATH:
		return &FileIconName{
			fileName:   fileName,
			objectName: fmt.Sprintf("%s/%s", "bio_images", fileName),
		}, nil
	default:
		return nil, errors.New("不正なパスです")
	}
}

func (f *FileIconName) GetFileName() string {
	return f.fileName
}

func (f *FileIconName) GetObjectName() string {
	return f.objectName
}
