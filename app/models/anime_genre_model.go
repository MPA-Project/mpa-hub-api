package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnimeGenre struct {
	ID uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`

	AnimeID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	Anime   Anime     `gorm:"foreignkey:AnimeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	GenreID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	Genre   Genre     `gorm:"foreignkey:GenreID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	gorm.Model
}

func (u *AnimeGenre) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
