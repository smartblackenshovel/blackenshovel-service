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

func GetShovels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var orgID *uuid.UUID
	if orgIDParam := r.URL.Query().Get("organization_id"); orgIDParam != "" {
		parsedID, err := uuid.Parse(orgIDParam)
		if err != nil {
			http.Error(w, "Invalid organization_id UUID", http.StatusBadRequest)
			return
		}
		orgID = &parsedID
	}

	var serialNumber *string
	if sn := r.URL.Query().Get("serial_number"); sn != "" {
		serialNumber = &sn
	}

	shovels, err := repository.GetAllShovels(orgID, serialNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shovels)
}

func CreateShovel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var shovel models.Shovel
	if err := json.NewDecoder(r.Body).Decode(&shovel); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shovel.ID = uuid.New()
	shovel.CreatedAt = now()
	shovel.UpdatedAt = now()

	if err := repository.CreateShovel(&shovel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shovel)
}

func GetShovelByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	shovel, err := repository.GetShovelByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if shovel == nil {
		http.Error(w, "Shovel not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shovel)
}

func UpdateShovel(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	shovel, err := repository.GetShovelByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if shovel == nil {
		http.Error(w, "Shovel not found", http.StatusNotFound)
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if orgIDStr, ok := updates["organization_id"].(string); ok {
		if orgID, err := uuid.Parse(orgIDStr); err == nil {
			shovel.OrganizationID = orgID
		}
	}

	if serial, ok := updates["serial_number"].(string); ok {
		shovel.SerialNumber = serial
	}

	shovel.UpdatedAt = time.Now()

	if err := repository.UpdateShovel(shovel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shovel)
}

func DeleteShovel(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	if err := repository.DeleteShovel(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func now() (t time.Time) {
	return time.Now()
}
