package models

import (
	"time"

	"github.com/google/uuid"
)

type SpotLog struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	SpotID    uuid.UUID  `json:"spot_id" gorm:"type:uuid;not null;index"`
	Spot      *Spot      `json:"spot,omitempty" gorm:"foreignKey:SpotID"`
	Status    string     `json:"status" gorm:"size:20;not null"`
	SessionID *uuid.UUID `json:"session_id,omitempty" gorm:"type:uuid;index"`
	Session   *Session   `json:"session,omitempty" gorm:"foreignKey:SessionID"`
	CreatedAt time.Time  `json:"created_at"`
}
