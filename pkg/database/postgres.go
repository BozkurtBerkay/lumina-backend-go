package database

import (
	"fmt"
	"log"
	"os"

	"github.com/macbook/lumina-backend-go/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Successfully connected to database")

	// Auto Migration
	err = db.AutoMigrate(
		&entities.Grade{},
		&entities.Course{},
		&entities.Unit{},
		&entities.Question{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database migration completed")
	DB = db
}
