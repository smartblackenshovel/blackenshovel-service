package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID  `json:"user_id" gorm:"type:uuid;not null;index"`
	User      *User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	ShovelID  uuid.UUID  `json:"shovel_id" gorm:"type:uuid;not null;index"`
	Shovel    *Shovel    `json:"shovel,omitempty" gorm:"foreignKey:ShovelID"`
	StartedAt *time.Time `json:"started_at,omitempty"`
	EndedAt   *time.Time `json:"ended_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
