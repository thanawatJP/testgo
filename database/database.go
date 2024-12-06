package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=password dbname=gopro port=5432 sslmode=disable"

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // output logger
		logger.Config{
			SlowThreshold: time.Second, // ช่วงเวลาที่ถือว่าช้า
			LogLevel:      logger.Info, // ระดับการ log เช่น Info, Warn, Error
			Colorful:      true,        // ใช้สีใน Log
		},
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")
}
