package useCase

import (
	"errors"

	"github.com/savioafs/findAFriendAPI/dto"
	"github.com/savioafs/findAFriendAPI/model"
	"github.com/savioafs/findAFriendAPI/repository"
)

type PetUseCase struct {
	petInterface repository.PetStorer
}

func NewPetUseCase() *PetUseCase {
	return &PetUseCase{
		petInterface: repository.NewPetRepository(),
	}
}

func (pu *PetUseCase) CreatePet(petRequest dto.PetDTO) error {

	if petRequest.Name == "" {
		return errors.New("pet name is required")
	}

	if len(petRequest.Name) > 100 {
		return errors.New("pet name must be less than 100 characteres")
	}

	petExists := true

	if petExists {
		return errors.New("pet with this name already exists")
	}

	pet := model.Pet{
		Name: petRequest.Name,
	}

	err := pu.petInterface.CreatePet(&pet)
	if err != nil {
		return err
	}

	return nil
}

func (pu *PetUseCase) FindByID(id string) (dto.PetDTO, error) {
	petDB, err := pu.petInterface.FindByID(id)
	if err != nil {
		return dto.PetDTO{}, err
	}

	pet := dto.PetDTO{
		Name:  petDB.Name,
		About: petDB.About,
		Age:   petDB.Age,
	}

	return pet, nil
}

func (pu *PetUseCase) FindAll(page, limit int, sort string) ([]dto.PetDTO, error) {

	petsDB, err := pu.petInterface.FindAll(page, limit, sort)
	if err != nil {
		return nil, err
	}

	var pets []dto.PetDTO

	for _, pet := range petsDB {
		petOk := dto.PetDTO{
			ID:    pet.ID.String(),
			Name:  pet.Name,
			About: pet.About,
			Age:   pet.Age,
		}

		pets = append(pets, petOk)
	}

	return pets, nil
}

func (pu *PetUseCase) Delete(id string) error {
	pet, err := pu.petInterface.FindByID(id)
	if err != nil {
		return err
	}

	err = pu.petInterface.Delete(pet)

	return err
}
