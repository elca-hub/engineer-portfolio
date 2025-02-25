package gorm_model

import (
	"gorm.io/gorm"
)

type Connect struct {
	gorm.Model
	id         uint64 `gorm:"primaryKey;autoIncrement"`
	parentItem Item   `gorm:"foreignKey:ParentItemID"`
	childItem  Item   `gorm:"foreignKey:ChildItemID"`
}
