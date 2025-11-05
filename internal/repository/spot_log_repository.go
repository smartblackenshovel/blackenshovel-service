package repository

import (
	"blackenshovel-service/internal/database"
	"blackenshovel-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateSpotLog(log *models.SpotLog) error {
	if log.ID == uuid.Nil {
		log.ID = uuid.New()
	}
	result := database.DB.Create(log)
	return result.Error
}

func GetSpotLogs(filter map[string]interface{}) ([]models.SpotLog, error) {
	var logs []models.SpotLog
	query := database.DB.Model(&models.SpotLog{})

	if spotID, ok := filter["spot_id"].(uuid.UUID); ok {
		query = query.Where("spot_id = ?", spotID)
	}
	if status, ok := filter["status"].(string); ok {
		query = query.Where("status = ?", status)
	}

	result := query.Find(&logs)
	return logs, result.Error
}

func GetSpotLogByID(id uuid.UUID) (*models.SpotLog, error) {
	var log models.SpotLog
	result := database.DB.First(&log, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &log, nil
}

func DeleteSpotLog(id uuid.UUID) error {
	result := database.DB.Delete(&models.SpotLog{}, "id = ?", id)
	return result.Error
}
