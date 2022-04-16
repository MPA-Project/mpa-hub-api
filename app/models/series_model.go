package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Series struct {
	ID             uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`
	Title          string    `json:"title" validate:"required,lte=500" gorm:"type:varchar(512);not null;"`
	Cover           string    `gorm:"type:varchar(255);size(255);null"`
	CoverBackground string    `gorm:"type:varchar(255);size(255);null"`
	TitleAlternate string    `json:"titleAlternate" validate:"required" gorm:"type:text;null;"`

	gorm.Model
}

func (u *Series) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

func (u *Series) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}
