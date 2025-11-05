package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"blackenshovel-service/internal/models"
	"blackenshovel-service/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetSessions(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	shovelIDStr := r.URL.Query().Get("shovel_id")

	var userID, shovelID *uuid.UUID

	if userIDStr != "" {
		if uid, err := uuid.Parse(userIDStr); err == nil {
			userID = &uid
		}
	}
	if shovelIDStr != "" {
		if sid, err := uuid.Parse(shovelIDStr); err == nil {
			shovelID = &sid
		}
	}

	sessions, err := repository.GetAllSessions(userID, shovelID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

func GetSessionByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	session, err := repository.GetSessionByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session == nil {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

func CreateSession(w http.ResponseWriter, r *http.Request) {
	var session models.Session
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	session.ID = uuid.New()
	now := time.Now()
	session.CreatedAt = now
	session.UpdatedAt = now

	if err := repository.CreateSession(&session); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(session)
}

func UpdateSession(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	session, err := repository.GetSessionByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session == nil {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if userIDStr, ok := updates["user_id"].(string); ok {
		if uid, err := uuid.Parse(userIDStr); err == nil {
			session.UserID = uid
		}
	}
	if shovelIDStr, ok := updates["shovel_id"].(string); ok {
		if sid, err := uuid.Parse(shovelIDStr); err == nil {
			session.ShovelID = sid
		}
	}
	if startedAtStr, ok := updates["started_at"].(string); ok {
		if t, err := time.Parse(time.RFC3339, startedAtStr); err == nil {
			session.StartedAt = &t
		}
	}
	if endedAtStr, ok := updates["ended_at"].(string); ok {
		if t, err := time.Parse(time.RFC3339, endedAtStr); err == nil {
			session.EndedAt = &t
		}
	}

	session.UpdatedAt = time.Now()

	if err := repository.UpdateSession(session); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	if err := repository.DeleteSession(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
