package file_uploader

import (
	"fmt"
	"os"
)

type FileUploaderConfig struct {
	endpoint        string
	accessKeyID     string
	secretAccessKey string
}

func NewFileUploaderConfig() *FileUploaderConfig {
	return &FileUploaderConfig{
		endpoint:        fmt.Sprintf("http://%s:%s", os.Getenv("MINIO_HOST"), os.Getenv("MINIO_PORT")),
		accessKeyID:     os.Getenv("MINIO_ACCESS_KEY"),
		secretAccessKey: os.Getenv("MINIO_SECRET_KEY"),
	}
}
