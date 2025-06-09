package storage_repository

import (
	"bytes"
	"context"
	"devport/infra/file_uploader/file_object"
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type BioSentenceStorage struct {
	uploader   *manager.Uploader
	client     *s3.Client
	bucketName string
}

func NewBioSentenceStorage(uploader *manager.Uploader, client *s3.Client, bucketName string) *BioSentenceStorage {
	return &BioSentenceStorage{
		uploader,
		client,
		bucketName,
	}
}

func (r *BioSentenceStorage) Upload(ctx context.Context, userId string, sentence string) (string, error) {
	objectName := file_object.NewBioFileName(userId)

	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, strings.NewReader(sentence)); err != nil {
		return "", err
	}

	fileByte := buf.Bytes()

	tagging := fmt.Sprintf("category=%d", file_object.BioSentenceObjectType)

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

	return objectNameStr, nil
}

func (r *BioSentenceStorage) Delete(ctx context.Context, userId string) error {
	objectName := file_object.NewBioFileName(userId)

	objectNameStr := objectName.GetObjectName()

	_, err := r.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &r.bucketName,
		Key:    &objectNameStr,
	})

	return err
}

func (r *BioSentenceStorage) IsExists(ctx context.Context, userId string) (bool, error) {
	objectName := file_object.NewHeaderObjectName(userId)

	objectNameStr := objectName.GetObjectName()

	_, err := r.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &r.bucketName,
		Key:    &objectNameStr,
	})

	return err == nil, nil
}

func (r *BioSentenceStorage) FindByUserId(ctx context.Context, userId string) (string, error) {
	objectName := file_object.NewBioFileName(userId)

	objectNameStr := objectName.GetObjectName()

	object, err := r.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &r.bucketName,
		Key:    &objectNameStr,
	})

	if err != nil {
		// 404の時は空文字を返す
		if strings.Contains(err.Error(), "NoSuchKey") {
			return "", nil
		}

		return "", err
	}

	body, err := io.ReadAll(object.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
