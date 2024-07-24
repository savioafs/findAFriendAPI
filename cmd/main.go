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
	"gorm.io/gorm"
)

func main() {
	server := gin.Default()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	connection, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	connection.AutoMigrate(&model.Pet{})

	petRepository := repository.NewPetRepository(connection)
	petUseCase := useCase.NewPetUseCase(petRepository)
	petController := controller.NewPetController(petUseCase)

	v1 := server.Group("/api/v1")
	{
		pets := v1.Group("/pets")
		{
			pets.GET("")
			pets.POST("", petController.CreatePet)
			pets.GET("/:id", petController.FindByID)
		}
	}

	server.Run(":8088")
}
