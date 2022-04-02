package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserTicket struct {
	ID uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`

	UserID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	User   User      `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	TicketData string `gorm:"type:TEXT;null;"`

	RequestType string    `gorm:"type:varchar(255);not null;index;"`
	Key         string    `gorm:"type:varchar(255);not null;"`
	KeyHash     string    `gorm:"type:varchar(32);size:32;not null;index;"`
	ExpiredAt   time.Time `gorm:"not null;"`

	gorm.Model
}

func (u *UserTicket) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
