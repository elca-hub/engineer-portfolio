package gorm_model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID                  string               `gorm:"size:255;not null;unique;primaryKey"`
	Name                string               `gorm:"size:255;not null"`
	Email               string               `gorm:"size:255;unique;not null"`
	Birthday            time.Time            `gorm:"not null,default:CURRENT_TIMESTAMP,type:date"`
	IconPath            string               `gorm:"size:255,default''"`
	HeaderPath          string               `gorm:"size:255,default''"`
	BioPath             string               `gorm:"size:255,default''"`
	OrganizationName    string               `gorm:"size:255,default:'"`
	Skills              []Skill              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ExternalServiceUrls []ExternalServiceUrl `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt           time.Time            `gorm:"autoCreateTime"`
	UpdatedAt           time.Time            `gorm:"autoUpdateTime"`
}
