//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE
package file_storage

import (
	"context"
	"mime/multipart"
)

type WorkThumbnailStorageRepository interface {
	Upload(context context.Context, userId string, workId string, file multipart.File, fileHeader *multipart.FileHeader) (string, string, error)
	Delete(context context.Context, userId string, workId string, imageName string) error
	FindByWorkId(context context.Context, userId string, workId string) ([]string, error)
}
