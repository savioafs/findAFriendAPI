package useCase

import (
	"github.com/savioafs/findAFriendAPI/dto"
	"github.com/savioafs/findAFriendAPI/model"
	"github.com/savioafs/findAFriendAPI/repository"
)

type OrganizationUseCase struct {
	organizationStorer repository.OrganizationRepository
}

func NewOrganizationUseCase() *OrganizationUseCase {
	return &OrganizationUseCase{
		organizationStorer: *repository.NewOrganizationRepository(),
	}
}

func (ou *OrganizationUseCase) CreateOrganization(orgRequest dto.OrganizationDTO) error {
	organization := model.Organization{
		Name:     orgRequest.Name,
		Manager:  orgRequest.Manager,
		Email:    orgRequest.Email,
		Password: orgRequest.Password,
		Cep:      orgRequest.Cep,
		Address: model.Address{
			Street:   orgRequest.Street,
			District: orgRequest.District,
			City:     orgRequest.City,
			State:    orgRequest.State,
		},
	}

	err := ou.organizationStorer.CreateOrganization(&organization)
	if err != nil {
		return err
	}

	return nil
}

func (ou *OrganizationUseCase) FindByID(id string) (dto.OrganizationDTO, error) {
	org, err := ou.organizationStorer.FindByID(id)
	if err != nil {
		return dto.OrganizationDTO{}, nil
	}

	var pets []dto.PetDTO

	for _, pet := range org.Pets {
		pets = append(pets, dto.NewPetDTO(pet))
	}

	orgDTO := dto.NewOrganizationDTO(*org, pets)

	return orgDTO, nil
}
