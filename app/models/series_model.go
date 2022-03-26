package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"myponyasia.com/hub-api/pkg/enums"
	"myponyasia.com/hub-api/pkg/utils/hash"
)

type Series struct {
	ID        uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`
	Title     string    `json:"title" validate:"required,lte=255" gorm:"type:varchar(512);not null;"`
	TitleHash string    `gorm:"type:varchar(32);size:32;not null;index"`
	Type      string    `gorm:"type:varchar(255);not null;comment:Watchable | Readable;"`
	Status    int       `gorm:"index;type:tinyint;size(1);UNSIGNED;null;"`

	UserID uuid.UUID `gorm:"index;type:uuid;null;size:36;"`
	User   User      `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	gorm.Model
}

func (u *Series) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.Status = enums.StatusEnum("INACTIVE")
	u.TitleHash = hash.GetMD5Hash(u.Title)
	return
}

func (u *Series) BeforeUpdate(tx *gorm.DB) (err error) {
	u.TitleHash = hash.GetMD5Hash(u.Title)
	return
}
