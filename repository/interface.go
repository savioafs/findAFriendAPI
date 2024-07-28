package repository

import "github.com/savioafs/findAFriendAPI/model"

type PetStorer interface {
	Create(pet *model.Pet) error
	FindByID(id string) (*model.Pet, error)
	FindAll(page, limit int, sort string) ([]model.Pet, error)
	FindByName(name string) (*model.Pet, error)
	Delete(pet *model.Pet) error
}

type OrganiztionStorer interface {
	Create(organization *model.Organization) error
	FindByName(name string) (*model.Organization, error)
	FindByID(id string) (*model.Organization, error)
}
