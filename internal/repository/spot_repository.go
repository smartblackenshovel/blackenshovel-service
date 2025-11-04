package repository

import (
	"blackenshovel-service/internal/database"
	"blackenshovel-service/internal/models"

	"github.com/google/uuid"
)

func CreateSpot(spot *models.Spot) error {
	if spot.ID == uuid.Nil {
		spot.ID = uuid.New()
	}
	result := database.DB.Create(spot)
	return result.Error
}

func GetAllSpots() ([]models.Spot, error) {
	var spots []models.Spot
	result := database.DB.Find(&spots)
	return spots, result.Error
}

func GetSpotByID(id uuid.UUID) (*models.Spot, error) {
	var spot models.Spot
	result := database.DB.First(&spot, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &spot, nil
}

func UpdateSpot(spot *models.Spot) error {
	result := database.DB.Save(spot)
	return result.Error
}

func DeleteSpot(id uuid.UUID) error {
	result := database.DB.Delete(&models.Spot{}, "id = ?", id)
	return result.Error
}

func GetSpotsByOrganization(orgID uuid.UUID) ([]models.Spot, error) {
	var spots []models.Spot
	result := database.DB.Where("organization_id = ?", orgID).Find(&spots)
	return spots, result.Error
}
