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

func GetOrganizationByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	org, err := repository.GetOrganizationByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if org == nil {
		http.Error(w, "Organization not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(org)
}

func GetOrganizations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orgs, err := repository.GetAllOrganizations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orgs)
}

func CreateOrganization(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var org models.Organization
	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if org.Name == "" {
		http.Error(w, "Missing required field: name", http.StatusBadRequest)
		return
	}

	if org.ID == uuid.Nil {
		org.ID = uuid.New()
	}

	org.CreatedAt = time.Now()
	org.UpdatedAt = time.Now()

	if err := repository.CreateOrganization(&org); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(org)
}

func UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	org, err := repository.GetOrganizationByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if org == nil {
		http.Error(w, "Organization not found", http.StatusNotFound)
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if name, ok := updates["name"].(string); ok {
		org.Name = name
	}
	if legalForm, ok := updates["legal_form"].(string); ok {
		org.LegalForm = legalForm
	}
	if street, ok := updates["street"].(string); ok {
		org.Street = street
	}
	if postal, ok := updates["postal_code"].(string); ok {
		org.PostalCode = postal
	}
	if city, ok := updates["city"].(string); ok {
		org.City = city
	}
	if country, ok := updates["country"].(string); ok {
		org.Country = country
	}

	org.UpdatedAt = time.Now()

	if err := repository.UpdateOrganization(org); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(org)
}

func DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	if err := repository.DeleteOrganization(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
