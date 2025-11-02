package models

import "time"

type Organization struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	LegalForm          string    `json:"legal_form,omitempty"`
	RegistrationNumber string    `json:"registration_number,omitempty"`
	Street             string    `json:"street,omitempty"`
	PostalCode         string    `json:"postal_code,omitempty"`
	City               string    `json:"city,omitempty"`
	Country            string    `json:"country,omitempty"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
