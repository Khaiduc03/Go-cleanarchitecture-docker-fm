package configuration

import (
	"FM/src/core/exception"
	"FM/src/entities"
	"log"
	"math/rand"
	"os"

	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDataBase(config Config) *gorm.DB {
	username := config.Get("POSTGRES_USER")
	password := config.Get("POSTGRES_PASSWORD")
	host := config.Get("POSTGRES_HOST")
	port := config.Get("POSTGRES_PORT")
	dbName := config.Get("POSTGRES_DB")
	TYPE := config.Get("TYPE")
	url := config.Get("POSTGRES_URL")
	maxPoolOpen, err := strconv.Atoi(config.Get("DATA_SOURCE_POOL_MAX_CONN"))
	exception.PanicLogging(err)

	maxPoolIdle, err := strconv.Atoi(config.Get("DATA_SOURCE_POOL_IDLE_CONN"))
	exception.PanicLogging(err)

	maxPollLifeTime, err := strconv.Atoi(config.Get("DATA_SOURCE_POOL_LIFE_TIME"))

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

	debug := "host=" + host + " user=" + username + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	//debug := "host=postgres" + " user=postgres " + " password=postgres" + " dbname=fm" + " port=5432" + " sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	var dsn string 
	if TYPE == "prod" {
		dsn = url
	} else {
		dsn = debug
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: loggerDb,
	})
	exception.PanicLogging(err)

	sqlDB, err := db.DB()
	exception.PanicLogging(err)

	sqlDB.SetMaxOpenConns(maxPoolOpen)
	sqlDB.SetMaxIdleConns(maxPoolIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)

	// auto migrate
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Category{})
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Image{})
	db.AutoMigrate(&entities.Rating{})
	db.AutoMigrate(&entities.Schedule{})
	db.AutoMigrate(&entities.FeedBack{})
	db.AutoMigrate(&entities.Room{})
	return db
}
