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
