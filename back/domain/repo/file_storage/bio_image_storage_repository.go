//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE
package file_storage

import (
	"context"
	"mime/multipart"
)

type BioImageStorageRepository interface {
	Upload(context context.Context, userId string, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
	Delete(context context.Context, userId string, imageName string) error
	IsExists(context context.Context, userId string, imageName string) (bool, error)
	FindByUserId(context context.Context, userId string) ([]string, error)
	GetPublicDomain(userId string) string // bioに含まれているimageの検索に必要
}
