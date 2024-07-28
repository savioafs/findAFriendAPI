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
