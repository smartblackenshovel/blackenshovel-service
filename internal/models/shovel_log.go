package models

import (
	"time"

	"github.com/google/uuid"
)

type ShovelLog struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	ShovelID  uuid.UUID `json:"shovel_id" gorm:"type:uuid;not null;index"`
	Status    string    `json:"status" gorm:"size:50;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
