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

type WorkImageStorage struct {
	uploader   *manager.Uploader
	client     *s3.Client
	bucketName string
	endpoint   string
}

func NewWorkImageStorage(uploader *manager.Uploader, client *s3.Client, bucketName string, endpoint string) *WorkImageStorage {
	return &WorkImageStorage{
		uploader,
		client,
		bucketName,
		endpoint,
	}
}

func (r *WorkImageStorage) Upload(ctx context.Context, userId string, workId string, file multipart.File, fileHeader *multipart.FileHeader) (string, string, error) {
	ext := strings.LastIndex(fileHeader.Filename, ".")

	fileName := fmt.Sprintf("%s%s", uuid.New().String(), fileHeader.Filename[ext:])
	objectName := file_object.NewWorkImageFileName(userId, workId, fileName)

	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, file); err != nil {
		return "", "", err
	}

	fileByte := buf.Bytes()

	tagging := fmt.Sprintf("category=%d", file_object.WorkImageObjectType)

	objectNameStr := objectName.GetObjectName()

	_, err := r.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:  &r.bucketName,
		Key:     &objectNameStr,
		Body:    bytes.NewReader(fileByte),
		Tagging: &tagging,
	})

	if err != nil {
		return "", "", err
	}

	return fmt.Sprintf("%s/%s", r.endpoint, objectNameStr), fileName, nil
}

func (r *WorkImageStorage) Delete(ctx context.Context, userId string, workId string, imageName string) error {
	objectName := file_object.NewWorkImageFileName(userId, workId, imageName)

	objectNameStr := objectName.GetObjectName()

	_, err := r.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &r.bucketName,
		Key:    &objectNameStr,
	})

	return err
}

func (r *WorkImageStorage) FindByWorkId(ctx context.Context, userId string, workId string) ([]string, error) {
	objectNameStr := file_object.WorkImageWorkPrefix(userId, workId)

	objects, err := r.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: &r.bucketName,
		Prefix: &objectNameStr,
	})

	if err != nil {
		return nil, err
	}

	imageNames := []string{}

	for _, object := range objects.Contents {
		imageNames = append(imageNames, strings.TrimPrefix(*object.Key, objectNameStr))
	}

	return imageNames, nil
}

func (r *WorkImageStorage) IsExists(ctx context.Context, userId string, workId string, imageName string) (bool, error) {
	objectName := file_object.NewWorkImageFileName(userId, workId, imageName)

	objectNameStr := objectName.GetObjectName()

	_, err := r.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &r.bucketName,
		Key:    &objectNameStr,
	})

	return err == nil, nil
}

func (r *WorkImageStorage) GetPublicDomain(userId string, workId string) string {
	return fmt.Sprintf("%s/%s/%s/%s", r.endpoint, file_object.WorkImageUri, userId, workId)
}