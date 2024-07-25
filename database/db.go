package database

import (
	"log"
	"os"

	"github.com/savioafs/findAFriendAPI/model"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func StartConnectionDB() *gorm.DB {
	var (
		connection *gorm.DB
		err        error
	)

	if os.Getenv("STAGE") == "goDocker" {
		databaseURL := os.Getenv("DATABASE_URL")
		if databaseURL == "" {
			log.Fatal("DATABASE_URL environment variable is not set")
		}

		connection, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to open db: %v", err)
		}

	} else {
		connection, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to open db: %v", err)
		}
	}

	connection.AutoMigrate(&model.Pet{})

	return connection
}
