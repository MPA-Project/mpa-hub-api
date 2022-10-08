package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnimeSeason struct {
	ID       uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`
	Name     string    `json:"name" validate:"required,lte=255" gorm:"type:varchar(512);not null;"`
	Position int       `gorm:"type:int;null;default 0"`

	gorm.Model
}

func (u *AnimeSeason) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
