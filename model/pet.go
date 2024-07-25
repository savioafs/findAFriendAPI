package model

import (
	"time"

	"github.com/google/uuid"
)

type Pet struct {
	ID             uuid.UUID `gorm:"type:VARCHAR(36);primary_key"`
	Name           string    `gorm:"type:VARCHAR(100)"`
	About          string    `gorm:"type:TEXT"`
	Age            string    `gorm:"type:VARCHAR(10)"`
	Size           string    `gorm:"type:VARCHAR(20)"`
	EnergyLevel    string    `gorm:"type:VARCHAR(20)"`
	DependeceLevel string    `gorm:"type:VARCHAR(20)"`
	Ambience       string    `gorm:"type:VARCHAR(50)"`
	Photos         []string  `gorm:"type:TEXT[]"`
	Requirements   string    `gorm:"type:TEXT"`
	CreatedAt      time.Time `gorm:"type:TIMESTAMP"`
}
