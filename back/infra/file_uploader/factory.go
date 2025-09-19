package file_uploader

import (
	"devport/domain/repo/file_storage"
	"errors"
)

var (
	errInvalidFileUploaderInstance = errors.New("invalid file uploader instance")
)

const (
	InstanceMinio int = iota
)

type StorageRepositoryInter interface {
	HeaderImageStorageRepository() file_storage.HeaderImageStorage
	IconImageStorageRepository() file_storage.IconImageStorage
	BioImageStorageRepository() file_storage.BioImageStorageRepository
	BioSentenceStorageRepository() file_storage.BioSentenceStorageRepository
	WorkImageStorageRepository() file_storage.WorkImageStorageRepository
	WorkSentenceStorageRepository() file_storage.WorkSentenceStorageRepository
	WorkThumbnailStorageRepository() file_storage.WorkThumbnailStorageRepository
}

func NewStorageRepositoryFactory(instance int) (StorageRepositoryInter, error) {
	switch instance {
	case InstanceMinio:
		return NewMinioHandler(NewFileUploaderConfig())
	default:
		return nil, errInvalidFileUploaderInstance
	}
}
