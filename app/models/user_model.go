package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"myponyasia.com/hub-api/pkg/utils/hash"
)

type User struct {
	ID               uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`
	Username         string    `json:"username" validate:"required,lte=255" gorm:"type:varchar(255);not null;"`
	Email            string    `json:"email" validate:"required,lte=255" gorm:"unique;type:varchar(255);not null;"`
	Password         string    `json:"password" validate:"required,lte=255" gorm:"type:varchar(255);not null;"`
	ExternalLink     string    `json:"externalLink" gorm:"type:text;null;"`
	DonateLink       string    `json:"donateLink" gorm:"type:text;null;"`
	Avatar           string    `json:"avatar" gorm:"type:text;null;"`
	AvatarBackground string    `json:"avatarBackground" gorm:"type:text;null;"`

	EmailVerify   bool      `gorm:"type:boolean;default:false;"`
	EmailVerifyAt time.Time `gorm:"default:null;null;"`
	TFAEnable     bool      `gorm:"type:boolean;default:false;"`
	TFAKey        string    `gorm:"type:varchar(255);null;"`
	TFABackup     string    `gorm:"type:text;null;"`

	gorm.Model
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	hash, errHash := hash.HashPassword(u.Password)
	if errHash != nil {
		return errHash
	}
	u.Password = hash

	u.EmailVerify = false
	u.TFAEnable = false

	// if u.Role == "admin" {
	// 	return errors.New("invalid role")
	// }
	return
}
