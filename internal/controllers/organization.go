package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"blackenshovel-service/internal/models"

	"github.com/go-chi/chi/v5"
)

var mockOrganizations = []models.Organization{
	{ID: "1", Name: "Acme AG", LegalForm: "AG", Country: "CH"},
	{ID: "2", Name: "Beta GmbH", LegalForm: "GmbH", Country: "CH"},
}

func GetOrganizationByID(w http.ResponseWriter, r *http.Request) {
	// Placeholder for future implementation
	id := chi.URLParam(r, "id")

	for _, org := range mockOrganizations {
		if org.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(org)
			return
		}
	}
	http.Error(w, "Organization not found", http.StatusNotFound)
}

func GetOrganizations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockOrganizations)
}

func CreateOrganization(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var org models.Organization

	err := json.NewDecoder(r.Body).Decode(&org)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if org.Name == "" {
		http.Error(w, "Missing required field: name", http.StatusBadRequest)
		return
	}

	// Generate a simple ID for now (in real app, use UUID)
	org.ID = fmt.Sprintf("%d", len(mockOrganizations)+1)

	// Set timestamps
	org.CreatedAt = time.Now()
	org.UpdatedAt = time.Now()

	// Add to mock slice
	mockOrganizations = append(mockOrganizations, org)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(org)
}

func UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, org := range mockOrganizations {
		if org.ID == id {
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
			mockOrganizations[i] = org

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(org)
			return
		}
	}

	http.Error(w, "Organization not found", http.StatusNotFound)
}

func DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for i, org := range mockOrganizations {
		if org.ID == id {
			// Remove from slice
			mockOrganizations = append(mockOrganizations[:i], mockOrganizations[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Organization not found", http.StatusNotFound)
}
