package models

import (
	"time"

	"github.com/google/uuid"
)

type SessionLog struct {
	ID                  uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	SessionID           uuid.UUID  `json:"session_id" gorm:"type:uuid;not null;index"`
	Session             *Session   `json:"session,omitempty" gorm:"foreignKey:SessionID"`
	SpotID              *uuid.UUID `json:"spot_id,omitempty" gorm:"type:uuid;index"`
	Timestamp           time.Time  `json:"timestamp"`
	Latitude            float64    `json:"latitude"`
	Longitude           float64    `json:"longitude"`
	Altitude            *float64   `json:"altitude,omitempty"`
	MovementType        *string    `json:"movement_type,omitempty" gorm:"size:20"`
	MovementProbability *float64   `json:"movement_probability,omitempty"`
	EmployeeConfirmed   *bool      `json:"employee_confirmed,omitempty"`
	SensorData          *string    `json:"sensor_data,omitempty" gorm:"type:jsonb"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}
