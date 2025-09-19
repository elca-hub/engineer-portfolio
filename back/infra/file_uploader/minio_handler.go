package file_uploader

import (
	"context"
	"devport/domain/repo/file_storage"
	"devport/infra/file_uploader/storage_repository"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type MinioHandler struct {
	uploader   *manager.Uploader
	client     *s3.Client
	bucketName string
	endpoint   string
}

func NewMinioHandler(c *FileUploaderConfig) (*MinioHandler, error) {
	creds := credentials.NewStaticCredentialsProvider(c.accessKeyID, c.secretAccessKey, "")
	sdkConfig, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion("us-east-1"),
		config.WithBaseEndpoint(c.endpoint),
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

	return &MinioHandler{
		uploader:   uploader,
		client:     s3Client,
		bucketName: c.bucketName,
		endpoint:   c.publicEndpoint,
	}, nil
}

func (m *MinioHandler) HeaderImageStorageRepository() file_storage.HeaderImageStorage {
	return storage_repository.NewHeaderImageStorage(m.uploader, m.client, m.bucketName)
}

func (m *MinioHandler) IconImageStorageRepository() file_storage.IconImageStorage {
	return storage_repository.NewIconImageStorage(m.uploader, m.client, m.bucketName)
}

func (m *MinioHandler) BioImageStorageRepository() file_storage.BioImageStorageRepository {
	return storage_repository.NewBioImageStorage(m.uploader, m.client, m.bucketName, m.endpoint)
}

func (m *MinioHandler) BioSentenceStorageRepository() file_storage.BioSentenceStorageRepository {
	return storage_repository.NewBioSentenceStorage(m.uploader, m.client, m.bucketName)
}

func (m *MinioHandler) WorkImageStorageRepository() file_storage.WorkImageStorageRepository {
	return storage_repository.NewWorkImageStorage(m.uploader, m.client, m.bucketName, m.endpoint)
}

func (m *MinioHandler) WorkSentenceStorageRepository() file_storage.WorkSentenceStorageRepository {
	return storage_repository.NewWorkSentenceStorage(m.uploader, m.client, m.bucketName)
}

func (m *MinioHandler) WorkThumbnailStorageRepository() file_storage.WorkThumbnailStorageRepository {
	return storage_repository.NewWorkThumbnailStorage(m.uploader, m.client, m.bucketName, m.endpoint)
}
