package controllers

import (
	"encoding/json"
	"net/http"

	"blackenshovel-service/internal/models"
	"blackenshovel-service/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetSpotLogs(w http.ResponseWriter, r *http.Request) {
	filter := make(map[string]interface{})

	if spotIDStr := r.URL.Query().Get("spot_id"); spotIDStr != "" {
		if spotID, err := uuid.Parse(spotIDStr); err == nil {
			filter["spot_id"] = spotID
		}
	}
	if status := r.URL.Query().Get("status"); status != "" {
		filter["status"] = status
	}

	logs, err := repository.GetSpotLogs(filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

func GetSpotLogByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	log, err := repository.GetSpotLogByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if log == nil {
		http.Error(w, "SpotLog not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(log)
}

func CreateSpotLog(w http.ResponseWriter, r *http.Request) {
	var log models.SpotLog
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if log.ID == uuid.Nil {
		log.ID = uuid.New()
	}

	if err := repository.CreateSpotLog(&log); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(log)
}

func DeleteSpotLog(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	if err := repository.DeleteSpotLog(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
