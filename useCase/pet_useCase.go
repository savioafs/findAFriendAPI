package useCase

import "github.com/savioafs/findAFriendAPI/model"

type PetUseCase struct {
	// repository
}

func NewPetUseCase() PetUseCase {
	return PetUseCase{}
}

func (pu *PetUseCase) GetAllPets() ([]model.Pet, error) {
	return []model.Pet{}, nil
}
