package configuration

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"myponyasia.com/hub-api/exception"
)

// Declare the variable for the database
var DB *gorm.DB

func NewDatabase() {
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	maxPoolOpen, err := strconv.Atoi(os.Getenv("DATABASE_POOL_MAX_CONN"))
	exception.PanicLogging(err)
	maxPoolIdle, err := strconv.Atoi(os.Getenv("DATABASE_POOL_IDLE_CONN"))
	exception.PanicLogging(err)
	maxPollLifeTime, err := strconv.Atoi(os.Getenv("DATABASE_POOL_LIFE_TIME"))
	exception.PanicLogging(err)

	loggerDb := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(mysql.Open(user+":"+pass+"@tcp("+host+":"+port+")/"+dbName+"?parseTime=true"), &gorm.Config{
		Logger: loggerDb,
	})
	exception.PanicLogging(err)

	sqlDB, err := db.DB()
	exception.PanicLogging(err)

	sqlDB.SetMaxOpenConns(maxPoolOpen)
	sqlDB.SetMaxIdleConns(maxPoolIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)

	//autoMigrate
	//err = db.AutoMigrate(&entity.Product{})
	//err = db.AutoMigrate(&entity.Transaction{})
	//err = db.AutoMigrate(&entity.TransactionDetail{})
	//err = db.AutoMigrate(&entity.User{})
	//err = db.AutoMigrate(&entity.UserRole{})
	//exception.PanicLogging(err)
	DB = db
}
