package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/savioafs/findAFriendAPI/model"
	"gorm.io/gorm"
)

type PetRepository struct {
	connection *gorm.DB
}

func NewPetRepository(connection *gorm.DB) *PetRepository {
	return &PetRepository{
		connection: connection,
	}
}

func (pr *PetRepository) CreatePet(pet *model.Pet) error {
	pet.ID = uuid.New()
	fmt.Println(pet.ID)
	return pr.connection.Create(pet).Error
}

func (pr *PetRepository) FindByID(id string) (*model.Pet, error) {
	var pet model.Pet
	err := pr.connection.First(&pet, "id = ?", id).Error
	return &pet, err
}

func (pu *PetRepository) FindAll(page, limit int, sort string) ([]model.Pet, error) {
	var pets []model.Pet

	err := pu.connection.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&pets).Error

	return pets, err
}
