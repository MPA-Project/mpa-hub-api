package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Artist struct {
	ID   uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`
	Name string    `json:"name" validate:"required,lte=255" gorm:"type:varchar(255);not null;"`

	gorm.Model
}

func (u *Artist) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
