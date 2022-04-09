package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnimeRelease struct {
	ID                uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`
	Title             string    `json:"name" validate:"required,lte=255" gorm:"type:varchar(512);not null;"`
	Position          int       `gorm:"type:int;null;default 1;"`
	Episode           int       `gorm:"type:int;null;default 0;"`
	Description       string    `json:"description" validate:"required" gorm:"type:text;null;"`
	DescriptionImages string    `gorm:"type:text;null;"`
	UseAds            bool      `gorm:"type:boolean;default:false;"`

	StreamingLink string `gorm:"type:text;null;"`
	DownloadLink  string `gorm:"type:text;null;"`

	AnimeID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	Anime   Anime     `gorm:"foreignkey:AnimeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	LanguageID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	Language   Language  `gorm:"foreignkey:LanguageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	UserID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	User   User      `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	gorm.Model
}

func (u *AnimeRelease) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
