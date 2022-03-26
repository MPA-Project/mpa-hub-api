package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeriesAuthor struct {
	ID uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`

	AuthorID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	Author   Author    `gorm:"foreignkey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	SeriesID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	Series   Series    `gorm:"foreignkey:SeriesID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	gorm.Model
}

func (u *SeriesAuthor) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
