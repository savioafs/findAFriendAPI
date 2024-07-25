package useCase

import (
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

func (pu *PetUseCase) CreatePet(pet *model.Pet) error {
	err := pu.petInterface.CreatePet(pet)
	if err != nil {
		return err
	}

	return nil
}

func (pu *PetUseCase) FindByID(id string) (*model.Pet, error) {
	pet, err := pu.petInterface.FindByID(id)
	if err != nil {
		return nil, err
	}

	return pet, nil
}

func (pu *PetUseCase) FindAll(page, limit int, sort string) ([]model.Pet, error) {

	pets, err := pu.petInterface.FindAll(page, limit, sort)
	if err != nil {
		return nil, err
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
