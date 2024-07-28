package model

import (
	"time"

	"github.com/google/uuid"
)

type Pet struct {
	ID             uuid.UUID    `gorm:"type:VARCHAR(36);primary_key;not null"`
	Name           string       `gorm:"type:VARCHAR(100);not null"`
	About          string       `gorm:"type:TEXT;not null"`
	Age            string       `gorm:"type:VARCHAR(10);not null"`
	Size           string       `gorm:"type:VARCHAR(20);not null"`
	EnergyLevel    string       `gorm:"type:VARCHAR(20);not null"`
	DependeceLevel string       `gorm:"type:VARCHAR(20);not null"`
	Ambience       string       `gorm:"type:VARCHAR(50);not null"`
	Requirements   string       `gorm:"type:TEXT;not null"`
	CreatedAt      time.Time    `gorm:"type:TIMESTAMP"`
	OrganizationID uuid.UUID    `gorm:"type:VARCHAR(36);not null"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
}
