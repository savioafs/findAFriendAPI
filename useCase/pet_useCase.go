package useCase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/savioafs/findAFriendAPI/dto"
	"github.com/savioafs/findAFriendAPI/model"
	"github.com/savioafs/findAFriendAPI/repository"
)

type PetUseCase struct {
	petStorer          repository.PetStorer
	organizationStorer repository.OrganiztionStorer
}

func NewPetUseCase() *PetUseCase {
	return &PetUseCase{
		petStorer:          repository.NewPetRepository(),
		organizationStorer: repository.NewOrganizationRepository(),
	}
}

func (pu *PetUseCase) CreatePet(petRequest dto.PetDTO) error {
	petFind, err := pu.petStorer.FindPetByName(petRequest.Name)
	if err != nil {
		return err
	}

	if petFind != nil {
		return fmt.Errorf("pet already registred: %v", err)
	}

	_, err = pu.organizationStorer.FindOrganizationByID(petRequest.OrganizationID)
	if err != nil {
		return fmt.Errorf("organization does not exists: %v", err)
	}

	organizationID, err := uuid.Parse(petRequest.OrganizationID)
	if err != nil {
		return fmt.Errorf("invalid organization ID: %v", err)
	}

	pet := model.Pet{
		Name:           petRequest.Name,
		About:          petRequest.About,
		Age:            petRequest.Age,
		Size:           petRequest.Size,
		EnergyLevel:    petRequest.EnergyLevel,
		DependeceLevel: petRequest.DependeceLevel,
		Ambience:       petRequest.Ambience,
		Requirements:   petRequest.Requirements,
		OrganizationID: organizationID,
	}

	err = pu.petStorer.CreatePet(&pet)
	if err != nil {
		return err
	}

	return nil
}

func (pu *PetUseCase) FindPetByID(id string) (dto.PetDTO, error) {
	petDB, err := pu.petStorer.FindPetByID(id)
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

func (pu *PetUseCase) FindAllPets(page, limit int, sort string) ([]dto.PetDTO, error) {

	petsDB, err := pu.petStorer.FindAllPets(page, limit, sort)
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

func (pu *PetUseCase) DeletePetByID(id string) error {
	pet, err := pu.petStorer.FindPetByID(id)
	if err != nil {
		return err
	}

	err = pu.petStorer.DeletePet(pet)

	return err
}
