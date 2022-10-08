package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Anime struct {
	ID              uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`
	Title           string    `json:"title" validate:"required,lte=500" gorm:"type:varchar(512);not null;"`
	TitleAlternate  string    `json:"titleAlternate" validate:"required" gorm:"type:text;null;"`
	Cover           string    `gorm:"type:varchar(255);size(255);null"`
	CoverBackground string    `gorm:"type:varchar(255);size(255);null"`
	VideoBackground string    `gorm:"type:varchar(255);size(255);null"`
	Rating          string    `gorm:"type:varchar(255);size(255);null"`
	Duration        string    `gorm:"type:varchar(255);size(255);null"`
	Status          string    `gorm:"type:varchar(25);size(25);null;index;comment:Not Yet Aired | Airing | Completed | On Delayed;"`
	Type            string    `gorm:"type:varchar(25);size(25);null;index;comment:TV | ONA | OVA | Special | Movie;"`
	PV              string    `gorm:"type:text;null;"`
	Synopsis        string    `gorm:"type:text;null;"`
	ExternalLink    string    `gorm:"type:text;null;"`
	NSFW            bool      `gorm:"type:boolean;default:false;"`

	SeriesID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	Series   Series    `gorm:"foreignkey:SeriesID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	AnimeSeasonID uuid.UUID   `gorm:"type:uuid;null;size:36;"`
	AnimeSeason   AnimeSeason `gorm:"foreignkey:AnimeSeasonID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	gorm.Model
}

func (u *Anime) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
