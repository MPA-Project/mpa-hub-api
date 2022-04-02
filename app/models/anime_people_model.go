package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnimePeople struct {
	ID uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`

	AnimeID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	Anime   Anime     `gorm:"foreignkey:AnimeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	PeopleID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	People   People    `gorm:"foreignkey:PeopleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	gorm.Model
}

func (u *AnimePeople) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
