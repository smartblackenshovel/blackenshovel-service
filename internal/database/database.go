package database

import (
	"fmt"
	"log"

	"blackenshovel-service/config"
	"blackenshovel-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect opens a connection to the database
func Connect() {
	c := config.AppConfig

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection established")

	// Auto-migrate Organization table
	err = DB.AutoMigrate(&models.Organization{})
	if err != nil {
		log.Fatal("Failed to auto-migrate:", err)
	}
}
