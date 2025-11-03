package models

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID                 uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name               string    `json:"name" gorm:"size:255;not null"`
	LegalForm          string    `json:"legal_form,omitempty" gorm:"size:50"`
	RegistrationNumber string    `json:"registration_number,omitempty" gorm:"size:20"`
	Street             string    `json:"street,omitempty" gorm:"size:255"`
	PostalCode         string    `json:"postal_code,omitempty" gorm:"size:20"`
	City               string    `json:"city,omitempty" gorm:"size:100"`
	Country            string    `json:"country,omitempty" gorm:"size:2;default:CH"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
