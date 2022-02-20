package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID            uuid.UUID `gorm:"primary_key,type:uuid"`
	Username      string    `json:"username" validate:"required,lte=255" gorm:"type:varchar(255);not null;"`
	SocialMedia   string    `json:"social_media" gorm:"type:text;null;"`
	Description   string    `json:"description" gorm:"type:text;null;"`
	AvatarProfile string    `gorm:"type:varchar(255);null;"`
	BannerProfile string    `gorm:"type:varchar(255);null;"`
	Status        string    `json:"status" gorm:"type:varchar(255);null;comment:Active | inactive;"`

	gorm.Model
}
