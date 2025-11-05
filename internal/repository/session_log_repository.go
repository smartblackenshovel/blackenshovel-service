package repository

import (
	"blackenshovel-service/internal/database"
	"blackenshovel-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateSessionLog(log *models.SessionLog) error {
	if log.ID == uuid.Nil {
		log.ID = uuid.New()
	}

	result := database.DB.Create(log)
	return result.Error
}

func GetSessionLogs(filter map[string]interface{}) ([]models.SessionLog, error) {
	var logs []models.SessionLog
	query := database.DB.Model(&models.SessionLog{})

	if sessionID, ok := filter["session_id"].(uuid.UUID); ok {
		query = query.Where("session_id = ?", sessionID)
	}
	if spotID, ok := filter["spot_id"].(uuid.UUID); ok {
		query = query.Where("spot_id = ?", spotID)
	}

	result := query.Find(&logs)
	return logs, result.Error
}

func GetSessionLogByID(id uuid.UUID) (*models.SessionLog, error) {
	var log models.SessionLog
	result := database.DB.First(&log, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &log, nil
}

func DeleteSessionLog(id uuid.UUID) error {
	result := database.DB.Delete(&models.SessionLog{}, "id = ?", id)
	return result.Error
}
