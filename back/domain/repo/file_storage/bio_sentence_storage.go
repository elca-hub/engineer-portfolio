//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE
package file_storage

import (
	"context"
)

type BioSentenceStorageRepository interface {
	Upload(context context.Context, userId string, sentence string) (string, error)
	Delete(context context.Context, userId string) error
	IsExists(context context.Context, userId string) (bool, error)
	FindByUserId(context context.Context, userId string) (string, error)
}
