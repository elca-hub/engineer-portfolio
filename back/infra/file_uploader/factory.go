package file_uploader

import (
	"errors"
)

var (
	errInvalidFileUploaderInstance = errors.New("invalid file uploader instance")
)

const (
	InstanceMinio int = iota
)

func NewFileUploaderFactory(instance int) (FileUploader, error) {
	switch instance {
	case InstanceMinio:
		return NewMinio(NewFileUploaderConfig())
	default:
		return nil, errInvalidFileUploaderInstance
	}
}
