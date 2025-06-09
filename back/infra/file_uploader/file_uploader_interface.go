//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package file_uploader

import (
	"devport/infra/file_uploader/file_object"
)

type FileUploader interface {
	UploadFile(objectName file_object.ObjectNameInterface) error
	DeleteFile(fileName file_object.FileNameInterface) error
	GetFiles(filePattern int) ([][]byte, error)
	GetFile(fileName file_object.FileNameInterface) ([]byte, error)
}
