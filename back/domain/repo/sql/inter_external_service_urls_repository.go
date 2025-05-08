//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE
package sql

import (
	"context"
	"devport/domain/model"
)

type ExternalServiceUrlsRepository interface {
	Create(context context.Context, u *model.User, e *model.ExternalServiceUrl) error
	Update(context context.Context, u *model.User, e *model.ExternalServiceUrl) error
	Delete(context context.Context, u *model.User, e *model.ExternalServiceUrl) error
	FindByUserId(context context.Context, u *model.User) ([]*model.ExternalServiceUrl, error)
	FindById(context context.Context, id string) (*model.ExternalServiceUrl, error)
	FindByServiceType(context context.Context, u *model.User, serviceType int) (*model.ExternalServiceUrl, error)
	IsExistsByServiceType(context context.Context, u *model.User, serviceType int) (bool, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
