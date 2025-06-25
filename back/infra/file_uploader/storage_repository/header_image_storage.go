package storage_repository

import (
	"bytes"
	"context"
	"devport/infra/file_uploader/file_object"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type HeaderImageStorage struct {
	uploader   *manager.Uploader
	client     *s3.Client
	bucketName string
}

func NewHeaderImageStorage(uploader *manager.Uploader, client *s3.Client, bucketName string) *HeaderImageStorage {
	return &HeaderImageStorage{
		uploader,
		client,
		bucketName,
	}
}

func (r *HeaderImageStorage) Upload(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	ext := strings.LastIndex(fileHeader.Filename, ".")

	fileName := fmt.Sprintf("%s%s", uuid.New().String(), fileHeader.Filename[ext:])
	objectName := file_object.NewHeaderObjectName(fileName)

	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}

	fileByte := buf.Bytes()

	tagging := fmt.Sprintf("category=%d", file_object.HeaderObjectType)

	objectNameStr := objectName.GetObjectName()

	_, err := r.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:  &r.bucketName,
		Key:     &objectNameStr,
		Body:    bytes.NewReader(fileByte),
		Tagging: &tagging,
	})

	if err != nil {
		return "", err
	}

	return fileName, nil
}

func (r *HeaderImageStorage) Delete(ctx context.Context, fileName string) error {
	objectName := file_object.NewHeaderObjectName(fileName)

	objectNameStr := objectName.GetObjectName()

	_, err := r.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &r.bucketName,
		Key:    &objectNameStr,
	})

	return err
}

func (r *HeaderImageStorage) IsExists(ctx context.Context, fileName string) (bool, error) {
	objectName := file_object.NewHeaderObjectName(fileName)

	objectNameStr := objectName.GetObjectName()

	_, err := r.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &r.bucketName,
		Key:    &objectNameStr,
	})

	return err == nil, nil
}
