package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var instance *gorm.DB

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold: time.Microsecond, // Slow SQL threshold
		LogLevel:      logger.Info,      // Log level
		Colorful:      false,            // Disable color
	},
)

func IniDB() *gorm.DB {

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	DB_TYPE := os.Getenv("DB_TYPE")

	var dsn gorm.Dialector

	switch DB_TYPE {
	case "mysql":
		dsn = mysql.Open(DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8&parseTime=True&loc=Local")
	case "postgres":
		dsn = postgres.Open("host=" + DB_HOST + " user=" + DB_USER + " password=" + DB_PASS + " dbname=" + DB_NAME + " port=" + DB_PORT + " sslmode=disable TimeZone=Asia/Jakarta")
	}

	db, err := gorm.Open(dsn, &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		fmt.Println("Cannot connect to database:", err)
		panic(err.Error())
	}

	// db.AutoMigrate(&entity.User{})

	return db
}
