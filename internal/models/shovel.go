package models

import (
	"time"

	"github.com/google/uuid"
)

type Shovel struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `json:"organization_id" gorm:"type:uuid;not null;index"`
	SerialNumber   string    `json:"serial_number,omitempty" gorm:"size:100;unique"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
