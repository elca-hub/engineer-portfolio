package gorm_model

import (
	"gorm.io/gorm"
)

type ExternalServiceUrl struct {
	gorm.Model

	Url    string `gorm:"size:255;not null"`
	UserId string
}
