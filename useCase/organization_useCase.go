package useCase

import (
	"github.com/savioafs/findAFriendAPI/dto"
	"github.com/savioafs/findAFriendAPI/model"
	"github.com/savioafs/findAFriendAPI/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/logger"
)

type OrganizationUseCase struct {
	organizationStorer repository.OrganiztionStorer
}

func NewOrganizationUseCase() *OrganizationUseCase {
	return &OrganizationUseCase{
		organizationStorer: repository.NewOrganizationRepository(),
	}
}

func (ou *OrganizationUseCase) CreateOrganization(orgRequest dto.OrganizationDTO) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(orgRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	organization := model.Organization{
		Name:     orgRequest.Name,
		Manager:  orgRequest.Manager,
		Email:    orgRequest.Email,
		Password: string(hash),
		Cep:      orgRequest.Cep,
		Address: model.Address{
			Street:   orgRequest.Street,
			District: orgRequest.District,
			City:     orgRequest.City,
			State:    orgRequest.State,
		},
	}

	err = ou.organizationStorer.CreateOrganization(&organization)
	if err != nil {
		return err
	}

	return nil
}

func (ou *OrganizationUseCase) FindOrganizationByID(id string) (dto.OrganizationDTO, error) {
	org, err := ou.organizationStorer.FindOrganizationByID(id)
	if err != nil && err != logger.ErrRecordNotFound {
		return dto.OrganizationDTO{}, nil
	}

	var pets []dto.PetDTO

	for _, pet := range org.Pets {
		pets = append(pets, dto.NewPetDTO(pet))
	}

	orgDTO := dto.NewOrganizationDTO(*org, pets)

	return orgDTO, nil
}
