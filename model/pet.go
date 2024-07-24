package model

import (
	"github.com/google/uuid"
)

type Pet struct {
	ID             uuid.UUID `json:"id" gorm:"type:TEXT;primary_key"`
	Name           string    `json:"name"`
	About          string    `json:"about"`
	Age            string    `json:"age"`
	Size           string    `json:"size"`
	EnergyLevel    string    `json:"energy_level"`
	DependeceLevel string    `json:"dependece_level"`
	Ambience       string    `json:"ambience"`
	Photos         []string  `json:"photos" gorm:"type:text[]"`
	Requirements   string    `json:"requirements"`
}
