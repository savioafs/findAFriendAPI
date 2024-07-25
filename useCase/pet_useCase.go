package useCase

import (
	"github.com/savioafs/findAFriendAPI/model"
	"github.com/savioafs/findAFriendAPI/repository"
)

type PetUseCase struct {
	petRepository *repository.PetRepository
}

func NewPetUseCase(petRepository *repository.PetRepository) *PetUseCase {
	return &PetUseCase{
		petRepository: petRepository,
	}
}

func (pu *PetUseCase) CreatePet(pet *model.Pet) error {
	err := pu.petRepository.CreatePet(pet)
	if err != nil {
		return err
	}

	return nil
}

func (pu *PetUseCase) FindByID(id string) (*model.Pet, error) {
	pet, err := pu.petRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return pet, nil
}

func (pu *PetUseCase) FindAll(page, limit int, sort string) ([]model.Pet, error) {

	pets, err := pu.petRepository.FindAll(page, limit, sort)
	if err != nil {
		return nil, err
	}

	return pets, nil
}
