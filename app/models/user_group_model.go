package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGroup struct {
	ID      uuid.UUID `gorm:"primary_key,type:uuid"`
	UserID  uuid.UUID `gorm:"type:varchar(255);not null;"`
	GroupID uuid.UUID `gorm:"type:varchar(255);not null;"`
	Role    string    `gorm:"type:varchar(255);not null;"`

	gorm.Model
}

func (u *UserGroup) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
