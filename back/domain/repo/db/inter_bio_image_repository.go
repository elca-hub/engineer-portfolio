//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE
package db

import (
	"context"
	"devport/domain/model"
)

type WorkImagesRepository interface {
	Create(context context.Context, workId string, fileName string) error
	Delete(context context.Context, work *model.Work, fileName string) error
	FindByWorkId(context context.Context, workId string) ([]string, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
