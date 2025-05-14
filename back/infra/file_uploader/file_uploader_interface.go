//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package file_uploader

import "devport/domain/model"

type FileUploader interface {
	UploadFile(fileBytes []byte, objectName string) error
	DeleteFile(objectName string) error
	GetIconNames() ([]*model.FileIconName, error)
	GetHeaderNames() ([]*model.FileIconName, error)
	GetFiles(filePattern int) ([]*model.FileIconName, error)
}
