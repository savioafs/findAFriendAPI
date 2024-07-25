package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/findAFriendAPI/controller"
	"github.com/savioafs/findAFriendAPI/model"
	"github.com/savioafs/findAFriendAPI/repository"
	"github.com/savioafs/findAFriendAPI/useCase"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	server := gin.Default()

	var (
		port       string
		connection *gorm.DB
		err        error
	)

	if os.Getenv("STAGE") == "localRun" {
		databaseURL := os.Getenv("DATABASE_URL")
		if databaseURL == "" {
			log.Fatal("DATABASE_URL environment variable is not set")
		}

		connection, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to open db: %v", err)
		}

		port = ":8088"
	} else {
		connection, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to open db: %v", err)
		}

		port = ":8000"
	}

	connection.AutoMigrate(&model.Pet{})

	petRepository := repository.NewPetRepository(connection)
	petUseCase := useCase.NewPetUseCase(petRepository)
	petController := controller.NewPetController(petUseCase)

	v1 := server.Group("/api/v1")
	{
		pets := v1.Group("/pets")
		{
			pets.POST("", petController.CreatePet)
			pets.GET("/:id", petController.FindByID)
			pets.GET("/", petController.FindAll)
			pets.DELETE("/:id", petController.Delete)
		}
	}

	server.Run(port)
}
