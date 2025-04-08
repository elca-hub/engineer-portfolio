package file_uploader

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Minio struct {
	uploader   *manager.Uploader
	client     *s3.Client
	bucketName string
}

func NewMinio(fileUploaderConfig *FileUploaderConfig) (*Minio, error) {
	creds := credentials.NewStaticCredentialsProvider(fileUploaderConfig.accessKeyID, fileUploaderConfig.secretAccessKey, "")
	sdkConfig, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion("us-east-1"),
		config.WithBaseEndpoint(fileUploaderConfig.endpoint),
		config.WithCredentialsProvider(creds),
	)

	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(sdkConfig, func(options *s3.Options) {
		options.UsePathStyle = true
	})

	uploader := manager.NewUploader(s3Client, func(options *manager.Uploader) {
		options.PartSize = 5 * 1024 * 1024 // 5 MB
		options.Concurrency = 5
	})

	return &Minio{
		uploader:   uploader,
		client:     s3Client,
		bucketName: fileUploaderConfig.bucketName,
	}, nil
}

func (m *Minio) UploadFile(fileBytes []byte, objectName string) error {
	_, err := m.uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &m.bucketName,
		Key:    &objectName,
		Body:   bytes.NewReader(fileBytes),
	})

	return err
}

func (m *Minio) DeleteFile(objectName string) error {
	_, err := m.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &m.bucketName,
		Key:    &objectName,
	})

	if err != nil {
		return err
	}

	return nil
}
