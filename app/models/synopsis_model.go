package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Synopsis struct {
	ID uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`

	LanguageID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	Language   Language  `gorm:"foreignkey:LanguageID;"`

	Synop string `json:"synopsis" validate:"required" gorm:"type:text;null;"`

	gorm.Model
}

func (u *Synopsis) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
