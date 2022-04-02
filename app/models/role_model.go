package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`

	Name        string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Description string `gorm:"type:varchar(255);not null;"`
	Level       int    `gorm:"type:tinyint;size(2);not null;UNSIGNED;"`

	gorm.Model
}

func (u *Role) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
