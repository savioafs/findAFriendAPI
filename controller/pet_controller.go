package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/findAFriendAPI/model"
)

type PetController struct {
	// useCase
}

func NewPetController() PetController {
	return PetController{}
}

func (pc *PetController) GetAllPets(c *gin.Context) {
	pets := []model.Pet{
		{
			Name:  "ChulescO",
			About: "Animado pra caramba",
		},
	}

	c.JSON(http.StatusOK, pets)
}
