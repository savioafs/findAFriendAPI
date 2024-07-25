package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/savioafs/findAFriendAPI/controller"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	petController := controller.NewPetController()

	main := router.Group("/api/v1")
	{
		pets := main.Group("pets")
		{
			pets.POST("/", petController.CreatePet)
			pets.GET("/:id", petController.FindByID)
			pets.GET("/", petController.FindAll)
			pets.DELETE("/:id", petController.Delete)
			// add pets put
		}
	}

	return router
}
