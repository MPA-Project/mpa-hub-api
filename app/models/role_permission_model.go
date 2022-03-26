package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RolePermissions struct {
	ID uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`

	PermissionID uuid.UUID  `gorm:"type:uuid;null;size:36;"`
	Permission   Permission `gorm:"foreignkey:PermissionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	RoleID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	Role   Role      `gorm:"foreignkey:RoleID"`

	gorm.Model
}

func (u *RolePermissions) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
