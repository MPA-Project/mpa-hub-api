package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type People struct {
	ID     uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`
	Name   string    `json:"name" validate:"required,lte=255" gorm:"type:varchar(512);not null;"`
	About  string    `gorm:"about:text;null;"`
	Avatar string    `gorm:"type:varchar(255);size(255);null"`

	gorm.Model
}

func (u *People) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
