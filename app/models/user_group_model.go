package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGroup struct {
	ID uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`

	UserID uuid.UUID `gorm:"type:uuid;not null;size:36;"`
	User   User      `gorm:"foreignkey:UserID"`

	GroupID uuid.UUID `gorm:"type:uuid;not null;size:36;"`
	Group   Group     `gorm:"foreignkey:GroupID"`

	Role string `gorm:"type:varchar(255);not null;"`

	gorm.Model
}

func (u *UserGroup) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
