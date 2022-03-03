package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRequest struct {
	ID          uuid.UUID `gorm:"primary_key,type:uuid"`
	UserID      uuid.UUID `gorm:"type:varchar(255);not null;index"`
	RequestType string    `gorm:"type:varchar(255);not null;index"`
	Key         string    `gorm:"type:varchar(255);not null;"`
	KeyHash     string    `gorm:"type:varchar(255);not null;index"`
	ExpiredAt   time.Time `gorm:"not null;"`

	gorm.Model
}

func (u *UserRequest) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
