//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE
package file_storage

import (
	"context"
)

type WorkSentenceStorageRepository interface {
	Upload(context context.Context, userId string, workId string, sentence string) (string, error)
	Delete(context context.Context, userId string, workId string) error
	IsExists(context context.Context, userId string, workId string) (bool, error)
	FindByUserId(context context.Context, userId string) ([]string, error)
	Find(context context.Context, userId string, workId string) (string, error)
}
