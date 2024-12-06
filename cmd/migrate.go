package main

import (
	"awesomeProject/database"
	"awesomeProject/database/models"
	"log"
)

// RunMigrations ทำการ migrate ตาราง
func main() {
	database.Connect()

	// ทำการ AutoMigrate
	err := database.DB.AutoMigrate(&models.User{}, &models.Blog{})
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	log.Println("Migration completed successfully")
}
