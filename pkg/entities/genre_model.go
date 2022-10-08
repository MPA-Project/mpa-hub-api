package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Genre struct {
	ID   uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`
	Name string    `json:"name" validate:"required,lte=255" gorm:"type:varchar(512);not null;"`

	gorm.Model
}

func (u *Genre) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
