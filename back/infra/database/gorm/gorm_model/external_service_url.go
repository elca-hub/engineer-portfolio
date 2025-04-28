package gorm_model

import (
	"gorm.io/gorm"
)

type ExternalServiceUrl struct {
	gorm.Model

	ID          string `gorm:"size:255;not null;unique;primaryKey"`
	Url         string `gorm:"size:255;not null"`
	UserId      string
	ServiceType int
}
