package controller

import (
	"net/http"
	"strconv"

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

func (pc *PetController) FindAll(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	sort := c.DefaultQuery("sort", "asc")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	pets, err := pc.petUseCase.FindAll(page, limit, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, pets)
}
