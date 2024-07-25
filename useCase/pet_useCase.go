package useCase

import (
	"github.com/savioafs/findAFriendAPI/model"
	"github.com/savioafs/findAFriendAPI/repository"
)

type PetUseCase struct {
	petInterface repository.PetStorer
}

type PetDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	About string `json:"about"`
	Age   string `json:"age"`
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

func (pu *PetUseCase) FindByID(id string) (PetDTO, error) {
	petDB, err := pu.petInterface.FindByID(id)
	if err != nil {
		return PetDTO{}, err
	}

	pet := PetDTO{
		Name:  petDB.Name,
		About: petDB.About,
		Age:   petDB.Age,
	}

	return pet, nil
}

func (pu *PetUseCase) FindAll(page, limit int, sort string) ([]PetDTO, error) {

	petsDB, err := pu.petInterface.FindAll(page, limit, sort)
	if err != nil {
		return nil, err
	}

	var pets []PetDTO

	for _, pet := range petsDB {
		petOk := PetDTO{
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
