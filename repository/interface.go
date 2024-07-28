package repository

import "github.com/savioafs/findAFriendAPI/model"

type PetStorer interface {
	CreatePet(pet *model.Pet) error
	FindPetByID(id string) (*model.Pet, error)
	FindPetByName(name string) (*model.Pet, error)
	FindAllPets(page, limit int, sort string) ([]model.Pet, error)
	DeletePet(pet *model.Pet) error
}

type OrganiztionStorer interface {
	CreateOrganization(organization *model.Organization) error
	FindOrganizationByID(id string) (*model.Organization, error)
	FindOrganizationByName(name string) (*model.Organization, error)
}
