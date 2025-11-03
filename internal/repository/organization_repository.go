package repository

import (
	"blackenshovel-service/internal/database"
	"blackenshovel-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateOrganization(org *models.Organization) error {

	if org.ID == uuid.Nil {
		org.ID = uuid.New()
	}

	result := database.DB.Create(org)
	return result.Error
}

func GetAllOrganizations() ([]models.Organization, error) {
	var orgs []models.Organization
	result := database.DB.Find(&orgs)
	return orgs, result.Error
}

func GetOrganizationByID(id uuid.UUID) (*models.Organization, error) {
	var org models.Organization
	result := database.DB.First(&org, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &org, nil
}

func UpdateOrganization(org *models.Organization) error {
	result := database.DB.Save(org)
	return result.Error
}

func DeleteOrganization(id uuid.UUID) error {
	result := database.DB.Delete(&models.Organization{}, "id = ?", id)
	return result.Error
}
