package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnimeStudioRelation struct {
	ID uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`

	AnimeID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	Anime   Anime     `gorm:"foreignkey:AnimeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	AnimeStudioID uuid.UUID   `gorm:"type:uuid;null;size:36;"`
	AnimeStudio   AnimeStudio `gorm:"foreignkey:AnimeStudioID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	gorm.Model
}

func (u *AnimeStudioRelation) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
