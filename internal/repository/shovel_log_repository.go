package repository

import (
	"blackenshovel-service/internal/database"
	"blackenshovel-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateShovelLog(log *models.ShovelLog) error {
	return database.DB.Create(log).Error
}

func GetShovelLogByID(id uuid.UUID) (*models.ShovelLog, error) {
	var log models.ShovelLog
	if err := database.DB.First(&log, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &log, nil
}

func ListShovelLogs(shovelID *uuid.UUID, status *string, limit, offset int) ([]models.ShovelLog, error) {
	var logs []models.ShovelLog
	query := database.DB.Model(&models.ShovelLog{})

	if shovelID != nil {
		query = query.Where("shovel_id = ?", *shovelID)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Order("created_at DESC").Find(&logs).Error; err != nil {
		return nil, err
	}

	return logs, nil
}

func DeleteShovelLog(id uuid.UUID) error {
	return database.DB.Delete(&models.ShovelLog{}, "id = ?", id).Error
}
