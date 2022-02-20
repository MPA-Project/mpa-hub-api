package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"myponyasia.com/hub-api/pkg/utils/hash"
)

// User struct to describe object.
type User struct {
	ID          uuid.UUID `gorm:"primary_key,type:uuid"`
	Username    string    `json:"username" validate:"required,lte=255" gorm:"type:varchar(255);not null;"`
	Email       string    `json:"email" validate:"required,lte=255" gorm:"unique;type:varchar(255);not null;"`
	Password    string    `json:"password" validate:"required,lte=255" gorm:"type:varchar(255);not null;"`
	Role        string    `gorm:"type:varchar(255);not null;"`
	SocialMedia string    `json:"social_media" gorm:"type:text;null;"`

	gorm.Model
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	hash, errHash := hash.HashPassword(u.Password)
	if errHash != nil {
		return errHash
	}
	u.Password = hash

	// if u.Role == "admin" {
	// 	return errors.New("invalid role")
	// }
	return
}
