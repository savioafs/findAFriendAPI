package main

import (
	"github.com/gin-gonic/gin"
	"github.com/savioafs/findAFriendAPI/controller"
)

func main() {
	server := gin.Default()

	petController := controller.NewPetController()

	v1 := server.Group("/api/v1")
	{
		pets := v1.Group("/pets")
		{
			pets.GET("", petController.GetAllPets)
		}
	}

	server.Run(":8088")
}
