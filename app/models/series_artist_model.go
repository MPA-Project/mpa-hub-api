package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeriesArtist struct {
	ID uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`

	ArtistID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	Artist   Artist    `gorm:"foreignkey:ArtistID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	SeriesID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	Series   Series    `gorm:"foreignkey:SeriesID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	gorm.Model
}

func (u *SeriesArtist) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
