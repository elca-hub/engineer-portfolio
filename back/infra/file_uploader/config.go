package file_uploader

import (
	"fmt"
	"os"
)

type FileUploaderConfig struct {
	endpoint        string
	publicEndpoint  string
	accessKeyID     string
	secretAccessKey string
	bucketName      string
}

func NewFileUploaderConfig() *FileUploaderConfig {
	return &FileUploaderConfig{
		endpoint:        fmt.Sprintf("http://%s:%s", os.Getenv("MINIO_HOST"), os.Getenv("MINIO_PORT")),
		publicEndpoint:  os.Getenv("BIO_RESOURCE_DOMAIN"),
		accessKeyID:     os.Getenv("MINIO_ACCESS_KEY"),
		secretAccessKey: os.Getenv("MINIO_SECRET_KEY"),
		bucketName:      os.Getenv("MINIO_BUCKET_NAME"),
	}
}
