package repository

import "github.com/savioafs/findAFriendAPI/model"

type PetInterface interface {
	CreatePet(pet *model.Pet) error
	FindByID(id string) (*model.Pet, error)
	FindAll(page, limit int, sort string) ([]model.Pet, error)
}
