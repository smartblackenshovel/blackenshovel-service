package repository

import (
	"blackenshovel-service/internal/database"
	"blackenshovel-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateRole(role *models.Role) error {
	if role.ID == uuid.Nil {
		role.ID = uuid.New()
	}
	result := database.DB.Create(role)
	return result.Error
}

func GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	result := database.DB.Find(&roles)
	return roles, result.Error
}

func GetRoleByID(id uuid.UUID) (*models.Role, error) {
	var role models.Role
	result := database.DB.First(&role, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &role, nil
}

func UpdateRole(role *models.Role) error {
	result := database.DB.Save(role)
	return result.Error
}

func DeleteRole(id uuid.UUID) error {
	result := database.DB.Delete(&models.Role{}, "id = ?", id)
	return result.Error
}
