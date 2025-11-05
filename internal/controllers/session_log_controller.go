package controllers

import (
	"encoding/json"
	"net/http"

	"blackenshovel-service/internal/models"
	"blackenshovel-service/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetSessionLogs(w http.ResponseWriter, r *http.Request) {
	filter := make(map[string]interface{})

	if sessionIDStr := r.URL.Query().Get("session_id"); sessionIDStr != "" {
		if sessionID, err := uuid.Parse(sessionIDStr); err == nil {
			filter["session_id"] = sessionID
		}
	}
	if spotIDStr := r.URL.Query().Get("spot_id"); spotIDStr != "" {
		if spotID, err := uuid.Parse(spotIDStr); err == nil {
			filter["spot_id"] = spotID
		}
	}

	logs, err := repository.GetSessionLogs(filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

func GetSessionLogByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	log, err := repository.GetSessionLogByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if log == nil {
		http.Error(w, "SessionLog not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(log)
}

func CreateSessionLog(w http.ResponseWriter, r *http.Request) {
	var log models.SessionLog
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if log.ID == uuid.Nil {
		log.ID = uuid.New()
	}

	if err := repository.CreateSessionLog(&log); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(log)
}

func DeleteSessionLog(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	if err := repository.DeleteSessionLog(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
