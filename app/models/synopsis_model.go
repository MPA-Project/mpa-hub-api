package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Synopsis struct {
	ID           uuid.UUID `gorm:"primary_key,type:uuid"`
	LanguageCode string    `json:"code" validate:"required,lte=255" gorm:"type:varchar(255);not null;"`
	Synopsis     string    `json:"synopsis" validate:"required" gorm:"type:text;not null;"`

	gorm.Model
}
