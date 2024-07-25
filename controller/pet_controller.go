package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/findAFriendAPI/dto"
	"github.com/savioafs/findAFriendAPI/useCase"
)

const (
	InvalidOrEmptyIDMsg     = "invalid or empty id"
	PetRegisteredSuccessMsg = "Pet registered with success"
	DeletedWithSuccessMsg   = "deleted with success"
)

type PetController struct {
	petUseCase *useCase.PetUseCase
}

func NewPetController() *PetController {
	return &PetController{
		petUseCase: useCase.NewPetUseCase(),
	}
}

func (pc *PetController) CreatePet(c *gin.Context) {
	var pet dto.PetDTO
	err := c.BindJSON(&pet)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = pc.petUseCase.CreatePet(pet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Message": PetRegisteredSuccessMsg,
	})
}

func (pc *PetController) FindByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": InvalidOrEmptyIDMsg,
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

func (pc *PetController) Delete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": InvalidOrEmptyIDMsg,
		})
		return
	}

	err := pc.petUseCase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": DeletedWithSuccessMsg,
	})
}
