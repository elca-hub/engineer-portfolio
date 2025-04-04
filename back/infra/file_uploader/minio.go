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
	uploader *manager.Uploader
	client   *s3.Client
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
		uploader: uploader,
		client:   s3Client,
	}, nil
}

func (m *Minio) UploadFile(fileBytes []byte, bucketName string, objectName string) (string, error) {
	res, err := m.uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectName,
		Body:   bytes.NewReader(fileBytes),
	})

	if err != nil {
		return "", err
	}

	return res.Location, nil
}

func (m *Minio) CrateBucket(bucketName string) error {
	_, err := m.client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: &bucketName,
	})

	if err != nil {
		if err.Error() != "BucketAlreadyOwnedByYou: Your previous request to create the named bucket succeeded and you already own it." {
			return err
		}

		return nil
	}

	return nil
}
