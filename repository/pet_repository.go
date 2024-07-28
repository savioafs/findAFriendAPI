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

func (pr *PetRepository) FindPetByID(id string) (*model.Pet, error) {
	var pet model.Pet
	err := pr.connection.First(&pet, "id = ?", id).Error
	return &pet, err
}

func (pr *PetRepository) FindPetByName(name string) (*model.Pet, error) {
	var pet model.Pet

	err := pr.connection.Where("name = ?", name).First(&pet).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &pet, err
}

func (pr *PetRepository) FindAllPets(page, limit int, sort string) ([]model.Pet, error) {
	var pets []model.Pet

	err := pr.connection.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&pets).Error

	return pets, err
}

func (pr *PetRepository) DeletePet(pet *model.Pet) error {
	err := pr.connection.Delete(pet).Error
	return err
}
