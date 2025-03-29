package gorm_model

import (
	"gorm.io/gorm"
)

type ExternalServiceUrl struct {
	gorm.Model

	Name string `gorm:"size:100;not null"`
	Url  string `gorm:"size:255;not null"`

	UserId uint
}
