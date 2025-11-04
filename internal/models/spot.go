package models

import (
	"time"

	"github.com/google/uuid"
)

type Spot struct {
	ID             uuid.UUID     `json:"id" gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID     `json:"organization_id" gorm:"type:uuid;not null;index"`
	Organization   *Organization `json:"organization,omitempty" gorm:"foreignKey:OrganizationID"`
	Latitude       float64       `json:"latitude" gorm:"not null"`
	Longitude      float64       `json:"longitude" gorm:"not null"`
	Altitude       float64       `json:"altitude,omitempty"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}
