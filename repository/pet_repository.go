package repository

import (
	"github.com/google/uuid"
	"github.com/savioafs/findAFriendAPI/database"
	"github.com/savioafs/findAFriendAPI/model"
	"gorm.io/gorm"
)

type PetRepository struct {
	connection *gorm.DB
}

func NewPetRepository() *PetRepository {
	return &PetRepository{
		connection: database.StartConnectionDB(),
	}
}

func (pr *PetRepository) CreatePet(pet *model.Pet) error {
	pet.ID = uuid.New()

	return pr.connection.Create(pet).Error
}

func (pr *PetRepository) FindByID(id string) (*model.Pet, error) {
	var pet model.Pet
	err := pr.connection.First(&pet, "id = ?", id).Error
	return &pet, err
}

func (pr *PetRepository) FindAll(page, limit int, sort string) ([]model.Pet, error) {
	var pets []model.Pet

	err := pr.connection.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&pets).Error

	return pets, err
}

func (pr *PetRepository) Delete(pet *model.Pet) error {
	err := pr.connection.Delete(pet).Error
	return err
}
