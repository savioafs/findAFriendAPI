package dto

import (
	"github.com/savioafs/findAFriendAPI/model"
)

type PetDTO struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name" binding:"required"`
	About          string `json:"about" binding:"required"`
	Age            string `json:"age" binding:"required"`
	Size           string `json:"size" binding:"required"`
	EnergyLevel    string `json:"energy_level" binding:"required"`
	DependeceLevel string `json:"dependence_level" binding:"required"`
	Ambience       string `json:"ambience" binding:"required"`
	Requirements   string `json:"requirements" binding:"required"`
	CreatedAt      string `json:"created_at"`
	OrganizationID string `json:"organization_id" binding:"required"`
}

type OrganizationDTO struct {
	ID       string   `json:"id,omitempty"`
	Name     string   `json:"name" binding:"required"`
	Manager  string   `json:"manager" binding:"required"`
	Email    string   `json:"email" binding:"required,email"`
	Password string   `json:"password" binding:"required"`
	Cep      string   `json:"cep" binding:"required"`
	Street   string   `json:"street" binding:"required"`
	District string   `json:"district" binding:"required"`
	City     string   `json:"city" binding:"required"`
	State    string   `json:"state" binding:"required"`
	Pets     []PetDTO `json:"pets"`
}

func NewPetDTO(pet model.Pet) PetDTO {
	return PetDTO{
		ID:             pet.ID.String(),
		Name:           pet.Name,
		About:          pet.About,
		Age:            pet.Age,
		Size:           pet.Size,
		EnergyLevel:    pet.EnergyLevel,
		DependeceLevel: pet.DependeceLevel,
		Ambience:       pet.Ambience,
		Requirements:   pet.Requirements,
		CreatedAt:      pet.CreatedAt.String(),
		OrganizationID: pet.OrganizationID.String(),
	}
}

func NewOrganizationDTO(org model.Organization, pets []PetDTO) OrganizationDTO {
	return OrganizationDTO{
		ID:      org.ID.String(),
		Name:    org.Name,
		Manager: org.Manager,
		Email:   org.Email,
		City:    org.Address.City,
		Pets:    pets,
	}
}
