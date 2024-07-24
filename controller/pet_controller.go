package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/findAFriendAPI/model"
	"github.com/savioafs/findAFriendAPI/useCase"
)

type PetController struct {
	petUseCase useCase.PetUseCase
}

func NewPetController(petUseCase useCase.PetUseCase) PetController {
	return PetController{
		petUseCase: petUseCase,
	}
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
