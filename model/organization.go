package model

import "github.com/google/uuid"

type Organization struct {
	ID       uuid.UUID `gorm:"type:VARCHAR(36);primary_key"`
	Name     string    `gorm:"type:VARCHAR(100);not null"`
	Manager  string    `gorm:"type:VARCHAR(100);not null"`
	Email    string    `gorm:"type:VARCHAR(100);not null"`
	Password string    `gorm:"type:VARCHAR(100);not null"`
	Cep      string    `gorm:"type:VARCHAR(100);not null"`
	Address  Address   `gorm:"embedded"`
	Pets     []Pet     `gorm:"foreignKey:OrganizationID"`
}

type Address struct {
	Street   string `gorm:"type:VARCHAR(100);not null"`
	District string `gorm:"type:VARCHAR(100);not null"`
	City     string `gorm:"type:VARCHAR(100);not null"`
	State    string `gorm:"type:VARCHAR(100);not null"`
}
