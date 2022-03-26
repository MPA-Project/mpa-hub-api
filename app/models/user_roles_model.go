package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoles struct {
	ID uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`

	UserID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	User   User      `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	RoleID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	Role   Role      `gorm:"foreignkey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	gorm.Model
}

func (u *UserRoles) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
