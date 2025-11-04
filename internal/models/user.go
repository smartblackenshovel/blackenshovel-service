package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID     `json:"id" gorm:"type:uuid;primaryKey"`
	Name           string        `json:"name" gorm:"size:255;not null"`
	Email          string        `json:"email,omitempty" gorm:"size:255"` // optional, not unique
	RoleID         uuid.UUID     `json:"role_id" gorm:"type:uuid;not null;index"`
	Role           *Role         `json:"role,omitempty" gorm:"foreignKey:RoleID"`
	OrganizationID uuid.UUID     `json:"organization_id" gorm:"type:uuid;not null;index"`
	Organization   *Organization `json:"organization,omitempty" gorm:"foreignKey:OrganizationID"`
	CreatedAt      time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
}
