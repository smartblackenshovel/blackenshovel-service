package repository

import (
	"blackenshovel-service/internal/database"
	"blackenshovel-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateSession(session *models.Session) error {
	if session.ID == uuid.Nil {
		session.ID = uuid.New()
	}
	result := database.DB.Create(session)
	return result.Error
}

func GetAllSessions(userID, shovelID *uuid.UUID) ([]models.Session, error) {
	var sessions []models.Session
	query := database.DB
	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}
	if shovelID != nil {
		query = query.Where("shovel_id = ?", *shovelID)
	}
	result := query.Find(&sessions)
	return sessions, result.Error
}

func GetSessionByID(id uuid.UUID) (*models.Session, error) {
	var session models.Session
	result := database.DB.First(&session, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &session, nil
}

func UpdateSession(session *models.Session) error {
	result := database.DB.Save(session)
	return result.Error
}

func DeleteSession(id uuid.UUID) error {
	result := database.DB.Delete(&models.Session{}, "id = ?", id)
	return result.Error
}
