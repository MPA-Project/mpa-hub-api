package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileManager struct {
	ID uuid.UUID `gorm:"primary_key,type:uuid;size:36;"`

	UserID uuid.UUID `gorm:"type:uuid;null;size:36;"`
	User   User      `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Filename  string `gorm:"type:varchar(100);null;size:100;"`
	Extension string `gorm:"type:varchar(25);null;size:25;"`
	MimeType  string `gorm:"type:varchar(50);null;size:50;"`
	PYear     string `gorm:"type:varchar(10);null;size:10;"`
	PMonth    string `gorm:"type:varchar(10);null;size:10;"`
	PDay      string `gorm:"type:varchar(10);null;size:10;"`

	Storage      string `gorm:"type:varchar(25);null;size:25;index;"`
	UploadStatus string `gorm:"type:varchar(25);null;size:25;index;default:QUEUE;comment:QUEUE | UPLOADING | UPLOADED | FAILED;"`
	Filesize     int64  `gorm:"type:bigint;null;default:0;"`

	ImageHeight           int      `gorm:"type:int;null;default:0;"`
	ImageWidth            int      `gorm:"type:int;null;default:0;"`
	ImageWebpSupport      bool     `gorm:"type:boolean;null;default:false;"`
	ImageThumbnailSupport bool     `gorm:"type:boolean;null;default:false;"`
	ImageAvailableRes     []string `gorm:"type:text;null;"`

	gorm.Model
}

func (u *FileManager) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type FileManagerModelImage struct {
	ID                uuid.UUID
	Filename          string
	PYear             string
	PMonth            string
	PDay              string
	ImageHeight       int
	ImageWidth        int
	ImageAvailableRes []string
}
