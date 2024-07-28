package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/findAFriendAPI/dto"
	"github.com/savioafs/findAFriendAPI/useCase"
)

type OrganizationController struct {
	organizationUseCase *useCase.OrganizationUseCase
}

func NewOrganizationController() *OrganizationController {
	return &OrganizationController{
		organizationUseCase: useCase.NewOrganizationUseCase(),
	}
}

func (oc *OrganizationController) CreateOrganization(c *gin.Context) {
	var org dto.OrganizationDTO
	err := c.BindJSON(&org)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": InvalidData,
		})
		return
	}

	err = oc.organizationUseCase.CreateOrganization(org)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Message": org.ID,
	})
}

func (oc *OrganizationController) FindByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": InvalidOrEmptyIDMsg,
		})
		return
	}

	org, err := oc.organizationUseCase.FindOrganizationByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, org)
}
