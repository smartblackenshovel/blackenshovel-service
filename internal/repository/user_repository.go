package repository

import (
	"blackenshovel-service/internal/database"
	"blackenshovel-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllUsers(filters map[string]interface{}) ([]models.User, error) {
	var users []models.User
	query := database.DB

	if orgID, ok := filters["organization_id"].(uuid.UUID); ok {
		query = query.Where("organization_id = ?", orgID)
	}
	if roleID, ok := filters["role_id"].(uuid.UUID); ok {
		query = query.Where("role_id = ?", roleID)
	}

	result := query.Find(&users)
	return users, result.Error
}

func GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	result := database.DB.Create(user)
	return result.Error
}

func UpdateUser(user *models.User) error {
	result := database.DB.Save(user)
	return result.Error
}

func DeleteUser(id uuid.UUID) error {
	result := database.DB.Delete(&models.User{}, "id = ?", id)
	return result.Error
}
