package db

import (
	"context"
	"devport/domain/model"
)

type CertificationsRepository interface {
	Create(ctx context.Context, certification *model.Certification, userID string) error
	FindByUserID(ctx context.Context, userID string) ([]*model.Certification, error)
	Update(ctx context.Context, certification *model.Certification, userID string) error
	Delete(ctx context.Context, id string) error
	GetMaxSortIndex(ctx context.Context, userID string) (int, error)
}