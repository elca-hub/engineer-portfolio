package gorm_model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	ID                string `gorm:"primaryKey"`
	Name              string `gorm:"size:255;not null"`
	Age               int    `gorm:"not null"`
	Email             string `gorm:"size:255;unique;not null"`
	Password          string `gorm:"size:255;not null"`
	EmailVerification int    `gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
