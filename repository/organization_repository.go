package repository

import (
	"github.com/google/uuid"
	"github.com/savioafs/findAFriendAPI/database"
	"github.com/savioafs/findAFriendAPI/model"
	"gorm.io/gorm"
)

type OrganizationRepository struct {
	connection *gorm.DB
}

func NewOrganizationRepository() *OrganizationRepository {
	return &OrganizationRepository{
		connection: database.StartConnectionDB(),
	}
}

func (or *OrganizationRepository) CreateOrganization(organization *model.Organization) error {
	organization.ID = uuid.New()
	return or.connection.Create(organization).Error
}

func (pr *OrganizationRepository) FindOrganizationByID(id string) (*model.Organization, error) {
	var org model.Organization

	err := pr.connection.Preload("Pets").First(&org, "id = ?", id).Error

	return &org, err
}

func (pr *OrganizationRepository) FindOrganizationByName(name string) (*model.Organization, error) {
	var org model.Organization

	err := pr.connection.Preload("Pets").First(&org, "name = ?", name).Error
	return &org, err
}
