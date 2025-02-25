package repository

import (
	"gorm.io/gorm"
	"devport/domain/model"
)

type GormItemRepository struct {
	db *gorm.DB
}

func NewGormItemRepository(db *gorm.DB) *GormItemRepository {
	return &GormItemRepository{
		db: db,
	}
}

func (r GormItemRepository) Create(item model.Item) error {
	return r.db.Create(&item).Error
}

func (r GormItemRepository) Update(item model.Item) model.Item {
	r.db.Save(&item)
	return item
}

func (r GormItemRepository) Delete(item model.Item) error {
	return r.db.Delete(&item).Error
}
