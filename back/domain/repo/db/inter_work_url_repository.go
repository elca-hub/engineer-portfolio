//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE
package db

import (
	"context"
	"devport/domain/model"
)

type WorkUrlRepository interface {
	CreateByWorkId(context context.Context, workId string, workUrl []*model.WorkUrl) error
	Delete(context context.Context, id string) error
	DeleteByWorkId(context context.Context, workId string) error
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
