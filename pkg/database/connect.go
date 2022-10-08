package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"myponyasia.com/hub-api/pkg/entities"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error

	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbHost := os.Getenv("DATABASE_HOST")
	dbName := os.Getenv("DATABASE_NAME")

	// Connection URL to connect to MYSql Database
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	DB.AutoMigrate(

		// Base User
		&entities.User{},
		&entities.UserRoles{},
		&entities.Role{},
		&entities.UserTicket{},
		&entities.Group{},
		&entities.UserGroup{},
		&entities.Permission{},
		&entities.RolePermissions{},

		// Base Data
		&entities.Language{},
		&entities.Genre{},
		&entities.People{},
		&entities.FileManager{},

		// Series
		&entities.Series{},
		&entities.SeriesPeople{},
		&entities.AnimeSeason{},
		&entities.Anime{},
		&entities.AnimeGenre{},
		&entities.AnimeStudio{},
		&entities.AnimeStudioRelation{},
		&entities.AnimePeople{},
		&entities.AnimeRelease{},

		// &models.Author{},
		// &models.Artist{},
		// &models.SeriesArtist{},
		// &models.SeriesAuthor{},
	)
	fmt.Println("Database Migrated")
}
