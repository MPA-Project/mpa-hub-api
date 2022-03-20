package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	ID uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`

	Name        string `gorm:"type:varchar(255);not null;"`
	Key         string `gorm:"type:varchar(255);not null;"`
	description string `gorm:"type:varchar(255);not null;"`

	gorm.Model
}

func (u *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
