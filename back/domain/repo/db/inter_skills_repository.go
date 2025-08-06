package db

import (
	"context"
	"devport/domain/model"
)

type SkillsRepository interface {
	Create(ctx context.Context, skill *model.Skill, userID string) error
	FindByUserID(ctx context.Context, userID string) ([]*model.Skill, error)
	Update(ctx context.Context, skill *model.Skill, userID string) error
	Delete(ctx context.Context, id string) error
	GetMaxSortIndex(ctx context.Context, userID string) (int, error)
	FindByID(ctx context.Context, userId string, skillId string) (*model.Skill, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}
