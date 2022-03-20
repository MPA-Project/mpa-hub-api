package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoles struct {
	ID uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`

	UserID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	User   User      `gorm:"foreignkey:UserID"`

	RoleID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	Role   Role      `gorm:"foreignkey:RoleID"`

	gorm.Model
}

func (u *UserRoles) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
