//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE
package db

import (
	"context"
	"devport/domain/model"
)

type ExternalServiceUrlsRepository interface {
	Create(context context.Context, userId string, e *model.ExternalServiceUrl) error
	Update(context context.Context, userId string, e *model.ExternalServiceUrl) error
	Delete(context context.Context, userId string, e *model.ExternalServiceUrl) error
	FindByUserId(context context.Context, userId string) ([]*model.ExternalServiceUrl, error)
	FindById(context context.Context, id string) (*model.ExternalServiceUrl, error)
	FindByServiceType(context context.Context, userId string, serviceType int) (*model.ExternalServiceUrl, error)
	IsExistsByServiceType(context context.Context, userId string, serviceType int) (bool, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
