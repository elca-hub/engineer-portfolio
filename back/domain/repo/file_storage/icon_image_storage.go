//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE
package file_storage

import (
	"context"
	"mime/multipart"
)

type IconImageStorage interface {
	Upload(context context.Context, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
	Delete(context context.Context, imageName string) error
	IsExists(context context.Context, imageName string) (bool, error)
}
