package repository

import "github.com/savioafs/findAFriendAPI/model"

type PetStorer interface {
	CreatePet(pet *model.Pet) error
	FindByID(id string) (*model.Pet, error)
	FindAll(page, limit int, sort string) ([]model.Pet, error)
	Delete(pet *model.Pet) error
}

type OrganiztionStorer interface {
	CreateOrganization(organization *model.Organization) error
}
