package gorm_model

import (
	"gorm.io/gorm"
	"time"
)

type Item struct {
	gorm.Model
	id        uint64 `gorm:"primaryKey;autoIncrement"`
	title     string `gorm:"size:255"`
	memo      string `gorm:"size:1024"`
	createdAt time.Time
	updatedAt time.Time
}
