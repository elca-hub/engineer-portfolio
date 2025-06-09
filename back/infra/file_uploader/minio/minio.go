package minio

import (
	"bytes"
	"context"
	"devport/infra/file_uploader/file_object"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Minio struct {
	uploader   *manager.Uploader
	client     *s3.Client
	bucketName string
}

func NewMinio(uploader *manager.Uploader, client *s3.Client, bucketName string) *Minio {
	return &Minio{
		uploader:   uploader,
		client:     client,
		bucketName: bucketName,
	}
}

func (m *Minio) UploadFile(ctx context.Context, objectName file_object.ObjectNameInterface) error {
	fileBytes := objectName.GetFileContent()
	objectNameStr := objectName.GetObjectName()
	filePattern := objectName.GetFilePattern()

	tagging := fmt.Sprintf("category=%d", filePattern)

	_, err := m.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:  &m.bucketName,
		Key:     &objectNameStr,
		Body:    bytes.NewReader(fileBytes),
		Tagging: &tagging,
	})

	return err
}

func (m *Minio) DeleteFile(ctx context.Context, fileName file_object.FileNameInterface) error {
	fileNameStr := fileName.GetFileName()
	_, err := m.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &m.bucketName,
		Key:    &fileNameStr,
	})

	if err != nil {
		return err
	}

	return nil
}

func (m *Minio) GetFile(ctx context.Context, fileName file_object.FileNameInterface) ([]byte, error) {
	fileNameStr := fileName.GetFileName()
	result, err := m.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &m.bucketName,
		Key:    &fileNameStr,
	})

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
