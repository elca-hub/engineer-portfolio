package gorm_model

import (
	"gorm.io/gorm"
	"time"
)

type Skill struct {
	gorm.Model

	Name      string    `gorm:"size:255;not null"`
	Status    string    `gorm:"size:255;not null"`
	When      time.Time `gorm:"not null"`
	SortIndex int       `gorm:"not null"`
	UserId    uint
}
