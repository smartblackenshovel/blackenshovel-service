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

func GetSpots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orgIDStr := r.URL.Query().Get("organization_id")

	if orgIDStr != "" {
		orgID, err := uuid.Parse(orgIDStr)
		if err != nil {
			http.Error(w, "Invalid organization_id", http.StatusBadRequest)
			return
		}
		spots, err := repository.GetSpotsByOrganization(orgID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(spots)
		return
	}

	spots, err := repository.GetAllSpots()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(spots)
}

func CreateSpot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var spot models.Spot
	if err := json.NewDecoder(r.Body).Decode(&spot); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if spot.OrganizationID == uuid.Nil {
		http.Error(w, "Missing organization_id", http.StatusBadRequest)
		return
	}

	spot.ID = uuid.New()
	spot.CreatedAt = time.Now()
	spot.UpdatedAt = time.Now()

	if err := repository.CreateSpot(&spot); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(spot)
}

func GetSpotByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	spot, err := repository.GetSpotByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if spot == nil {
		http.Error(w, "Spot not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(spot)
}

func UpdateSpot(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	spot, err := repository.GetSpotByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if spot == nil {
		http.Error(w, "Spot not found", http.StatusNotFound)
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if orgIDStr, ok := updates["organization_id"].(string); ok {
		if orgID, err := uuid.Parse(orgIDStr); err == nil {
			spot.OrganizationID = orgID
		}
	}
	if lat, ok := updates["latitude"].(float64); ok {
		spot.Latitude = lat
	}
	if lng, ok := updates["longitude"].(float64); ok {
		spot.Longitude = lng
	}
	if alt, ok := updates["altitude"].(float64); ok {
		spot.Altitude = alt
	}

	spot.UpdatedAt = time.Now()
	if err := repository.UpdateSpot(spot); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(spot)
}

func DeleteSpot(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	if err := repository.DeleteSpot(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
