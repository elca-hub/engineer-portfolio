package model

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type FileIconName struct {
	fileName   string
	objectName string
	iconType   int
}

const (
	ICON_PATH = iota
	HEADER_PATH
	BIO_IMAGE_PATH
	WORK_CONTENT_PATH
	WORK_THUMBNAIL_PATH
)

func NewFileName(fileName string, path int) (*FileIconName, error) {
	prefixes := []string{"icon/", "header/", "bio_images/", "works/"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(fileName, prefix) {
			return &FileIconName{
				fileName:   strings.TrimPrefix(fileName, prefix),
				objectName: fileName,
				iconType:   path,
			}, nil
		}
	}

	prefixes = []string{"icon", "header", "bio_images", "works"}
	if path < 0 || path >= len(prefixes) {
		return nil, errors.New("不正なパスです")
	}
	return &FileIconName{
		fileName:   fileName,
		objectName: fmt.Sprintf("%s/%s", prefixes[path], fileName),
		iconType:   path,
	}, nil
}

func (f *FileIconName) GetFileName() string {
	return f.fileName
}

func (f *FileIconName) GetObjectName() string {
	return f.objectName
}

func (f *FileIconName) GetPath() int {
	return f.iconType
}

func (f *FileIconName) GetUrl() string {
	protocol := "http"
	hostName := "localhost"
	port := os.Getenv("MINIO_PORT")
	bucketName := os.Getenv("MINIO_BUCKET_NAME")
	switch f.iconType {
	case ICON_PATH:
		return fmt.Sprintf("%s://%s:%s/%s/icon/%s", protocol, hostName, port, bucketName, f.fileName)
	case HEADER_PATH:
		return fmt.Sprintf("%s://%s:%s/%s/header/%s", protocol, hostName, port, bucketName, f.fileName)
	case BIO_IMAGE_PATH:
		return fmt.Sprintf("%s://%s:%s/%s/bio_images/%s", protocol, hostName, port, bucketName, f.fileName)
	case WORK_CONTENT_PATH, WORK_THUMBNAIL_PATH:
		return fmt.Sprintf("%s://%s:%s/%s/works/%s", protocol, hostName, port, bucketName, f.fileName)
	default:
		return ""
	}
}
