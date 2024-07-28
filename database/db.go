package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/savioafs/findAFriendAPI/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartConnectionDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	fmt.Println("DATABASE_URL:", databaseURL)

	connection, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	connection.AutoMigrate(&model.Organization{}, &model.Pet{})

	return connection
}
