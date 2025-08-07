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

type WorkSentenceStorage struct {
	uploader   *manager.Uploader
	client     *s3.Client
	bucketName string
}

func NewWorkSentenceStorage(uploader *manager.Uploader, client *s3.Client, bucketName string) *WorkSentenceStorage {
	return &WorkSentenceStorage{
		uploader,
		client,
		bucketName,
	}
}

func (r *WorkSentenceStorage) Upload(ctx context.Context, userId string, workId string, sentence string) (string, error) {
	objectName := file_object.NewWorkSentenceFileName(userId, workId)

	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, strings.NewReader(sentence)); err != nil {
		return "", err
	}

	fileByte := buf.Bytes()

	tagging := fmt.Sprintf("category=%d", file_object.WorkSentenceObjectType)

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

func (r *WorkSentenceStorage) Delete(ctx context.Context, userId string, workId string) error {
	objectName := file_object.NewWorkSentenceFileName(userId, workId)

	objectNameStr := objectName.GetObjectName()

	_, err := r.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &r.bucketName,
		Key:    &objectNameStr,
	})

	return err
}

func (r *WorkSentenceStorage) IsExists(ctx context.Context, userId string, workId string) (bool, error) {
	objectName := file_object.NewWorkSentenceFileName(userId, workId)

	objectNameStr := objectName.GetObjectName()

	_, err := r.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &r.bucketName,
		Key:    &objectNameStr,
	})

	return err == nil, nil
}

func (r *WorkSentenceStorage) FindByUserId(ctx context.Context, userId string) ([]string, error) {
	// TODO: 実装が必要な場合は後で追加
	return []string{}, nil
}

func (r *WorkSentenceStorage) Find(ctx context.Context, userId string, workId string) (string, error) {
	objectName := file_object.NewWorkSentenceFileName(userId, workId)

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