package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/findAFriendAPI/model"
	"github.com/savioafs/findAFriendAPI/useCase"
)

type PetController struct {
	petUseCase *useCase.PetUseCase
}

func NewPetController(petUseCase *useCase.PetUseCase) *PetController {
	return &PetController{
		petUseCase: petUseCase,
	}
}

func (pc *PetController) CreatePet(c *gin.Context) {
	var pet model.Pet
	err := c.BindJSON(&pet)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = pc.petUseCase.CreatePet(&pet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Message: ": "Pet registred with succeess",
	})
}

func (pc *PetController) FindByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message: ": "invalid or empty id",
		})
		return
	}

	petOk, err := pc.petUseCase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, petOk)
}
