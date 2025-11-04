package repository

import (
	"blackenshovel-service/internal/database"
	"blackenshovel-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateShovel(shovel *models.Shovel) error {
	if shovel.ID == uuid.Nil {
		shovel.ID = uuid.New()
	}
	result := database.DB.Create(shovel)
	return result.Error
}

func GetAllShovels(organizationID *uuid.UUID) ([]models.Shovel, error) {
	var shovels []models.Shovel
	query := database.DB
	if organizationID != nil {
		query = query.Where("organization_id = ?", *organizationID)
	}
	result := query.Find(&shovels)
	return shovels, result.Error
}

func GetShovelByID(id uuid.UUID) (*models.Shovel, error) {
	var shovel models.Shovel
	result := database.DB.First(&shovel, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &shovel, nil
}

func UpdateShovel(shovel *models.Shovel) error {
	result := database.DB.Save(shovel)
	return result.Error
}

func DeleteShovel(id uuid.UUID) error {
	result := database.DB.Delete(&models.Shovel{}, "id = ?", id)
	return result.Error
}
